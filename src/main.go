package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/quintilesims/eks-sso/pkg/auth"
	"github.com/quintilesims/eks-sso/pkg/aws"
	"github.com/quintilesims/eks-sso/pkg/config"
	"github.com/quintilesims/eks-sso/pkg/controllers"
	"github.com/quintilesims/eks-sso/pkg/kubernetes"
	"github.com/quintilesims/eks-sso/pkg/logging"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	initFlags(app)

	app.Before = func(c *cli.Context) error {
		if err := validateConfig(c); err != nil {
			return err
		}

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetOutput(logging.NewLogWriter(c.Bool("debug")))

		return nil
	}

	app.Action = func(c *cli.Context) error {
		// router & session store
		router := mux.NewRouter()
		sessionStore := sessions.NewCookieStore([]byte(c.String(config.FlagClusterName)))

		// create controller dependencies
		k8sClient, err := kubernetes.NewKubernetesClient(
			c.Bool(config.FlagInCluster),
			c.String(config.FlagKubeConfigPath),
			nil,
		)
		if err != nil {
			return err
		}

		auth0 := auth.NewAuth0Authenticator(
			c.String(config.FlagAuth0Domain),
			c.String(config.FlagAuth0Connection),
			c.String(config.FlagAuth0ClientID),
			c.String(config.FlagClusterName),
			newHTTPClient(),
		)
		awsClient := aws.NewClient(
			c.String(config.FlagClusterName),
			c.String(config.FlagAWSRegion),
		)

		// create controllers
		cs := []interface {
			GetRoutes() []controllers.Route
		}{
			controllers.NewAuthController(auth0, sessionStore),
			controllers.NewAWSController(awsClient, k8sClient, sessionStore),
			controllers.NewKubernetesController(k8sClient, sessionStore),
		}

		// register routes for all the controllers
		restrictedRouteMap := map[string]bool{}
		for _, controller := range cs {
			for _, route := range controller.GetRoutes() {
				router.HandleFunc(route.Path, route.Handler).Methods(route.Method...)
				restrictedRouteMap[route.Path] = route.Restricted
			}
		}

		isRestricted := func(route string) bool {
			restricted, ok := restrictedRouteMap[route]
			if !ok {
				return false
			}

			return restricted
		}

		// add middleware
		router.Use(controllers.NewLoggingMiddleware())
		router.Use(controllers.NewAuthenticationMiddleware(isRestricted, sessionStore))
		router.Use(controllers.NewCredentialsMiddleware(isRestricted, sessionStore))

		// register file server at root
		fs := http.FileServer(http.Dir("./ui/build"))
		router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

		// create server and start listening
		srv := &http.Server{
			Handler:      router,
			Addr:         "0.0.0.0:" + c.String("port"),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		go func() {
			log.Println("[INFO] listening on", srv.Addr)
			if err := srv.ListenAndServe(); err != nil {
				log.Fatalln("[ERROR]", err)
			}
		}()

		// gracefully shutdown via SIGINT (Ctrl+C), SIGKILL, SIGQUIT or SIGTERM
		channel := make(chan os.Signal, 1)
		signal.Notify(channel, os.Interrupt)
		<-channel

		// wait for a maximum of 30 seconds for all connections to close
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		log.Println("[INFO] stopping server")
		srv.Shutdown(ctx)
		os.Exit(0)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: time.Second * 15,
	}
}

func initFlags(app *cli.App) {
	app.Name = config.AppName
	app.Usage = config.AppDescription
	app.HelpName = config.AppName

	for _, f := range config.GetAppFlags() {
		switch f.FlagType {
		case config.FTString:
			fl := cli.StringFlag{
				Name:   f.Name,
				EnvVar: f.EnvVar,
				Value:  f.Value,
			}
			app.Flags = append(app.Flags, fl)
		case config.FTBool:
			fl := cli.BoolFlag{
				Name:   f.Name,
				EnvVar: f.EnvVar,
			}
			app.Flags = append(app.Flags, fl)
		}
	}
}

func validateConfig(c *cli.Context) error {
	err := multierror.Error{}

	for _, f := range config.GetAppFlags() {
		log.Printf("[DEBUG] flag:'%s' value:'%s'\n", f.Name, c.String(f.Name))
		value := c.String(f.Name)

		if f.Required && value == "" {
			if err.Errors == nil {
				err.Errors = []error{}
			}

			errorMessage := fmt.Sprintf("Flag --%s ", f.Name)
			if f.EnvVar != "" {
				errorMessage += fmt.Sprintf("or EnvVar: %s ", f.EnvVar)
			}
			errorMessage += "not set"
			err.Errors = append(err.Errors, fmt.Errorf(errorMessage))
		}

		if f.Validate != nil && value != "" {
			if er := f.Validate(value); er != nil {
				err.Errors = append(err.Errors, er)
			}
		}
	}

	return err.ErrorOrNil()
}
