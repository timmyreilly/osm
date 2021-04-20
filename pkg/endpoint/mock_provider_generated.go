// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/endpoint (interfaces: Provider)

// Package endpoint is a generated GoMock package.
package endpoint

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	identity "github.com/openservicemesh/osm/pkg/identity"
	service "github.com/openservicemesh/osm/pkg/service"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// GetID mocks base method
func (m *MockProvider) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockProviderMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockProvider)(nil).GetID))
}

// GetResolvableEndpointsForService mocks base method
func (m *MockProvider) GetResolvableEndpointsForService(arg0 service.MeshService) ([]Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvableEndpointsForService", arg0)
	ret0, _ := ret[0].([]Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolvableEndpointsForService indicates an expected call of GetResolvableEndpointsForService
func (mr *MockProviderMockRecorder) GetResolvableEndpointsForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvableEndpointsForService", reflect.TypeOf((*MockProvider)(nil).GetResolvableEndpointsForService), arg0)
}

// GetServicesForServiceAccount mocks base method
func (m *MockProvider) GetServicesForServiceAccount(arg0 identity.K8sServiceAccount) ([]service.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServicesForServiceAccount", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServicesForServiceAccount indicates an expected call of GetServicesForServiceAccount
func (mr *MockProviderMockRecorder) GetServicesForServiceAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServicesForServiceAccount", reflect.TypeOf((*MockProvider)(nil).GetServicesForServiceAccount), arg0)
}

// GetTargetPortToProtocolMappingForService mocks base method
func (m *MockProvider) GetTargetPortToProtocolMappingForService(arg0 service.MeshService) (map[uint32]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetPortToProtocolMappingForService", arg0)
	ret0, _ := ret[0].(map[uint32]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetPortToProtocolMappingForService indicates an expected call of GetTargetPortToProtocolMappingForService
func (mr *MockProviderMockRecorder) GetTargetPortToProtocolMappingForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetPortToProtocolMappingForService", reflect.TypeOf((*MockProvider)(nil).GetTargetPortToProtocolMappingForService), arg0)
}

// ListEndpointsForIdentity mocks base method
func (m *MockProvider) ListEndpointsForIdentity(arg0 identity.ServiceIdentity) []Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsForIdentity", arg0)
	ret0, _ := ret[0].([]Endpoint)
	return ret0
}

// ListEndpointsForIdentity indicates an expected call of ListEndpointsForIdentity
func (mr *MockProviderMockRecorder) ListEndpointsForIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsForIdentity", reflect.TypeOf((*MockProvider)(nil).ListEndpointsForIdentity), arg0)
}

// ListEndpointsForService mocks base method
func (m *MockProvider) ListEndpointsForService(arg0 service.MeshService) []Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsForService", arg0)
	ret0, _ := ret[0].([]Endpoint)
	return ret0
}

// ListEndpointsForService indicates an expected call of ListEndpointsForService
func (mr *MockProviderMockRecorder) ListEndpointsForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsForService", reflect.TypeOf((*MockProvider)(nil).ListEndpointsForService), arg0)
}
