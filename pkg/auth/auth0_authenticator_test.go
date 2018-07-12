package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth0AuthenticatorAuthenticate_ValidCreds(t *testing.T) {
	accessToken := "access-token"
	name := "name"
	givenName := "givenName"
	pictureURL := "http://url-image.io/profile.png"
	clusterName := "cluster-name"

	handleAuthentication := func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{
			"id_token":     "id-token",
			"access_token": accessToken,
			"token_type":   "bearer",
		}

		respBody, _ := json.Marshal(resp)
		fmt.Fprintf(w, "%s", string(respBody))

		assert.Equal(t, r.Method, "POST")
	}

	handleGetProfile := func(w http.ResponseWriter, r *http.Request) {
		resp := auth0ProfileResponse{
			Name:      name,
			GivenName: givenName,
			Picture:   pictureURL,
		}

		respBody, _ := json.Marshal(resp)
		fmt.Fprintf(w, "%s", string(respBody))

		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, fmt.Sprintf("Bearer %s", accessToken), r.Header.Get("Authorization"))
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case authEndpoint:
			handleAuthentication(w, r)
		case profileEndpoint:
			handleGetProfile(w, r)
		}
	}))
	defer server.Close()

	auth0Authenticator := NewAuth0Authenticator(
		server.URL,
		"connection",
		"clientID",
		clusterName,
		server.Client(),
	)

	auth0Resp, err := auth0Authenticator.Authenticate("valid username", "valid password")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, auth0Resp.Authenticated)
	assert.Equal(t, name, auth0Resp.Username)
	assert.Equal(t, givenName, auth0Resp.GivenName)
	assert.Equal(t, pictureURL, auth0Resp.Picture)
	assert.Equal(t, clusterName, auth0Resp.Cluster)
}

func TestAuth0AuthenticatorAuthenticate_InvalidCreds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
	}))
	defer server.Close()

	auth0Authenticator := NewAuth0Authenticator(
		server.URL,
		"connection",
		"clientID",
		"clusterName",
		server.Client(),
	)

	_, err := auth0Authenticator.Authenticate("invalid username", "invalid password")
	if err == nil {
		t.Fatal("expected authentication error")
	}
}
