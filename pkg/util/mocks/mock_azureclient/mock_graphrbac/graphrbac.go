// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/graphrbac (interfaces: RBACApplicationsClient,RBACGroupsClient,ServicePrincipalsClient)

// Package mock_graphrbac is a generated GoMock package.
package mock_graphrbac

import (
	context "context"
	reflect "reflect"

	graphrbac "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockRBACApplicationsClient is a mock of RBACApplicationsClient interface
type MockRBACApplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRBACApplicationsClientMockRecorder
}

// MockRBACApplicationsClientMockRecorder is the mock recorder for MockRBACApplicationsClient
type MockRBACApplicationsClientMockRecorder struct {
	mock *MockRBACApplicationsClient
}

// NewMockRBACApplicationsClient creates a new mock instance
func NewMockRBACApplicationsClient(ctrl *gomock.Controller) *MockRBACApplicationsClient {
	mock := &MockRBACApplicationsClient{ctrl: ctrl}
	mock.recorder = &MockRBACApplicationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRBACApplicationsClient) EXPECT() *MockRBACApplicationsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRBACApplicationsClient) Create(arg0 context.Context, arg1 graphrbac.ApplicationCreateParameters) (graphrbac.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRBACApplicationsClientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRBACApplicationsClient)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockRBACApplicationsClient) Delete(arg0 context.Context, arg1 string) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockRBACApplicationsClientMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRBACApplicationsClient)(nil).Delete), arg0, arg1)
}

// List mocks base method
func (m *MockRBACApplicationsClient) List(arg0 context.Context, arg1 string) (graphrbac.ApplicationListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.ApplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRBACApplicationsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRBACApplicationsClient)(nil).List), arg0, arg1)
}

// Patch mocks base method
func (m *MockRBACApplicationsClient) Patch(arg0 context.Context, arg1 string, arg2 graphrbac.ApplicationUpdateParameters) (autorest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1, arg2)
	ret0, _ := ret[0].(autorest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockRBACApplicationsClientMockRecorder) Patch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockRBACApplicationsClient)(nil).Patch), arg0, arg1, arg2)
}

// MockRBACGroupsClient is a mock of RBACGroupsClient interface
type MockRBACGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRBACGroupsClientMockRecorder
}

// MockRBACGroupsClientMockRecorder is the mock recorder for MockRBACGroupsClient
type MockRBACGroupsClientMockRecorder struct {
	mock *MockRBACGroupsClient
}

// NewMockRBACGroupsClient creates a new mock instance
func NewMockRBACGroupsClient(ctrl *gomock.Controller) *MockRBACGroupsClient {
	mock := &MockRBACGroupsClient{ctrl: ctrl}
	mock.recorder = &MockRBACGroupsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRBACGroupsClient) EXPECT() *MockRBACGroupsClientMockRecorder {
	return m.recorder
}

// GetGroupMembers mocks base method
func (m *MockRBACGroupsClient) GetGroupMembers(arg0 context.Context, arg1 string) ([]graphrbac.BasicDirectoryObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMembers", arg0, arg1)
	ret0, _ := ret[0].([]graphrbac.BasicDirectoryObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMembers indicates an expected call of GetGroupMembers
func (mr *MockRBACGroupsClientMockRecorder) GetGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMembers", reflect.TypeOf((*MockRBACGroupsClient)(nil).GetGroupMembers), arg0, arg1)
}

// MockServicePrincipalsClient is a mock of ServicePrincipalsClient interface
type MockServicePrincipalsClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicePrincipalsClientMockRecorder
}

// MockServicePrincipalsClientMockRecorder is the mock recorder for MockServicePrincipalsClient
type MockServicePrincipalsClientMockRecorder struct {
	mock *MockServicePrincipalsClient
}

// NewMockServicePrincipalsClient creates a new mock instance
func NewMockServicePrincipalsClient(ctrl *gomock.Controller) *MockServicePrincipalsClient {
	mock := &MockServicePrincipalsClient{ctrl: ctrl}
	mock.recorder = &MockServicePrincipalsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServicePrincipalsClient) EXPECT() *MockServicePrincipalsClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServicePrincipalsClient) Create(arg0 context.Context, arg1 graphrbac.ServicePrincipalCreateParameters) (graphrbac.ServicePrincipal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.ServicePrincipal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServicePrincipalsClientMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServicePrincipalsClient)(nil).Create), arg0, arg1)
}

// List mocks base method
func (m *MockServicePrincipalsClient) List(arg0 context.Context, arg1 string) (graphrbac.ServicePrincipalListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(graphrbac.ServicePrincipalListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServicePrincipalsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicePrincipalsClient)(nil).List), arg0, arg1)
}
