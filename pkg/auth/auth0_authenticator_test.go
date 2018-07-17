package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth0AuthenticatorGetAccessToken(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.Path, "/oauth/ro")

		var req oauthReq
		Unmarshal(t, r, &req)

		assert.Equal(t, req.Username, "valid username")
		assert.Equal(t, req.Password, "valid password")

		resp := map[string]string{"access_token": "token"}
		MarshalAndWrite(t, w, resp, 200)
	}

	authenticator, server := newAuth0AuthenticatorAndServer(handler)
	defer server.Close()

	token, err := authenticator.getAccessToken("valid username", "valid password")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "token", token)
}

func TestAuth0AuthenticatorGetAccessTokenInvalidCreds(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		MarshalAndWrite(t, w, nil, 401)
	}

	authenticator, server := newAuth0AuthenticatorAndServer(handler)
	defer server.Close()

	if _, err := authenticator.getAccessToken("", ""); err == nil {
		t.Fatalf("Error was nil!")
	}
}

func TestAuth0AuthenticatorGetProfile(t *testing.T) {
	resp := auth0Profile{
		Name:      "name",
		GivenName: "given_name",
		Picture:   "picture",
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/userinfo")
		assert.Equal(t, "Bearer token", r.Header.Get("Authorization"))

		MarshalAndWrite(t, w, resp, 200)
	}

	authenticator, server := newAuth0AuthenticatorAndServer(handler)
	defer server.Close()

	profile, err := authenticator.getProfile("token")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp, *profile)
}

func TestAuth0AuthenticatorGetProfileInvalidToken(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		MarshalAndWrite(t, w, nil, 401)
	}

	authenticator, server := newAuth0AuthenticatorAndServer(handler)
	defer server.Close()

	if _, err := authenticator.getProfile(""); err == nil {
		t.Fatalf("Error was nil!")
	}
}
