package config

// const application defaults
const (
	AppName            = "eks-sso"
	AppDescription     = "single sign-on solution for aws eks"
	DefaultPort        = "8080"
	DefaultAWSRegion   = "us-west-2"
	DefaultAuth0Domain = "https://imshealth.auth0.com"
)

// const command line flags
const (
	FlagPort            = "p, port"
	FlagDebug           = "d, debug"
	FlagAWSAccessKey    = "aws-access-key"
	FlagAWSSecretKey    = "aws-secret-key"
	FlagAWSRegion       = "aws-region"
	FlagAuth0Domain     = "auth0-domain"
	FlagAuth0ClientID   = "auth0-client-id"
	FlagAuth0Connection = "auth0-connection"
	FlagClusterName     = "cluster-name"
	FlagClusterRegion   = "cluster-region"
	FlagInCluster       = "in-cluster"
	FlagKubeConfigPath  = "kube-config-path"
)

// const environment varialbes
const (
	EnvVarPort            = "EKS_SSO_PORT"
	EnvVarDebug           = "EKS_SSO_DEBUG"
	EnvVarInCluster       = "EKS_SSO_IN_CLUSTER"
	EnvVarAWSRegion       = "EKS_SSO_AWS_REGION"
	EnvVarAuth0Domain     = "EKS_SSO_AUTH0_DOMAIN"
	EnvVarAuth0ClientID   = "EKS_SSO_AUTH0_CLIENT_ID"
	EnvVarAuth0Connection = "EKS_SSO_AUTH0_CONNECTION"
	EnvVarClusterName     = "EKS_SSO_CLUSTER_NAME"
	EnvVarClusterRegion   = "EKS_SSO_CLUSTER_REGION"
	EnvVarKubeConfigPath  = "EKS_SSO_KUBE_CONFIG_PATH"
)
