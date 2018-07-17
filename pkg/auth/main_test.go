package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

type Handler func(w http.ResponseWriter, r *http.Request)

func newAuth0AuthenticatorAndServer(handler Handler) (*Auth0Authenticator, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	return NewAuth0Authenticator(server.URL, "", "", ""), server
}

func MarshalAndWrite(t *testing.T, w http.ResponseWriter, body interface{}, status int) {
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	w.WriteHeader(status)
	if _, err := fmt.Fprintln(w, string(b)); err != nil {
		t.Fatal(err)
	}
}

func Unmarshal(t *testing.T, r *http.Request, v interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.Unmarshal(body, &v); err != nil {
		t.Fatal(err)
	}
}
