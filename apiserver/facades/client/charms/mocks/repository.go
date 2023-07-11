// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm (interfaces: Repository,RepositoryFactory,CharmArchive)

// Package mocks is a generated GoMock package.
package mocks

import (
	url "net/url"
	reflect "reflect"

	charm "github.com/juju/charm/v8"
	resource "github.com/juju/charm/v8/resource"
	charm0 "github.com/juju/juju/core/charm"
	gomock "go.uber.org/mock/gomock"
	macaroon "gopkg.in/macaroon.v2"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// DownloadCharm mocks base method.
func (m *MockRepository) DownloadCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 macaroon.Slice, arg3 string) (charm0.CharmArchive, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadCharm", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm0.CharmArchive)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadCharm indicates an expected call of DownloadCharm.
func (mr *MockRepositoryMockRecorder) DownloadCharm(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadCharm", reflect.TypeOf((*MockRepository)(nil).DownloadCharm), arg0, arg1, arg2, arg3)
}

// GetDownloadURL mocks base method.
func (m *MockRepository) GetDownloadURL(arg0 *charm.URL, arg1 charm0.Origin, arg2 macaroon.Slice) (*url.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDownloadURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDownloadURL indicates an expected call of GetDownloadURL.
func (mr *MockRepositoryMockRecorder) GetDownloadURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDownloadURL", reflect.TypeOf((*MockRepository)(nil).GetDownloadURL), arg0, arg1, arg2)
}

// GetEssentialMetadata mocks base method.
func (m *MockRepository) GetEssentialMetadata(arg0 ...charm0.MetadataRequest) ([]charm0.EssentialMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEssentialMetadata", varargs...)
	ret0, _ := ret[0].([]charm0.EssentialMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEssentialMetadata indicates an expected call of GetEssentialMetadata.
func (mr *MockRepositoryMockRecorder) GetEssentialMetadata(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEssentialMetadata", reflect.TypeOf((*MockRepository)(nil).GetEssentialMetadata), arg0...)
}

// ListResources mocks base method.
func (m *MockRepository) ListResources(arg0 *charm.URL, arg1 charm0.Origin, arg2 macaroon.Slice) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResources", arg0, arg1, arg2)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockRepositoryMockRecorder) ListResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockRepository)(nil).ListResources), arg0, arg1, arg2)
}

// ResolveWithPreferredChannel mocks base method.
func (m *MockRepository) ResolveWithPreferredChannel(arg0 *charm.URL, arg1 charm0.Origin, arg2 macaroon.Slice) (*charm.URL, charm0.Origin, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveWithPreferredChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ResolveWithPreferredChannel indicates an expected call of ResolveWithPreferredChannel.
func (mr *MockRepositoryMockRecorder) ResolveWithPreferredChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveWithPreferredChannel", reflect.TypeOf((*MockRepository)(nil).ResolveWithPreferredChannel), arg0, arg1, arg2)
}

// MockRepositoryFactory is a mock of RepositoryFactory interface.
type MockRepositoryFactory struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryFactoryMockRecorder
}

// MockRepositoryFactoryMockRecorder is the mock recorder for MockRepositoryFactory.
type MockRepositoryFactoryMockRecorder struct {
	mock *MockRepositoryFactory
}

// NewMockRepositoryFactory creates a new mock instance.
func NewMockRepositoryFactory(ctrl *gomock.Controller) *MockRepositoryFactory {
	mock := &MockRepositoryFactory{ctrl: ctrl}
	mock.recorder = &MockRepositoryFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryFactory) EXPECT() *MockRepositoryFactoryMockRecorder {
	return m.recorder
}

// GetCharmRepository mocks base method.
func (m *MockRepositoryFactory) GetCharmRepository(arg0 charm0.Source) (charm0.Repository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmRepository", arg0)
	ret0, _ := ret[0].(charm0.Repository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharmRepository indicates an expected call of GetCharmRepository.
func (mr *MockRepositoryFactoryMockRecorder) GetCharmRepository(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmRepository", reflect.TypeOf((*MockRepositoryFactory)(nil).GetCharmRepository), arg0)
}

// MockCharmArchive is a mock of CharmArchive interface.
type MockCharmArchive struct {
	ctrl     *gomock.Controller
	recorder *MockCharmArchiveMockRecorder
}

// MockCharmArchiveMockRecorder is the mock recorder for MockCharmArchive.
type MockCharmArchiveMockRecorder struct {
	mock *MockCharmArchive
}

// NewMockCharmArchive creates a new mock instance.
func NewMockCharmArchive(ctrl *gomock.Controller) *MockCharmArchive {
	mock := &MockCharmArchive{ctrl: ctrl}
	mock.recorder = &MockCharmArchiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmArchive) EXPECT() *MockCharmArchiveMockRecorder {
	return m.recorder
}

// Actions mocks base method.
func (m *MockCharmArchive) Actions() *charm.Actions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(*charm.Actions)
	return ret0
}

// Actions indicates an expected call of Actions.
func (mr *MockCharmArchiveMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockCharmArchive)(nil).Actions))
}

// Config mocks base method.
func (m *MockCharmArchive) Config() *charm.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*charm.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockCharmArchiveMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCharmArchive)(nil).Config))
}

// LXDProfile mocks base method.
func (m *MockCharmArchive) LXDProfile() *charm.LXDProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(*charm.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockCharmArchiveMockRecorder) LXDProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockCharmArchive)(nil).LXDProfile))
}

// Manifest mocks base method.
func (m *MockCharmArchive) Manifest() *charm.Manifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*charm.Manifest)
	return ret0
}

// Manifest indicates an expected call of Manifest.
func (mr *MockCharmArchiveMockRecorder) Manifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*MockCharmArchive)(nil).Manifest))
}

// Meta mocks base method.
func (m *MockCharmArchive) Meta() *charm.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(*charm.Meta)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockCharmArchiveMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockCharmArchive)(nil).Meta))
}

// Metrics mocks base method.
func (m *MockCharmArchive) Metrics() *charm.Metrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metrics")
	ret0, _ := ret[0].(*charm.Metrics)
	return ret0
}

// Metrics indicates an expected call of Metrics.
func (mr *MockCharmArchiveMockRecorder) Metrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockCharmArchive)(nil).Metrics))
}

// Revision mocks base method.
func (m *MockCharmArchive) Revision() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision.
func (mr *MockCharmArchiveMockRecorder) Revision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockCharmArchive)(nil).Revision))
}

// Version mocks base method.
func (m *MockCharmArchive) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockCharmArchiveMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCharmArchive)(nil).Version))
}
