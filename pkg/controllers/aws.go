package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/quintilesims/eks-sso/pkg/aws"
	"github.com/quintilesims/eks-sso/pkg/kubernetes"
	"github.com/quintilesims/eks-sso/pkg/models"
)

// NewAWSController returns new AWSController
func NewAWSController(a aws.Client, k kubernetes.Client, s sessions.Store) AWSController {
	return AWSController{
		aws:   a,
		kube:  k,
		store: s,
	}
}

// AWSController exposes endpoints related to aws resources
type AWSController struct {
	aws   aws.Client
	kube  kubernetes.Client
	store sessions.Store
}

// GetRoutes returns an array of Route
func (a AWSController) GetRoutes() []Route {
	return []Route{
		{
			Path:       "/cluster",
			Method:     []string{"GET"},
			Handler:    a.GetCluster,
			Restricted: true,
		},
		{
			Path:       "/credentials",
			Method:     []string{"POST"},
			Handler:    a.AddCredentials,
			Restricted: true,
		},
	}
}

// GetCluster returns configuration needed to connect to kubernetes cluster
func (a *AWSController) GetCluster(w http.ResponseWriter, r *http.Request) {
	iamCreds, ok := iamCredsFromContext(r.Context())
	if !ok {
		http.Error(w, "credentials not present", http.StatusBadRequest)
		return
	}

	var awsResponse *models.ClusterInfoResponse
	var err error

	namespaceName := r.URL.Query().Get("name")
	if namespaceName == "" {
		awsResponse, err = a.aws.GetClusterInfo(iamCreds.AccessID, iamCreds.SecretKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	} else {
		sa, token, err := a.kube.GetTokenForNamespace(namespaceName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		awsResponse, err = a.aws.GetKubeConfig(iamCreds.AccessID, iamCreds.SecretKey, sa, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}

	resp, err := json.Marshal(awsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

// AddCredentials stores validated IAM credentials in a secure cookie and also
// provisions kubernetes resources for the IAM user
func (a *AWSController) AddCredentials(w http.ResponseWriter, r *http.Request) {
	session, err := a.store.Get(r, cookieSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req models.IAMCredentials
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	awsResponse, err := a.aws.AddIAMCredentials(req.AccessID, req.SecretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err := a.kube.ProvisionNamespace(awsResponse.Username, awsResponse.UserARN); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(awsResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// store credentials in a secure session cookie
	creds := models.IAMCredentials{
		AccessID:  req.AccessID,
		SecretKey: req.SecretKey,
	}
	session.Values[credentialsSessionKey] = creds

	if err := session.Save(r, w); err != nil {
		log.Println("[DEBUG] error session info", err)
	}

	w.Write(resp)
}
