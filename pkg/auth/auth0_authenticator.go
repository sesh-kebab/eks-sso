package auth

import (
	"fmt"
	"log"

	"github.com/quintilesims/eks-sso/pkg/models"
	"github.com/zpatrick/rclient"
)

// Authenticator represents the interface to handle user authentication,
// typically with user credentials such as username and password.
type Authenticator interface {
	Authenticate(username, password string) (*models.AuthReponse, error)
}

// Auth0Authenticator exposes functionality to authenticate a user via the
// auth0 identity provider
type Auth0Authenticator struct {
	clientID   string
	connection string
	domain     string
	cluster    string
	client     *rclient.RestClient
}

func NewAuth0Authenticator(domain, connection, clientID, cluster string) *Auth0Authenticator {
	return &Auth0Authenticator{
		clientID:   clientID,
		connection: connection,
		domain:     domain,
		cluster:    cluster,
		client:     rclient.NewRestClient(domain),
	}
}

type oauthReq struct {
	ClientID   string `json:"client_id"`
	Connection string `json:"connection"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	GrantType  string `json:"grant_type"`
	Scope      string `json:"scope"`
}

type auth0Profile struct {
	Name      string `json:"name"`
	GivenName string `json:"given_name"`
	Picture   string `json:"picture"`
}

// Authenticates user credentials and returns the user's profile information
func (a *Auth0Authenticator) Authenticate(username, password string) (*models.AuthReponse, error) {
	log.Printf("[DEBUG] Attempting to authenticate user '%s' through Auth0", username)

	token, err := a.getAccessToken(username, password)
	if err != nil {
		return nil, err
	}

	profile, err := a.getProfile(token)
	if err != nil {
		return nil, err
	}

	resp := &models.AuthReponse{
		Authenticated: true,
		Username:      profile.Name,
		GivenName:     profile.GivenName,
		Picture:       profile.Picture,
		Cluster:       a.cluster,
	}

	return resp, nil
}

func (a *Auth0Authenticator) getAccessToken(username, password string) (string, error) {
	req := oauthReq{
		ClientID:   a.clientID,
		Connection: a.connection,
		Username:   username,
		Password:   password,
		GrantType:  "password",
		Scope:      "openid",
	}

	var resp map[string]string
	if err := a.client.Post("/oauth/ro", req, &resp); err != nil {
		return "", err
	}

	token, ok := resp["access_token"]
	if !ok {
		return "", fmt.Errorf("'access_token' not present in auth0 auth response")
	}

	return token, nil
}

func (a *Auth0Authenticator) getProfile(token string) (*auth0Profile, error) {
	var resp auth0Profile
	header := rclient.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	if err := a.client.Get("/userinfo", &resp, header); err != nil {
		return nil, err
	}

	return &resp, nil

}
