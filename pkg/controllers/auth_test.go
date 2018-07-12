package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/quasoft/memstore"
	"github.com/quintilesims/eks-sso/pkg/mock"
	"github.com/quintilesims/eks-sso/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "test-user"
	password := "test-password"
	values := map[string]string{
		"username": username,
		"password": password,
	}
	reqBody, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", "/authenticate", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ma := mock.NewMockAuthenticator(ctrl)
	a := NewAuthController(ma, memstore.NewMemStore([]byte("key")))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Authenticate)

	output := models.AuthReponse{
		Username:      username,
		Authenticated: true,
	}
	ma.EXPECT().
		Authenticate(username, password).
		Return(&output, nil)

	handler.ServeHTTP(rr, req)

	expected, _ := json.Marshal(output)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, string(expected), rr.Body.String())
}

func TestAuthenticateUnauthorized(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	username := "test-user"
	password := "test-password"
	values := map[string]string{
		"username": username,
		"password": password,
	}
	reqBody, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", "/authenticate", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ma := mock.NewMockAuthenticator(ctrl)
	a := NewAuthController(ma, memstore.NewMemStore([]byte("key")))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.Authenticate)

	ma.EXPECT().
		Authenticate(username, password).
		Return(nil, fmt.Errorf("auth0 authenticate returned an error"))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
