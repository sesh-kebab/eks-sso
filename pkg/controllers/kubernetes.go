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
