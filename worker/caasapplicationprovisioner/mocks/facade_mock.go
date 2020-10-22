// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/worker/caasapplicationprovisioner (interfaces: CAASProvisionerFacade)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v8"
	caasapplicationprovisioner "github.com/juju/juju/api/caasapplicationprovisioner"
	charms "github.com/juju/juju/api/common/charms"
	params "github.com/juju/juju/apiserver/params"
	life "github.com/juju/juju/core/life"
	resources "github.com/juju/juju/core/resources"
	status "github.com/juju/juju/core/status"
	watcher "github.com/juju/juju/core/watcher"
	names "github.com/juju/names/v4"
	reflect "reflect"
)

// MockCAASProvisionerFacade is a mock of CAASProvisionerFacade interface
type MockCAASProvisionerFacade struct {
	ctrl     *gomock.Controller
	recorder *MockCAASProvisionerFacadeMockRecorder
}

// MockCAASProvisionerFacadeMockRecorder is the mock recorder for MockCAASProvisionerFacade
type MockCAASProvisionerFacadeMockRecorder struct {
	mock *MockCAASProvisionerFacade
}

// NewMockCAASProvisionerFacade creates a new mock instance
func NewMockCAASProvisionerFacade(ctrl *gomock.Controller) *MockCAASProvisionerFacade {
	mock := &MockCAASProvisionerFacade{ctrl: ctrl}
	mock.recorder = &MockCAASProvisionerFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCAASProvisionerFacade) EXPECT() *MockCAASProvisionerFacadeMockRecorder {
	return m.recorder
}

// ApplicationCharmURL mocks base method
func (m *MockCAASProvisionerFacade) ApplicationCharmURL(arg0 string) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationCharmURL", arg0)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationCharmURL indicates an expected call of ApplicationCharmURL
func (mr *MockCAASProvisionerFacadeMockRecorder) ApplicationCharmURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationCharmURL", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ApplicationCharmURL), arg0)
}

// ApplicationOCIResources mocks base method
func (m *MockCAASProvisionerFacade) ApplicationOCIResources(arg0 string) (map[string]resources.DockerImageDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationOCIResources", arg0)
	ret0, _ := ret[0].(map[string]resources.DockerImageDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationOCIResources indicates an expected call of ApplicationOCIResources
func (mr *MockCAASProvisionerFacadeMockRecorder) ApplicationOCIResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationOCIResources", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ApplicationOCIResources), arg0)
}

// CharmInfo mocks base method
func (m *MockCAASProvisionerFacade) CharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo
func (mr *MockCAASProvisionerFacadeMockRecorder) CharmInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).CharmInfo), arg0)
}

// GarbageCollect mocks base method
func (m *MockCAASProvisionerFacade) GarbageCollect(arg0 string, arg1 []names.Tag, arg2 int, arg3 []string, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GarbageCollect", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// GarbageCollect indicates an expected call of GarbageCollect
func (mr *MockCAASProvisionerFacadeMockRecorder) GarbageCollect(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GarbageCollect", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).GarbageCollect), arg0, arg1, arg2, arg3, arg4)
}

// Life mocks base method
func (m *MockCAASProvisionerFacade) Life(arg0 string) (life.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Life", arg0)
	ret0, _ := ret[0].(life.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Life indicates an expected call of Life
func (mr *MockCAASProvisionerFacadeMockRecorder) Life(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Life", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).Life), arg0)
}

// ProvisioningInfo mocks base method
func (m *MockCAASProvisionerFacade) ProvisioningInfo(arg0 string) (caasapplicationprovisioner.ProvisioningInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvisioningInfo", arg0)
	ret0, _ := ret[0].(caasapplicationprovisioner.ProvisioningInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvisioningInfo indicates an expected call of ProvisioningInfo
func (mr *MockCAASProvisionerFacadeMockRecorder) ProvisioningInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvisioningInfo", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).ProvisioningInfo), arg0)
}

// SetOperatorStatus mocks base method
func (m *MockCAASProvisionerFacade) SetOperatorStatus(arg0 string, arg1 status.Status, arg2 string, arg3 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOperatorStatus", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOperatorStatus indicates an expected call of SetOperatorStatus
func (mr *MockCAASProvisionerFacadeMockRecorder) SetOperatorStatus(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOperatorStatus", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).SetOperatorStatus), arg0, arg1, arg2, arg3)
}

// SetPassword mocks base method
func (m *MockCAASProvisionerFacade) SetPassword(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPassword indicates an expected call of SetPassword
func (mr *MockCAASProvisionerFacadeMockRecorder) SetPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPassword", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).SetPassword), arg0, arg1)
}

// Units mocks base method
func (m *MockCAASProvisionerFacade) Units(arg0 string) ([]names.Tag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units", arg0)
	ret0, _ := ret[0].([]names.Tag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockCAASProvisionerFacadeMockRecorder) Units(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).Units), arg0)
}

// UpdateUnits mocks base method
func (m *MockCAASProvisionerFacade) UpdateUnits(arg0 params.UpdateApplicationUnits) (*params.UpdateApplicationUnitsInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUnits", arg0)
	ret0, _ := ret[0].(*params.UpdateApplicationUnitsInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUnits indicates an expected call of UpdateUnits
func (mr *MockCAASProvisionerFacadeMockRecorder) UpdateUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUnits", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).UpdateUnits), arg0)
}

// WatchApplication mocks base method
func (m *MockCAASProvisionerFacade) WatchApplication(arg0 string) (watcher.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplication", arg0)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplication indicates an expected call of WatchApplication
func (mr *MockCAASProvisionerFacadeMockRecorder) WatchApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplication", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).WatchApplication), arg0)
}

// WatchApplications mocks base method
func (m *MockCAASProvisionerFacade) WatchApplications() (watcher.StringsWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplications")
	ret0, _ := ret[0].(watcher.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplications indicates an expected call of WatchApplications
func (mr *MockCAASProvisionerFacadeMockRecorder) WatchApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplications", reflect.TypeOf((*MockCAASProvisionerFacade)(nil).WatchApplications))
}
