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
	"github.com/quintilesims/eks-sso/pkg/auth"
	"github.com/quintilesims/eks-sso/pkg/aws"
	"github.com/quintilesims/eks-sso/pkg/controllers"
	"github.com/quintilesims/eks-sso/pkg/kubernetes"
	"github.com/quintilesims/eks-sso/pkg/logging"
	"github.com/urfave/cli"
)

const (
	FlagPort            = "port"
	EVPort              = "EKS_SSO_PORT"
	FlagDebug           = "debug"
	EVDebug             = "EKS_SSO_DEBUG"
	FlagAWSRegion       = "aws-region"
	EVAWSRegion         = "EKS_SSO_AWS_REGION"
	FlagAuth0Domain     = "auth0-domain"
	EVAuth0Domain       = "EKS_SSO_AUTH0_DOMAIN"
	FlagAuth0ClientID   = "auth0-client-id"
	EVAuth0ClientID     = "EKS_SSO_AUTH0_CLIENT_ID"
	FlagAuth0Connection = "auth0-connection"
	EVAuth0Connection   = "EKS_SSO_AUTH0_CONNECTION"
	FlagClusterName     = "cluster-name"
	EVClusterName       = "EKS_SSO_CLUSTER_NAME"
	FlagEncryptionKey   = "encryption-key"
	EVEncryptionKey     = "EKS_SSO_ENCRYPTION_KEY"
	FlagInCluster       = "in-cluster"
	EVInCluster         = "EKS_SSO_IN_CLUSTER"
	FlagKubeConfigPath  = "kube-config-path"
	EVKubeConfigPath    = "EKS_SSO_KUBE_CONFIG_PATH"
)

func main() {
	app := cli.NewApp()
	app.Name = "eks-sso"
	app.Usage = "single sign solution for aws eks"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   FlagPort,
			EnvVar: EVPort,
			Value:  8080,
		},
		cli.BoolFlag{
			Name:   FlagDebug,
			EnvVar: EVDebug,
		},
		cli.StringFlag{
			Name:   FlagAWSRegion,
			EnvVar: EVAWSRegion,
			Value:  "us-west-2",
		},
		cli.StringFlag{
			Name:   FlagAuth0Domain,
			EnvVar: EVAuth0Domain,
			Value:  "https://imshealth.auth0.com",
		},
		cli.StringFlag{
			Name:   FlagAuth0ClientID,
			EnvVar: EVAuth0ClientID,
		},
		cli.StringFlag{
			Name:   FlagAuth0Connection,
			EnvVar: EVAuth0Connection,
		},
		cli.StringFlag{
			Name:   FlagClusterName,
			EnvVar: EVClusterName,
		},
		cli.StringFlag{
			Name:   FlagEncryptionKey,
			EnvVar: EVEncryptionKey,
		},
		cli.BoolFlag{
			Name:   FlagInCluster,
			EnvVar: EVInCluster,
		},
		cli.StringFlag{
			Name:   FlagKubeConfigPath,
			EnvVar: EVKubeConfigPath,
		},
	}

	app.Before = func(c *cli.Context) error {
		requiredFlags := []string{
			FlagAuth0ClientID,
			FlagAuth0Connection,
			FlagClusterName,
			FlagEncryptionKey,
			FlagKubeConfigPath,
		}

		for _, flag := range requiredFlags {
			if !c.IsSet(flag) {
				return fmt.Errorf("Required flag '%s' is not set!", flag)
			}
		}

		path := c.String(FlagKubeConfigPath)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("Invalid kubeconfig path; file '%s' does not exist", path)
		}

		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.SetOutput(logging.NewLogWriter(c.Bool(FlagDebug)))

		return nil
	}

	app.Action = func(c *cli.Context) error {
		// router & cookie store
		router := mux.NewRouter()

		key := []byte(c.String(FlagEncryptionKey))
		cookieStore := sessions.NewCookieStore(key)

		// create controller dependencies
		k8sClient, err := kubernetes.NewKubernetesClient(
			c.Bool(FlagInCluster),
			c.String(FlagKubeConfigPath),
			nil,
		)
		if err != nil {
			return err
		}

		auth0Authenticator := auth.NewAuth0Authenticator(
			c.String(FlagAuth0Domain),
			c.String(FlagAuth0Connection),
			c.String(FlagAuth0ClientID),
			c.String(FlagClusterName),
			newHTTPClient(),
		)
		awsClient := aws.NewClient(
			c.String(FlagClusterName),
			c.String(FlagAWSRegion),
		)

		// create controllers
		cs := []interface {
			GetRoutes() []controllers.Route
		}{
<<<<<<< HEAD
			controllers.NewAuthController(auth0Authenticator, cookieStore),
			controllers.NewAWSController(awsClient, k8sClient, cookieStore),
=======
			controllers.NewAuthController(auth0, sessionStore),
			controllers.NewAWSController(awsClient, k8sClient, sessionStore),
			controllers.NewKubernetesController(k8sClient, sessionStore),
>>>>>>> 2e1bdad... Add frontend and backend code for displaying namespaces in namespace tab
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

		// register file server at root
		fs := http.FileServer(http.Dir("./ui/build"))
		router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

		// add middleware
		router.Use(controllers.NewLoggingMiddleware())
		router.Use(controllers.NewAuthenticationMiddleware(isRestricted, cookieStore))
		router.Use(controllers.NewCredentialsMiddleware(isRestricted, cookieStore))

		// create server and start listening
		srv := &http.Server{
			Handler:      router,
			Addr:         fmt.Sprintf("0.0.0.0:%d", c.Int(FlagPort)),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		go func() {
			log.Println("[INFO] listening on", srv.Addr)
			if err := srv.ListenAndServe(); err != nil {
				log.Fatal(err)
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
