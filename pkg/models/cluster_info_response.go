package models

import (
	"github.com/aws/aws-sdk-go/service/eks"
)

// ClusterInfoResponse represents output from aws eks describe cluster call
type ClusterInfoResponse struct {
	Cluster    eks.Cluster `json:"cluster"`
	KubeConfig string      `json:"kubeconfig"`
}
