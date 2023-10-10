// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/secrets (interfaces: SecretsState,SecretsConsumer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	state "github.com/juju/juju/state"
	names "github.com/juju/names/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretsState is a mock of SecretsState interface.
type MockSecretsState struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsStateMockRecorder
}

// MockSecretsStateMockRecorder is the mock recorder for MockSecretsState.
type MockSecretsStateMockRecorder struct {
	mock *MockSecretsState
}

// NewMockSecretsState creates a new mock instance.
func NewMockSecretsState(ctrl *gomock.Controller) *MockSecretsState {
	mock := &MockSecretsState{ctrl: ctrl}
	mock.recorder = &MockSecretsStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsState) EXPECT() *MockSecretsStateMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockSecretsState) CreateSecret(arg0 *secrets.URI, arg1 state.CreateSecretParams) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockSecretsStateMockRecorder) CreateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretsState)(nil).CreateSecret), arg0, arg1)
}

// DeleteSecret mocks base method.
func (m *MockSecretsState) DeleteSecret(arg0 *secrets.URI, arg1 ...int) ([]secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSecret", varargs...)
	ret0, _ := ret[0].([]secrets.ValueRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretsStateMockRecorder) DeleteSecret(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretsState)(nil).DeleteSecret), varargs...)
}

// GetSecret mocks base method.
func (m *MockSecretsState) GetSecret(arg0 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretsStateMockRecorder) GetSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretsState)(nil).GetSecret), arg0)
}

// GetSecretRevision mocks base method.
func (m *MockSecretsState) GetSecretRevision(arg0 *secrets.URI, arg1 int) (*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretRevision", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretRevisionMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretRevision indicates an expected call of GetSecretRevision.
func (mr *MockSecretsStateMockRecorder) GetSecretRevision(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretRevision", reflect.TypeOf((*MockSecretsState)(nil).GetSecretRevision), arg0, arg1)
}

// GetSecretValue mocks base method.
func (m *MockSecretsState) GetSecretValue(arg0 *secrets.URI, arg1 int) (secrets.SecretValue, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretsStateMockRecorder) GetSecretValue(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsState)(nil).GetSecretValue), arg0, arg1)
}

// ListSecretRevisions mocks base method.
func (m *MockSecretsState) ListSecretRevisions(arg0 *secrets.URI) ([]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretRevisions", arg0)
	ret0, _ := ret[0].([]*secrets.SecretRevisionMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretRevisions indicates an expected call of ListSecretRevisions.
func (mr *MockSecretsStateMockRecorder) ListSecretRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretRevisions", reflect.TypeOf((*MockSecretsState)(nil).ListSecretRevisions), arg0)
}

// ListSecrets mocks base method.
func (m *MockSecretsState) ListSecrets(arg0 state.SecretsFilter) ([]*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsStateMockRecorder) ListSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsState)(nil).ListSecrets), arg0)
}

// ListUnusedSecretRevisions mocks base method.
func (m *MockSecretsState) ListUnusedSecretRevisions(arg0 *secrets.URI) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnusedSecretRevisions", arg0)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUnusedSecretRevisions indicates an expected call of ListUnusedSecretRevisions.
func (mr *MockSecretsStateMockRecorder) ListUnusedSecretRevisions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnusedSecretRevisions", reflect.TypeOf((*MockSecretsState)(nil).ListUnusedSecretRevisions), arg0)
}

// UpdateSecret mocks base method.
func (m *MockSecretsState) UpdateSecret(arg0 *secrets.URI, arg1 state.UpdateSecretParams) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockSecretsStateMockRecorder) UpdateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretsState)(nil).UpdateSecret), arg0, arg1)
}

// MockSecretsConsumer is a mock of SecretsConsumer interface.
type MockSecretsConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsConsumerMockRecorder
}

// MockSecretsConsumerMockRecorder is the mock recorder for MockSecretsConsumer.
type MockSecretsConsumerMockRecorder struct {
	mock *MockSecretsConsumer
}

// NewMockSecretsConsumer creates a new mock instance.
func NewMockSecretsConsumer(ctrl *gomock.Controller) *MockSecretsConsumer {
	mock := &MockSecretsConsumer{ctrl: ctrl}
	mock.recorder = &MockSecretsConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsConsumer) EXPECT() *MockSecretsConsumerMockRecorder {
	return m.recorder
}

// GrantSecretAccess mocks base method.
func (m *MockSecretsConsumer) GrantSecretAccess(arg0 *secrets.URI, arg1 state.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantSecretAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantSecretAccess indicates an expected call of GrantSecretAccess.
func (mr *MockSecretsConsumerMockRecorder) GrantSecretAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantSecretAccess", reflect.TypeOf((*MockSecretsConsumer)(nil).GrantSecretAccess), arg0, arg1)
}

// RevokeSecretAccess mocks base method.
func (m *MockSecretsConsumer) RevokeSecretAccess(arg0 *secrets.URI, arg1 state.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecretAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeSecretAccess indicates an expected call of RevokeSecretAccess.
func (mr *MockSecretsConsumerMockRecorder) RevokeSecretAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecretAccess", reflect.TypeOf((*MockSecretsConsumer)(nil).RevokeSecretAccess), arg0, arg1)
}

// SecretAccess mocks base method.
func (m *MockSecretsConsumer) SecretAccess(arg0 *secrets.URI, arg1 names.Tag) (secrets.SecretRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretAccess", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretAccess indicates an expected call of SecretAccess.
func (mr *MockSecretsConsumerMockRecorder) SecretAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretAccess", reflect.TypeOf((*MockSecretsConsumer)(nil).SecretAccess), arg0, arg1)
}
