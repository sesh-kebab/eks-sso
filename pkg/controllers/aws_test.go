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

func TestGetCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req, err := http.NewRequest("GET", "/cluster", nil)
	if err != nil {
		t.Fatal(err)
	}

	accessID := "access-id"
	secretKey := "secret-key"
	iamCreds := models.IAMCredentials{
		AccessID:  accessID,
		SecretKey: secretKey,
	}
	req = req.WithContext(newContextWithCredentials(req.Context(), iamCreds))

	ma := mock.NewMockAWS(ctrl)
	k := mock.NewMockKubernetes(ctrl)
	a := NewAWSController(ma, k, memstore.NewMemStore())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetCluster)

	output := models.ClusterInfoResponse{}
	ma.EXPECT().
		GetClusterInfo(accessID, secretKey).
		Return(&output, nil)

	handler.ServeHTTP(rr, req)

	expected, _ := json.Marshal(output)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, string(expected), rr.Body.String())
}

func TestGetCluster_InvalidCreds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req, err := http.NewRequest("GET", "/cluster", nil)
	if err != nil {
		t.Fatal(err)
	}

	accessID := "access-id"
	secretKey := "invalid-secret-key"
	iamCreds := models.IAMCredentials{
		AccessID:  accessID,
		SecretKey: secretKey,
	}
	req = req.WithContext(newContextWithCredentials(req.Context(), iamCreds))

	ma := mock.NewMockAWS(ctrl)
	k := mock.NewMockKubernetes(ctrl)
	a := NewAWSController(ma, k, memstore.NewMemStore())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.GetCluster)

	ma.EXPECT().
		GetClusterInfo(accessID, secretKey).
		Return(nil, fmt.Errorf("invalid iam credentials"))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}

func TestAddCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accessID := "access-id"
	secretKey := "secret-key"
	values := map[string]string{
		"accessId":  accessID,
		"secretKey": secretKey,
	}
	reqBody, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", "/credentials", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	store := memstore.NewMemStore()
	session, err := store.Get(req, cookieSession)
	assert.NoError(t, err)

	ma := mock.NewMockAWS(ctrl)
	k := mock.NewMockKubernetes(ctrl)
	a := NewAWSController(ma, k, memstore.NewMemStore())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.AddCredentials)

	username := "username"
	userARN := "user-arn"
	output := models.AddCredentialsResponse{
		Username: username,
		UserARN:  userARN,
	}
	ma.EXPECT().
		AddIAMCredentials(accessID, secretKey).
		Return(&output, nil)

	k.EXPECT().
		ProvisionNamespace(username, userARN).
		Return(nil)

	handler.ServeHTTP(rr, req)

	expected, _ := json.Marshal(output)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, string(expected), rr.Body.String())
	assert.Contains(t, session.Values, credentialsSessionKey)
	assert.Equal(t, models.IAMCredentials{
		AccessID:  accessID,
		SecretKey: secretKey,
	}, session.Values[credentialsSessionKey])
}

func TestAddCredentials_InvalidCreds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accessID := "access-id"
	secretKey := "secret-key"
	values := map[string]string{
		"accessId":  accessID,
		"secretKey": secretKey,
	}
	reqBody, _ := json.Marshal(values)
	req, err := http.NewRequest("POST", "/credentials", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	ma := mock.NewMockAWS(ctrl)
	k := mock.NewMockKubernetes(ctrl)
	a := NewAWSController(ma, k, memstore.NewMemStore())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.AddCredentials)

	ma.EXPECT().
		AddIAMCredentials(accessID, secretKey).
		Return(nil, fmt.Errorf("invalid iam credentials"))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
