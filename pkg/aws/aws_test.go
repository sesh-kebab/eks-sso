package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/golang/mock/gomock"
	"github.com/quintilesims/eks-sso/pkg/mock"
	"github.com/quintilesims/eks-sso/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestVerifyIAMCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stsMock := mock.NewMockSTSAPI(ctrl)

	a := ClientAPI{
		clusterName:   "cluster-name",
		clusterRegion: "us-west-2",
		stsMock:       stsMock,
	}

	output := sts.GetCallerIdentityOutput{}
	output.SetArn("arn/username")

	stsMock.EXPECT().
		GetCallerIdentity(&sts.GetCallerIdentityInput{}).
		Return(&output, nil)

	ci, err := a.AddIAMCredentials("access-id", "secret-key")
	if err != nil {
		t.Fatal(err)
	}

	result := models.AddCredentialsResponse{
		Username: "username",
		UserARN:  "arn/username",
	}

	assert.Equal(t, result, *ci)
}

func TestGetClusterInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eksMock := mock.NewMockEKSAPI(ctrl)
	clusterName := "cluster-name"
	clusterURL := "https://cluster-endpoint"

	a := ClientAPI{
		clusterName:   clusterName,
		clusterRegion: "us-west-2",
		eksMock:       eksMock,
	}

	input := eks.DescribeClusterInput{}
	input.SetName(clusterName)
	output := eks.DescribeClusterOutput{}
	output.SetCluster(&eks.Cluster{
		Name:     aws.String(clusterName),
		Endpoint: aws.String(clusterURL),
		CertificateAuthority: &eks.Certificate{
			Data: aws.String("cert-data"),
		},
	})

	eksMock.EXPECT().
		DescribeCluster(&input).
		Return(&output, nil)

	cir, err := a.GetClusterInfo("access-id", "secret-key")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, clusterName, aws.StringValue(cir.Cluster.Name))
	assert.Equal(t, clusterURL, aws.StringValue(cir.Cluster.Endpoint))
	assert.NotEmpty(t, cir.KubeConfig)
}
