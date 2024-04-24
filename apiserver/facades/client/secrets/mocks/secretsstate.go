// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/secrets (interfaces: SecretService)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/secretsstate.go github.com/juju/juju/apiserver/facades/client/secrets SecretService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	secret "github.com/juju/juju/domain/secret"
	service "github.com/juju/juju/domain/secret/service"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretService is a mock of SecretService interface.
type MockSecretService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceMockRecorder
}

// MockSecretServiceMockRecorder is the mock recorder for MockSecretService.
type MockSecretServiceMockRecorder struct {
	mock *MockSecretService
}

// NewMockSecretService creates a new mock instance.
func NewMockSecretService(ctrl *gomock.Controller) *MockSecretService {
	mock := &MockSecretService{ctrl: ctrl}
	mock.recorder = &MockSecretServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretService) EXPECT() *MockSecretServiceMockRecorder {
	return m.recorder
}

// CreateSecret mocks base method.
func (m *MockSecretService) CreateSecret(arg0 context.Context, arg1 *secrets.URI, arg2 service.CreateSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecret indicates an expected call of CreateSecret.
func (mr *MockSecretServiceMockRecorder) CreateSecret(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockSecretService)(nil).CreateSecret), arg0, arg1, arg2)
}

// DeleteSecret mocks base method.
func (m *MockSecretService) DeleteSecret(arg0 context.Context, arg1 *secrets.URI, arg2 service.DeleteSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretServiceMockRecorder) DeleteSecret(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretService)(nil).DeleteSecret), arg0, arg1, arg2)
}

// GetSecret mocks base method.
func (m *MockSecretService) GetSecret(arg0 context.Context, arg1 *secrets.URI) (*secrets.SecretMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(*secrets.SecretMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockSecretServiceMockRecorder) GetSecret(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockSecretService)(nil).GetSecret), arg0, arg1)
}

// GetSecretGrants mocks base method.
func (m *MockSecretService) GetSecretGrants(arg0 context.Context, arg1 *secrets.URI, arg2 secrets.SecretRole) ([]service.SecretAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretGrants", arg0, arg1, arg2)
	ret0, _ := ret[0].([]service.SecretAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretGrants indicates an expected call of GetSecretGrants.
func (mr *MockSecretServiceMockRecorder) GetSecretGrants(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretGrants", reflect.TypeOf((*MockSecretService)(nil).GetSecretGrants), arg0, arg1, arg2)
}

// GetSecretValue mocks base method.
func (m *MockSecretService) GetSecretValue(arg0 context.Context, arg1 *secrets.URI, arg2 int, arg3 service.SecretAccessor) (secrets.SecretValue, *secrets.ValueRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretValue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(*secrets.ValueRef)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretServiceMockRecorder) GetSecretValue(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretService)(nil).GetSecretValue), arg0, arg1, arg2, arg3)
}

// GetUserSecretURIByLabel mocks base method.
func (m *MockSecretService) GetUserSecretURIByLabel(arg0 context.Context, arg1 string) (*secrets.URI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSecretURIByLabel", arg0, arg1)
	ret0, _ := ret[0].(*secrets.URI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSecretURIByLabel indicates an expected call of GetUserSecretURIByLabel.
func (mr *MockSecretServiceMockRecorder) GetUserSecretURIByLabel(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSecretURIByLabel", reflect.TypeOf((*MockSecretService)(nil).GetUserSecretURIByLabel), arg0, arg1)
}

// GrantSecretAccess mocks base method.
func (m *MockSecretService) GrantSecretAccess(arg0 context.Context, arg1 *secrets.URI, arg2 service.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantSecretAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantSecretAccess indicates an expected call of GrantSecretAccess.
func (mr *MockSecretServiceMockRecorder) GrantSecretAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantSecretAccess", reflect.TypeOf((*MockSecretService)(nil).GrantSecretAccess), arg0, arg1, arg2)
}

// ListCharmSecrets mocks base method.
func (m *MockSecretService) ListCharmSecrets(arg0 context.Context, arg1 ...service.CharmSecretOwner) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCharmSecrets", varargs...)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCharmSecrets indicates an expected call of ListCharmSecrets.
func (mr *MockSecretServiceMockRecorder) ListCharmSecrets(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmSecrets", reflect.TypeOf((*MockSecretService)(nil).ListCharmSecrets), varargs...)
}

// ListSecrets mocks base method.
func (m *MockSecretService) ListSecrets(arg0 context.Context, arg1 *secrets.URI, arg2 *int, arg3 secret.Labels) ([]*secrets.SecretMetadata, [][]*secrets.SecretRevisionMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*secrets.SecretMetadata)
	ret1, _ := ret[1].([][]*secrets.SecretRevisionMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretServiceMockRecorder) ListSecrets(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretService)(nil).ListSecrets), arg0, arg1, arg2, arg3)
}

// RevokeSecretAccess mocks base method.
func (m *MockSecretService) RevokeSecretAccess(arg0 context.Context, arg1 *secrets.URI, arg2 service.SecretAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSecretAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeSecretAccess indicates an expected call of RevokeSecretAccess.
func (mr *MockSecretServiceMockRecorder) RevokeSecretAccess(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSecretAccess", reflect.TypeOf((*MockSecretService)(nil).RevokeSecretAccess), arg0, arg1, arg2)
}

// UpdateSecret mocks base method.
func (m *MockSecretService) UpdateSecret(arg0 context.Context, arg1 *secrets.URI, arg2 service.UpdateSecretParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret.
func (mr *MockSecretServiceMockRecorder) UpdateSecret(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockSecretService)(nil).UpdateSecret), arg0, arg1, arg2)
}
