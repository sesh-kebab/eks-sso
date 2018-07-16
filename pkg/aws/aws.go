package aws

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/quintilesims/eks-sso/pkg/models"
)

// Client represents the interface to handle aws api calls.
type Client interface {
	AddIAMCredentials(accessID, secretKey string) (*models.AddCredentialsResponse, error)
	GetClusterInfo(accessID, secretKey string) (*models.ClusterInfoResponse, error)
}

// NewClient returns an instance of aws.client
func NewClient(clusterName, clusterRegion string) ClientAPI {
	log.Printf("[DEBUG] new aws manager with cluster:'%s', cluster-region:'%s'", clusterName, clusterRegion)

	return ClientAPI{
		clusterName:   clusterName,
		clusterRegion: clusterRegion,
	}
}

// ClientAPI exposes functionality to create/read aws resources
type ClientAPI struct {
	clusterName   string
	clusterRegion string
	stsMock       stsiface.STSAPI
	eksMock       eksiface.EKSAPI
}

func (c ClientAPI) sts(accessID, secretKey string) stsiface.STSAPI {
	if c.stsMock != nil {
		return c.stsMock
	}

	return sts.New(c.newSession(accessID, secretKey))
}

func (c ClientAPI) eks(accessID, secretKey string) eksiface.EKSAPI {
	if c.eksMock != nil {
		return c.eksMock
	}

	return eks.New(c.newSession(accessID, secretKey))
}

func (c ClientAPI) newSession(accessID, secretKey string) *session.Session {
	config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessID, secretKey, ""),
		Region:      aws.String(c.clusterRegion),
	}

	return session.New(config)
}

// AddIAMCredentials validates and stores valid credentials in an encrypted cookie
func (c ClientAPI) AddIAMCredentials(accessID, secretKey string) (*models.AddCredentialsResponse, error) {
	svc := c.sts(accessID, secretKey)

	input := &sts.GetCallerIdentityInput{}
	output, err := svc.GetCallerIdentity(input)
	if err != nil {
		return nil, err
	}

	splitARN := strings.Split(aws.StringValue(output.Arn), "/")
	if len(splitARN) < 2 {
		return nil, fmt.Errorf("arn: %s doesn't contain '/'", aws.StringValue(output.Arn))
	}

	name := splitARN[1]
	return &models.AddCredentialsResponse{
		Username: name,
		UserARN:  aws.StringValue(output.Arn),
	}, nil
}

// GetClusterInfo returns eks cluster information of the cluster this service is running on
// and kubeconfig to connect to the cluster
func (c ClientAPI) GetClusterInfo(accessID, secretKey string) (*models.ClusterInfoResponse, error) {
	svc := c.eks(accessID, secretKey)

	iamUser, err := c.AddIAMCredentials(accessID, secretKey)
	if err != nil {
		return nil, err
	}

	input := &eks.DescribeClusterInput{}
	input.SetName(c.clusterName)
	output, err := svc.DescribeCluster(input)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("kubeConfig").Parse(kubeConfigTemplate)
	if err != nil {
		return nil, err
	}

	vals := map[string]string{
		"Endpoint":    aws.StringValue(output.Cluster.Endpoint),
		"CertData":    aws.StringValue(output.Cluster.CertificateAuthority.Data),
		"Name":        aws.StringValue(output.Cluster.Name),
		"IamUserName": iamUser.Username,
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, vals); err != nil {
		return nil, err
	}

	if output.Cluster == nil {
		return nil, fmt.Errorf("describe cluster response returned empty cluster")
	}

	resp := &models.ClusterInfoResponse{
		Cluster:    *output.Cluster,
		KubeConfig: buf.String(),
	}
	return resp, nil
}

// todo: add default user namespace once we bootstrap k8s user
const kubeConfigTemplate = `
apiVersion: v1
kind: Config
preferences: {}

current-context: eks

contexts:
- context:
    cluster: eks-{{.Name}}
    namespace: {{.IamUserName}}
    user: iam-user
  name: eks

clusters:
- cluster:
    server: {{.Endpoint}}
    certificate-authority-data: {{.CertData}}
  name: eks-{{.Name}}

users:
- name: iam-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1alpha1
      command: heptio-authenticator-aws
      args:
        - "token"
        - "-i"
        - "{{.Name}}"
`
