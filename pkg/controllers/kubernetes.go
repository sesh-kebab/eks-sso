package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/quintilesims/eks-sso/pkg/kubernetes"
)

// NewKubernetesController returns new KubernetesController
func NewKubernetesController(k kubernetes.Client, s sessions.Store) KubernetesController {
	return KubernetesController{
		kube:  k,
		store: s,
	}
}

// KubernetesController exposes endpoints related to aws resources
type KubernetesController struct {
	kube  kubernetes.Client
	store sessions.Store
}

// GetRoutes returns an array of Route
func (k KubernetesController) GetRoutes() []Route {
	return []Route{
		{
			Path:       "/namespaces",
			Method:     []string{"GET"},
			Handler:    k.GetNamespaces,
			Restricted: true,
		},
		{
			Path:        "/namespace",
			Method:      []string{"POST"},
			Handler:     k.CreatePrivateNamespace,
			Restricted:  true,
			QueryParams: []string{"name", "{name}"},
		},
	}
}

// GetNamespaces returns configuration needed to connect to kubernetes cluster
func (k *KubernetesController) GetNamespaces(w http.ResponseWriter, r *http.Request) {
	namspaces, err := k.kube.GetNamespaces()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(namspaces)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

// CreatePrivateNamespace ...
// todo: probably should read name from request body instead of query string
func (a *KubernetesController) CreatePrivateNamespace(w http.ResponseWriter, r *http.Request) {
	namespaceName := r.URL.Query().Get("name")
	if namespaceName == "" {
		http.Error(w, "invalid request: missing name query string parameter", http.StatusBadRequest)
		return
	}

	if err := a.kube.ProvisionPrivateNamespace(namespaceName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
