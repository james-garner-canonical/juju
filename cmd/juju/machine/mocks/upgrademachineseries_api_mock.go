// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/machine (interfaces: UpgradeMachineSeriesAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	watcher "github.com/juju/juju/core/watcher"
	gomock "go.uber.org/mock/gomock"
)

// MockUpgradeMachineSeriesAPI is a mock of UpgradeMachineSeriesAPI interface.
type MockUpgradeMachineSeriesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockUpgradeMachineSeriesAPIMockRecorder
}

// MockUpgradeMachineSeriesAPIMockRecorder is the mock recorder for MockUpgradeMachineSeriesAPI.
type MockUpgradeMachineSeriesAPIMockRecorder struct {
	mock *MockUpgradeMachineSeriesAPI
}

// NewMockUpgradeMachineSeriesAPI creates a new mock instance.
func NewMockUpgradeMachineSeriesAPI(ctrl *gomock.Controller) *MockUpgradeMachineSeriesAPI {
	mock := &MockUpgradeMachineSeriesAPI{ctrl: ctrl}
	mock.recorder = &MockUpgradeMachineSeriesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpgradeMachineSeriesAPI) EXPECT() *MockUpgradeMachineSeriesAPIMockRecorder {
	return m.recorder
}

// BestAPIVersion mocks base method.
func (m *MockUpgradeMachineSeriesAPI) BestAPIVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestAPIVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// BestAPIVersion indicates an expected call of BestAPIVersion.
func (mr *MockUpgradeMachineSeriesAPIMockRecorder) BestAPIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestAPIVersion", reflect.TypeOf((*MockUpgradeMachineSeriesAPI)(nil).BestAPIVersion))
}

// Close mocks base method.
func (m *MockUpgradeMachineSeriesAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUpgradeMachineSeriesAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUpgradeMachineSeriesAPI)(nil).Close))
}

// GetUpgradeSeriesMessages mocks base method.
func (m *MockUpgradeMachineSeriesAPI) GetUpgradeSeriesMessages(arg0, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpgradeSeriesMessages", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpgradeSeriesMessages indicates an expected call of GetUpgradeSeriesMessages.
func (mr *MockUpgradeMachineSeriesAPIMockRecorder) GetUpgradeSeriesMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpgradeSeriesMessages", reflect.TypeOf((*MockUpgradeMachineSeriesAPI)(nil).GetUpgradeSeriesMessages), arg0, arg1)
}

// UpgradeSeriesComplete mocks base method.
func (m *MockUpgradeMachineSeriesAPI) UpgradeSeriesComplete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesComplete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeSeriesComplete indicates an expected call of UpgradeSeriesComplete.
func (mr *MockUpgradeMachineSeriesAPIMockRecorder) UpgradeSeriesComplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesComplete", reflect.TypeOf((*MockUpgradeMachineSeriesAPI)(nil).UpgradeSeriesComplete), arg0)
}

// UpgradeSeriesPrepare mocks base method.
func (m *MockUpgradeMachineSeriesAPI) UpgradeSeriesPrepare(arg0, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesPrepare", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpgradeSeriesPrepare indicates an expected call of UpgradeSeriesPrepare.
func (mr *MockUpgradeMachineSeriesAPIMockRecorder) UpgradeSeriesPrepare(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesPrepare", reflect.TypeOf((*MockUpgradeMachineSeriesAPI)(nil).UpgradeSeriesPrepare), arg0, arg1, arg2)
}

// WatchUpgradeSeriesNotifications mocks base method.
func (m *MockUpgradeMachineSeriesAPI) WatchUpgradeSeriesNotifications(arg0 string) (watcher.NotifyWatcher, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpgradeSeriesNotifications", arg0)
	ret0, _ := ret[0].(watcher.NotifyWatcher)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WatchUpgradeSeriesNotifications indicates an expected call of WatchUpgradeSeriesNotifications.
func (mr *MockUpgradeMachineSeriesAPIMockRecorder) WatchUpgradeSeriesNotifications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpgradeSeriesNotifications", reflect.TypeOf((*MockUpgradeMachineSeriesAPI)(nil).WatchUpgradeSeriesNotifications), arg0)
}
