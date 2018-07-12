package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/quintilesims/eks-sso/pkg/models"
)

const (
	authEndpoint    = "/oauth/ro"
	profileEndpoint = "/userinfo"
)

// Authenticator represents the interface to handle user authentication,
// typically with user credentials such as username and password.
type Authenticator interface {
	Authenticate(username, password string) (*models.AuthReponse, error)
}

func NewAuth0Authenticator(domain, connection, clientID, cluster string, client *http.Client) *Auth0Authenticator {
	a := &Auth0Authenticator{
		clientID:   clientID,
		connection: connection,
		domain:     domain,
		cluster:    cluster,
		httpClient: client,
	}

	log.Printf("[DEBUG] new auth0 authenticator %#v\n", a)
	return a
}

// Auth0Authenticator exposes functionality to authenticate a user via the
// auth0 identity provider
type Auth0Authenticator struct {
	clientID   string
	connection string
	domain     string
	cluster    string
	httpClient *http.Client
}

type auth0ProfileResponse struct {
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Picture   string `json:"picture"`
}

// Authenticates user credentials and returns the user's profile information
func (a *Auth0Authenticator) Authenticate(username, password string) (*models.AuthReponse, error) {
	path, err := a.joinURL(authEndpoint)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] requesting authentication for user: '%s' from: '%s'\n",
		username,
		path,
	)

	values := map[string]string{
		"username":   username,
		"password":   password,
		"client_id":  a.clientID,
		"connection": a.connection,
		"scope":      "openid",
	}

	reqBody, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(path, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not authenticate. status:%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse map[string]string
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	token, ok := tokenResponse["access_token"]
	if !ok {
		return nil, fmt.Errorf("'access_token' not present in auth0 auth response")
	}

	pr, err := a.getAuth0Profile(token)
	if err != nil {
		return nil, err
	}

	return &models.AuthReponse{
		Authenticated: true,
		Username:      pr.Name,
		GivenName:     pr.GivenName,
		Picture:       pr.Picture,
		Cluster:       a.cluster,
	}, nil
}

func (a *Auth0Authenticator) getAuth0Profile(token string) (*auth0ProfileResponse, error) {
	path, err := a.joinURL(profileEndpoint)
	if err != nil {
		return nil, err
	}

	log.Println("[DEBUG] requesting profile info from:", path)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Could not authenticate")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var pr auth0ProfileResponse
	if err := json.Unmarshal(body, &pr); err != nil {
		return nil, err
	}

	return &pr, nil
}

func (a Auth0Authenticator) joinURL(path string) (string, error) {
	base, err := url.Parse(a.domain)
	if err != nil {
		return "", err
	}

	p, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", base.ResolveReference(p)), nil
}
