package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestProvisionNamespace(t *testing.T) {
	km, err := NewKubernetesClient(false, "", fake.NewSimpleClientset())
	if err != nil {
		t.Fatal(err)
	}

	namespaceName := "username-ns"
	userARN := "user-arn"
	roleBindgingName := namespaceName
	roleName := "cluster-admin"

	if err := km.ProvisionNamespace(namespaceName, userARN); err != nil {
		t.Fatal(err)
	}

	ns, err := km.client.CoreV1().Namespaces().Get(namespaceName, v1.GetOptions{})
	if err != nil {
		t.Fatal(err)
	}

	rb, err := km.client.RbacV1beta1().RoleBindings(namespaceName).Get(roleBindgingName, v1.GetOptions{})
	if err != nil {
		t.Fatal(err)
	}

	roleBindingSubject := v1beta1.Subject{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "User",
		Name:     userARN,
	}

	assert.Equal(t, namespaceName, ns.Name)
	assert.Equal(t, roleBindgingName, rb.Name)
	assert.Contains(t, rb.Subjects, roleBindingSubject)
	assert.Equal(t, roleName, rb.RoleRef.Name)
	assert.Equal(t, "ClusterRole", rb.RoleRef.Kind)
}

func TestGetNamespacesEmpty(t *testing.T) {
	km, err := NewKubernetesClient(false, "", fake.NewSimpleClientset())
	if err != nil {
		t.Fatal(err)
	}

	namespaces, err := km.GetNamespaces()
	if err != nil {
		t.Fatal(err)
	}

	ns, err := km.client.CoreV1().Namespaces().List(v1.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, namespaces, 0, ns.Items)
}
