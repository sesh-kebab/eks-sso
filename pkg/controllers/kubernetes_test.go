package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/quasoft/memstore"
	"github.com/quintilesims/eks-sso/pkg/mock"
	"github.com/quintilesims/eks-sso/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetNamespaces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req, err := http.NewRequest("GET", "/namespaces", nil)
	if err != nil {
		t.Fatal(err)
	}

	mk := mock.NewMockKubernetes(ctrl)
	k := NewKubernetesController(mk, memstore.NewMemStore())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(k.GetNamespaces)

	output := []models.NamespaceResponse{}
	mk.EXPECT().
		GetNamespaces().
		Return(output, nil)

	handler.ServeHTTP(rr, req)

	expected, _ := json.Marshal(output)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, string(expected), rr.Body.String())
}
