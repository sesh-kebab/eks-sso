// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quintilesims/eks-sso/pkg/auth (interfaces: Kubernetes)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKubernetes is a mock of Kubernetes interface
type MockKubernetes struct {
	ctrl     *gomock.Controller
	recorder *MockKubernetesMockRecorder
}

// MockKubernetesMockRecorder is the mock recorder for MockKubernetes
type MockKubernetesMockRecorder struct {
	mock *MockKubernetes
}

// NewMockKubernetes creates a new mock instance
func NewMockKubernetes(ctrl *gomock.Controller) *MockKubernetes {
	mock := &MockKubernetes{ctrl: ctrl}
	mock.recorder = &MockKubernetesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubernetes) EXPECT() *MockKubernetesMockRecorder {
	return m.recorder
}

// ProvisionNamespace mocks base method
func (m *MockKubernetes) ProvisionNamespace(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "ProvisionNamespace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProvisionNamespace indicates an expected call of ProvisionNamespace
func (mr *MockKubernetesMockRecorder) ProvisionNamespace(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvisionNamespace", reflect.TypeOf((*MockKubernetes)(nil).ProvisionNamespace), arg0, arg1)
}
