package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/quintilesims/eks-sso/pkg/models"
)

// NewCredentialsMiddleware is responsible for checking whether the IAM Creds
// exist in the session and add credentials to the context if a restricted path
// is being accessed
func NewCredentialsMiddleware(isRestricted func(string) bool, store sessions.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, cookieSession)
			if err != nil {
				log.Println("[ERROR]", err)
				next.ServeHTTP(w, r)
				return
			}

			// if uri specified as restricted, add credentials to context if they exist
			if isRestricted(r.RequestURI) {
				creds, ok := session.Values[credentialsSessionKey]
				if !ok {
					next.ServeHTTP(w, r)
					return
				}

				iamCreds, ok := creds.(*models.IAMCredentials)
				if !ok {
					fmt.Println("couldn't decode iamcreds")
					next.ServeHTTP(w, r)
					return
				}

				// updating context with iam credentials
				r = r.WithContext(newContextWithCredentials(r.Context(), *iamCreds))
				next.ServeHTTP(w, r)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
