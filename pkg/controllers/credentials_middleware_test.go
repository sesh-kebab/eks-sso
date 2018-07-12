package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/quasoft/memstore"
	"github.com/quintilesims/eks-sso/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestCredentialsMiddleware_ContextSetIfCredsExist(t *testing.T) {
	isRestricted := func(r string) bool { return true }
	req := httptest.NewRequest("GET", "/home", nil)
	store := memstore.NewMemStore([]byte("test-secret"))
	session, err := store.Get(req, cookieSession)
	assert.NoError(t, err)

	mw := NewCredentialsMiddleware(isRestricted, store)
	rr := httptest.NewRecorder()

	iamCreds := models.IAMCredentials{
		AccessID:  "access-id",
		SecretKey: "secret-key",
	}
	session.Values[credentialsSessionKey] = iamCreds
	err = session.Save(req, rr)
	assert.NoError(t, err)

	setCookieHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			iamCreds := models.IAMCredentials{
				AccessID:  "access-id",
				SecretKey: "secret-key",
			}
			session.Values[credentialsSessionKey] = iamCreds
			err = session.Save(req, rr)
			assert.NoError(t, err)
		})
	}

	testHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("testhandler")

			val, ok := r.Context().Value(credentialsCtxKey).(models.IAMCredentials)
			assert.Equal(t, true, ok)
			assert.Equal(t, val, iamCreds)
		})
	}

	router := mux.NewRouter()
	router.HandleFunc("/home", dummyHandler).Methods("GET")
	router.Use(setCookieHandler)
	router.Use(mw)
	router.Use(testHandler)

	router.ServeHTTP(rr, req)
}

func TestCredentialsMiddleware_ContextNotSetIfCredsDontExist(t *testing.T) {
	isRestricted := func(r string) bool { return true }
	req := httptest.NewRequest("GET", "/home", nil)
	store := memstore.NewMemStore([]byte("test-secret"))
	_, err := store.Get(req, cookieSession)
	assert.NoError(t, err)

	mw := NewCredentialsMiddleware(isRestricted, store)
	rr := httptest.NewRecorder()

	testHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, ok := r.Context().Value(credentialsCtxKey).(models.IAMCredentials)
			assert.Equal(t, false, ok)
		})
	}

	router := mux.NewRouter()
	router.HandleFunc("/home", dummyHandler).Methods("GET")
	router.Use(mw)
	router.Use(testHandler)

	router.ServeHTTP(rr, req)
}

func TestCredentialsMiddleware_ContextNotSetIfNotRestrictedPath(t *testing.T) {
	isRestricted := func(r string) bool { return false }
	req := httptest.NewRequest("GET", "/not-restricted", nil)
	store := memstore.NewMemStore([]byte("test-secret"))
	session, err := store.Get(req, cookieSession)
	assert.NoError(t, err)

	mw := NewCredentialsMiddleware(isRestricted, store)
	rr := httptest.NewRecorder()

	setCookieHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			iamCreds := models.IAMCredentials{
				AccessID:  "access-id",
				SecretKey: "secret-key",
			}
			session.Values[credentialsSessionKey] = iamCreds
			err = session.Save(req, rr)
			assert.NoError(t, err)
		})
	}

	testHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, ok := r.Context().Value(credentialsCtxKey).(models.IAMCredentials)
			assert.Equal(t, false, ok)
		})
	}

	router := mux.NewRouter()
	router.HandleFunc("/not-restricted", dummyHandler).Methods("GET")
	router.Use(setCookieHandler)
	router.Use(mw)
	router.Use(testHandler)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
