package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/quintilesims/eks-sso/pkg/auth"
	"github.com/quintilesims/eks-sso/pkg/models"
)

// NewAuthController returns an instance of AuthController
func NewAuthController(a auth.Authenticator, s sessions.Store) AuthController {
	return AuthController{
		authenticator: a,
		store:         s,
	}
}

// AuthController handles initial service authentication
type AuthController struct {
	authenticator auth.Authenticator
	store         sessions.Store
}

// GetRoutes returns an array of Route
func (a AuthController) GetRoutes() []Route {
	return []Route{
		{
			Path:    "/authenticate",
			Method:  []string{"POST"},
			Handler: a.Authenticate,
		},
		{
			Path:    "/logout",
			Method:  []string{"GET"},
			Handler: a.Authenticate,
		},
	}
}

// Authenticate authenticates user and if successful also updates session cookie
// with authentication state
func (a *AuthController) Authenticate(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, cookieSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var req models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := a.authenticator.Authenticate(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp, err := json.Marshal(authResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	if err := session.Save(r, w); err != nil {
		http.Error(w, "error saving authenticated session", http.StatusInternalServerError)
	}
	w.Write(resp)
}

// Logout updates the authenticated state in the session and redirects to '/'
func (a *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, cookieSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Values["authenticated"] = false
	if err := session.Save(r, w); err != nil {
		log.Println("[ERROR] error saving authenticated session")
	}

	http.Redirect(w, r, "/", http.StatusNoContent)
}
