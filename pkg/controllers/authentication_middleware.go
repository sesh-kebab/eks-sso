package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

const cookieSession = "eks-sso.session"

// NewAuthenticationMiddleware returns middleware func to handle paths that are secure
// by checking the authenticated session state and redirecting if false
func NewAuthenticationMiddleware(isRestricted func(string) bool, store sessions.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, cookieSession)
			if err != nil {
				log.Println("[ERROR]", err)
			}

			// if uri specified as restricted, check for auth
			if isRestricted(r.RequestURI) {
				val, ok := session.Values["authenticated"]
				if !ok {
					log.Println("[DEBUG] not authenticated. redirecting")
					http.Redirect(w, r, "/", http.StatusForbidden)
					return
				}

				if authenticated, ok := val.(bool); !ok || !authenticated {
					log.Println("[DEBUG] session expired. redirecting")
					http.Redirect(w, r, "/", http.StatusForbidden)
					return
				}
			}

			// initialize new session if not set
			if len(session.Values) == 0 {
				log.Println("[DEBUG] no session found - initializing new session")
				session.Values["authenticated"] = false
				if err := session.Save(r, w); err != nil {
					http.Error(w, "error saving authenticated session", http.StatusInternalServerError)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}
