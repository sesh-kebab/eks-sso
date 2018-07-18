package kubernetes

import (
	"log"

	"github.com/quintilesims/eks-sso/pkg/models"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client exposes Kubernetes functionality
type Client interface {
	ProvisionNamespace(username string, userARN string) error
	GetNamespaces() (namespaces []models.NamespaceResponse, err error)
}

// NewKubernetesClient returns a new instance of KubernetesManager or an error if an instance couldn't
// be constructed.
func NewKubernetesClient(inCluster bool, kubeConfigPath string, client kubernetes.Interface) (ClientAPI, error) {
	log.Printf("[DEBUG] new kubernetes manager for inCluster:'%v', kubeConfigPath:'%s'",
		inCluster,
		kubeConfigPath,
	)

	if client == nil {
		log.Printf("[DEBUG] creating a new k8s client")
		config, err := getConfig(inCluster, kubeConfigPath)
		if err != nil {
			return ClientAPI{}, err
		}

		client, err = kubernetes.NewForConfig(config)
		if err != nil {
			return ClientAPI{}, err
		}
	}

	return ClientAPI{
		inCluster:      inCluster,
		kubeConfigPath: kubeConfigPath,
		client:         client,
	}, nil
}

// ClientAPI exposes functionality to create/delete resources on a k8s cluster
type ClientAPI struct {
	inCluster      bool
	kubeConfigPath string
	client         kubernetes.Interface
}

// ProvisionNamespace creates a namespace and a rolebinding to give that user
// cluster-admin permissions within the newly created namespace
func (k ClientAPI) ProvisionNamespace(username, userARN string) error {
	namespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: username,
			Labels: map[string]string{
				"eks-sso": "true",
			},
		},
	}

	// create namespace
	log.Println("[DEBUG] creating namespace:", namespace.Name)
	if _, err := k.client.CoreV1().Namespaces().Create(namespace); err != nil {
		if !errors.IsAlreadyExists(err) {
			return err
		}
	}

	// create rbac role binding for the user for ther created namespace
	log.Println("[DEBUG] creating rolebinding:", namespace.Name)
	rb := &v1beta1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: namespace.Name,
			Labels: map[string]string{
				"eks-sso": "true",
			},
		},
		Subjects: []v1beta1.Subject{
			v1beta1.Subject{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "User",
				Name:     userARN,
			},
		},
		RoleRef: v1beta1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "cluster-admin",
		},
	}

	if _, err := k.client.RbacV1beta1().RoleBindings(namespace.Name).Create(rb); err != nil {
		if !errors.IsAlreadyExists(err) {
			return err
		}
	}

	return nil
}

// GetNamespaces reads namespace and a rolebinding to give that user
// cluster-admin permissions within the newly created namespace
func (k ClientAPI) GetNamespaces() (namespaces []models.NamespaceResponse, err error) {
	namespacesList, err := k.client.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, namespace := range namespacesList.Items {
		namespaces = append(namespaces, models.NamespaceResponse{Name: namespace.Name, CreationTime: namespace.CreationTimestamp.Unix()})

	}

	log.Println("[DEBUG] listing:", namespaces)
	return namespaces, nil
}

func getConfig(inCluster bool, kubeConfigPath string) (*rest.Config, error) {
	if inCluster {
		return rest.InClusterConfig()
	}

	return clientcmd.BuildConfigFromFlags("", kubeConfigPath)
}
