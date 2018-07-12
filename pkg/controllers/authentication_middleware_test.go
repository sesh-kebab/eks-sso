package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/quasoft/memstore"
	"github.com/stretchr/testify/assert"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {}

func TestAuthenticationMiddleware(t *testing.T) {
	isRestricted := func(r string) bool { return false }
	req := httptest.NewRequest("GET", "/", nil)
	store := memstore.NewMemStore([]byte("test-secret"))

	mw := NewAuthenticationMiddleware(isRestricted, store)
	rr := httptest.NewRecorder()

	testHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, cookieSession)

			// check session initialized with authenticated=false
			assert.NoError(t, err)
			assert.Equal(t, false, session.Values["authenticated"])
		})
	}

	router := mux.NewRouter()
	router.HandleFunc("/", dummyHandler).Methods("GET")
	router.Use(mw)
	router.Use(testHandler)

	router.ServeHTTP(rr, req)

	// check that a session cookie will be initialized
	assert.Equal(t, http.StatusOK, rr.Code)
	val, ok := rr.HeaderMap["Set-Cookie"]
	assert.Equal(t, true, ok)
	assert.Len(t, val, 1)
	assert.Equal(t, true, strings.Contains(val[0], cookieSession))
}

func TestAuthenticationMiddleware_ForbiddenRedirect(t *testing.T) {
	isRestricted := func(r string) bool { return true }
	req := httptest.NewRequest("GET", "/home", nil)
	store := memstore.NewMemStore([]byte("test-secret"))

	mw := NewAuthenticationMiddleware(isRestricted, store)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/home", dummyHandler).Methods("GET")
	router.Use(mw)

	router.ServeHTTP(rr, req)

	// check redirect reponse code and redirect header
	assert.Equal(t, http.StatusForbidden, rr.Code)
	val, ok := rr.HeaderMap["Location"]
	assert.Equal(t, true, ok)
	assert.Len(t, val, 1)
}
