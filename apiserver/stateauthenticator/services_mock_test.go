// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/stateauthenticator (interfaces: ControllerConfigService,AccessService,MacaroonService,AgentAuthenticatorFactory)
//
// Generated by this command:
//
//	mockgen -typed -package stateauthenticator -destination services_mock_test.go github.com/juju/juju/apiserver/stateauthenticator ControllerConfigService,AccessService,MacaroonService,AgentAuthenticatorFactory
//

// Package stateauthenticator is a generated GoMock package.
package stateauthenticator

import (
	context "context"
	reflect "reflect"
	time "time"

	bakery "github.com/go-macaroon-bakery/macaroon-bakery/v3/bakery"
	dbrootkeystore "github.com/go-macaroon-bakery/macaroon-bakery/v3/bakery/dbrootkeystore"
	authentication "github.com/juju/juju/apiserver/authentication"
	controller "github.com/juju/juju/controller"
	model "github.com/juju/juju/core/model"
	permission "github.com/juju/juju/core/permission"
	user "github.com/juju/juju/core/user"
	auth "github.com/juju/juju/internal/auth"
	state "github.com/juju/juju/state"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerConfigService is a mock of ControllerConfigService interface.
type MockControllerConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockControllerConfigServiceMockRecorder
}

// MockControllerConfigServiceMockRecorder is the mock recorder for MockControllerConfigService.
type MockControllerConfigServiceMockRecorder struct {
	mock *MockControllerConfigService
}

// NewMockControllerConfigService creates a new mock instance.
func NewMockControllerConfigService(ctrl *gomock.Controller) *MockControllerConfigService {
	mock := &MockControllerConfigService{ctrl: ctrl}
	mock.recorder = &MockControllerConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerConfigService) EXPECT() *MockControllerConfigServiceMockRecorder {
	return m.recorder
}

// ControllerConfig mocks base method.
func (m *MockControllerConfigService) ControllerConfig(arg0 context.Context) (controller.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerConfig", arg0)
	ret0, _ := ret[0].(controller.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerConfig indicates an expected call of ControllerConfig.
func (mr *MockControllerConfigServiceMockRecorder) ControllerConfig(arg0 any) *MockControllerConfigServiceControllerConfigCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerConfig", reflect.TypeOf((*MockControllerConfigService)(nil).ControllerConfig), arg0)
	return &MockControllerConfigServiceControllerConfigCall{Call: call}
}

// MockControllerConfigServiceControllerConfigCall wrap *gomock.Call
type MockControllerConfigServiceControllerConfigCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockControllerConfigServiceControllerConfigCall) Return(arg0 controller.Config, arg1 error) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockControllerConfigServiceControllerConfigCall) Do(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockControllerConfigServiceControllerConfigCall) DoAndReturn(f func(context.Context) (controller.Config, error)) *MockControllerConfigServiceControllerConfigCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockAccessService is a mock of AccessService interface.
type MockAccessService struct {
	ctrl     *gomock.Controller
	recorder *MockAccessServiceMockRecorder
}

// MockAccessServiceMockRecorder is the mock recorder for MockAccessService.
type MockAccessServiceMockRecorder struct {
	mock *MockAccessService
}

// NewMockAccessService creates a new mock instance.
func NewMockAccessService(ctrl *gomock.Controller) *MockAccessService {
	mock := &MockAccessService{ctrl: ctrl}
	mock.recorder = &MockAccessServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessService) EXPECT() *MockAccessServiceMockRecorder {
	return m.recorder
}

// GetUserByAuth mocks base method.
func (m *MockAccessService) GetUserByAuth(arg0 context.Context, arg1 string, arg2 auth.Password) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuth indicates an expected call of GetUserByAuth.
func (mr *MockAccessServiceMockRecorder) GetUserByAuth(arg0, arg1, arg2 any) *MockAccessServiceGetUserByAuthCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuth", reflect.TypeOf((*MockAccessService)(nil).GetUserByAuth), arg0, arg1, arg2)
	return &MockAccessServiceGetUserByAuthCall{Call: call}
}

// MockAccessServiceGetUserByAuthCall wrap *gomock.Call
type MockAccessServiceGetUserByAuthCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessServiceGetUserByAuthCall) Return(arg0 user.User, arg1 error) *MockAccessServiceGetUserByAuthCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessServiceGetUserByAuthCall) Do(f func(context.Context, string, auth.Password) (user.User, error)) *MockAccessServiceGetUserByAuthCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessServiceGetUserByAuthCall) DoAndReturn(f func(context.Context, string, auth.Password) (user.User, error)) *MockAccessServiceGetUserByAuthCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUserByName mocks base method.
func (m *MockAccessService) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockAccessServiceMockRecorder) GetUserByName(arg0, arg1 any) *MockAccessServiceGetUserByNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockAccessService)(nil).GetUserByName), arg0, arg1)
	return &MockAccessServiceGetUserByNameCall{Call: call}
}

// MockAccessServiceGetUserByNameCall wrap *gomock.Call
type MockAccessServiceGetUserByNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessServiceGetUserByNameCall) Return(arg0 user.User, arg1 error) *MockAccessServiceGetUserByNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessServiceGetUserByNameCall) Do(f func(context.Context, string) (user.User, error)) *MockAccessServiceGetUserByNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessServiceGetUserByNameCall) DoAndReturn(f func(context.Context, string) (user.User, error)) *MockAccessServiceGetUserByNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadUserAccessLevelForTargetAddingMissingUser mocks base method.
func (m *MockAccessService) ReadUserAccessLevelForTargetAddingMissingUser(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessLevelForTargetAddingMissingUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessLevelForTargetAddingMissingUser indicates an expected call of ReadUserAccessLevelForTargetAddingMissingUser.
func (mr *MockAccessServiceMockRecorder) ReadUserAccessLevelForTargetAddingMissingUser(arg0, arg1, arg2 any) *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessLevelForTargetAddingMissingUser", reflect.TypeOf((*MockAccessService)(nil).ReadUserAccessLevelForTargetAddingMissingUser), arg0, arg1, arg2)
	return &MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall{Call: call}
}

// MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall wrap *gomock.Call
type MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall) Return(arg0 permission.Access, arg1 error) *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall) Do(f func(context.Context, string, permission.ID) (permission.Access, error)) *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall) DoAndReturn(f func(context.Context, string, permission.ID) (permission.Access, error)) *MockAccessServiceReadUserAccessLevelForTargetAddingMissingUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateLastModelLogin mocks base method.
func (m *MockAccessService) UpdateLastModelLogin(arg0 context.Context, arg1 string, arg2 model.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastModelLogin", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastModelLogin indicates an expected call of UpdateLastModelLogin.
func (mr *MockAccessServiceMockRecorder) UpdateLastModelLogin(arg0, arg1, arg2 any) *MockAccessServiceUpdateLastModelLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastModelLogin", reflect.TypeOf((*MockAccessService)(nil).UpdateLastModelLogin), arg0, arg1, arg2)
	return &MockAccessServiceUpdateLastModelLoginCall{Call: call}
}

// MockAccessServiceUpdateLastModelLoginCall wrap *gomock.Call
type MockAccessServiceUpdateLastModelLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAccessServiceUpdateLastModelLoginCall) Return(arg0 error) *MockAccessServiceUpdateLastModelLoginCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAccessServiceUpdateLastModelLoginCall) Do(f func(context.Context, string, model.UUID) error) *MockAccessServiceUpdateLastModelLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAccessServiceUpdateLastModelLoginCall) DoAndReturn(f func(context.Context, string, model.UUID) error) *MockAccessServiceUpdateLastModelLoginCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockMacaroonService is a mock of MacaroonService interface.
type MockMacaroonService struct {
	ctrl     *gomock.Controller
	recorder *MockMacaroonServiceMockRecorder
}

// MockMacaroonServiceMockRecorder is the mock recorder for MockMacaroonService.
type MockMacaroonServiceMockRecorder struct {
	mock *MockMacaroonService
}

// NewMockMacaroonService creates a new mock instance.
func NewMockMacaroonService(ctrl *gomock.Controller) *MockMacaroonService {
	mock := &MockMacaroonService{ctrl: ctrl}
	mock.recorder = &MockMacaroonServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMacaroonService) EXPECT() *MockMacaroonServiceMockRecorder {
	return m.recorder
}

// FindLatestKeyContext mocks base method.
func (m *MockMacaroonService) FindLatestKeyContext(arg0 context.Context, arg1, arg2, arg3 time.Time) (dbrootkeystore.RootKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLatestKeyContext", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(dbrootkeystore.RootKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLatestKeyContext indicates an expected call of FindLatestKeyContext.
func (mr *MockMacaroonServiceMockRecorder) FindLatestKeyContext(arg0, arg1, arg2, arg3 any) *MockMacaroonServiceFindLatestKeyContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLatestKeyContext", reflect.TypeOf((*MockMacaroonService)(nil).FindLatestKeyContext), arg0, arg1, arg2, arg3)
	return &MockMacaroonServiceFindLatestKeyContextCall{Call: call}
}

// MockMacaroonServiceFindLatestKeyContextCall wrap *gomock.Call
type MockMacaroonServiceFindLatestKeyContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMacaroonServiceFindLatestKeyContextCall) Return(arg0 dbrootkeystore.RootKey, arg1 error) *MockMacaroonServiceFindLatestKeyContextCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMacaroonServiceFindLatestKeyContextCall) Do(f func(context.Context, time.Time, time.Time, time.Time) (dbrootkeystore.RootKey, error)) *MockMacaroonServiceFindLatestKeyContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMacaroonServiceFindLatestKeyContextCall) DoAndReturn(f func(context.Context, time.Time, time.Time, time.Time) (dbrootkeystore.RootKey, error)) *MockMacaroonServiceFindLatestKeyContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetExternalUsersThirdPartyKey mocks base method.
func (m *MockMacaroonService) GetExternalUsersThirdPartyKey(arg0 context.Context) (*bakery.KeyPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalUsersThirdPartyKey", arg0)
	ret0, _ := ret[0].(*bakery.KeyPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalUsersThirdPartyKey indicates an expected call of GetExternalUsersThirdPartyKey.
func (mr *MockMacaroonServiceMockRecorder) GetExternalUsersThirdPartyKey(arg0 any) *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalUsersThirdPartyKey", reflect.TypeOf((*MockMacaroonService)(nil).GetExternalUsersThirdPartyKey), arg0)
	return &MockMacaroonServiceGetExternalUsersThirdPartyKeyCall{Call: call}
}

// MockMacaroonServiceGetExternalUsersThirdPartyKeyCall wrap *gomock.Call
type MockMacaroonServiceGetExternalUsersThirdPartyKeyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall) Return(arg0 *bakery.KeyPair, arg1 error) *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall) Do(f func(context.Context) (*bakery.KeyPair, error)) *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall) DoAndReturn(f func(context.Context) (*bakery.KeyPair, error)) *MockMacaroonServiceGetExternalUsersThirdPartyKeyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetKeyContext mocks base method.
func (m *MockMacaroonService) GetKeyContext(arg0 context.Context, arg1 []byte) (dbrootkeystore.RootKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyContext", arg0, arg1)
	ret0, _ := ret[0].(dbrootkeystore.RootKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyContext indicates an expected call of GetKeyContext.
func (mr *MockMacaroonServiceMockRecorder) GetKeyContext(arg0, arg1 any) *MockMacaroonServiceGetKeyContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyContext", reflect.TypeOf((*MockMacaroonService)(nil).GetKeyContext), arg0, arg1)
	return &MockMacaroonServiceGetKeyContextCall{Call: call}
}

// MockMacaroonServiceGetKeyContextCall wrap *gomock.Call
type MockMacaroonServiceGetKeyContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMacaroonServiceGetKeyContextCall) Return(arg0 dbrootkeystore.RootKey, arg1 error) *MockMacaroonServiceGetKeyContextCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMacaroonServiceGetKeyContextCall) Do(f func(context.Context, []byte) (dbrootkeystore.RootKey, error)) *MockMacaroonServiceGetKeyContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMacaroonServiceGetKeyContextCall) DoAndReturn(f func(context.Context, []byte) (dbrootkeystore.RootKey, error)) *MockMacaroonServiceGetKeyContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLocalUsersKey mocks base method.
func (m *MockMacaroonService) GetLocalUsersKey(arg0 context.Context) (*bakery.KeyPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalUsersKey", arg0)
	ret0, _ := ret[0].(*bakery.KeyPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocalUsersKey indicates an expected call of GetLocalUsersKey.
func (mr *MockMacaroonServiceMockRecorder) GetLocalUsersKey(arg0 any) *MockMacaroonServiceGetLocalUsersKeyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalUsersKey", reflect.TypeOf((*MockMacaroonService)(nil).GetLocalUsersKey), arg0)
	return &MockMacaroonServiceGetLocalUsersKeyCall{Call: call}
}

// MockMacaroonServiceGetLocalUsersKeyCall wrap *gomock.Call
type MockMacaroonServiceGetLocalUsersKeyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMacaroonServiceGetLocalUsersKeyCall) Return(arg0 *bakery.KeyPair, arg1 error) *MockMacaroonServiceGetLocalUsersKeyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMacaroonServiceGetLocalUsersKeyCall) Do(f func(context.Context) (*bakery.KeyPair, error)) *MockMacaroonServiceGetLocalUsersKeyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMacaroonServiceGetLocalUsersKeyCall) DoAndReturn(f func(context.Context) (*bakery.KeyPair, error)) *MockMacaroonServiceGetLocalUsersKeyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetLocalUsersThirdPartyKey mocks base method.
func (m *MockMacaroonService) GetLocalUsersThirdPartyKey(arg0 context.Context) (*bakery.KeyPair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocalUsersThirdPartyKey", arg0)
	ret0, _ := ret[0].(*bakery.KeyPair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocalUsersThirdPartyKey indicates an expected call of GetLocalUsersThirdPartyKey.
func (mr *MockMacaroonServiceMockRecorder) GetLocalUsersThirdPartyKey(arg0 any) *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocalUsersThirdPartyKey", reflect.TypeOf((*MockMacaroonService)(nil).GetLocalUsersThirdPartyKey), arg0)
	return &MockMacaroonServiceGetLocalUsersThirdPartyKeyCall{Call: call}
}

// MockMacaroonServiceGetLocalUsersThirdPartyKeyCall wrap *gomock.Call
type MockMacaroonServiceGetLocalUsersThirdPartyKeyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall) Return(arg0 *bakery.KeyPair, arg1 error) *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall) Do(f func(context.Context) (*bakery.KeyPair, error)) *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall) DoAndReturn(f func(context.Context) (*bakery.KeyPair, error)) *MockMacaroonServiceGetLocalUsersThirdPartyKeyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InsertKeyContext mocks base method.
func (m *MockMacaroonService) InsertKeyContext(arg0 context.Context, arg1 dbrootkeystore.RootKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertKeyContext", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertKeyContext indicates an expected call of InsertKeyContext.
func (mr *MockMacaroonServiceMockRecorder) InsertKeyContext(arg0, arg1 any) *MockMacaroonServiceInsertKeyContextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertKeyContext", reflect.TypeOf((*MockMacaroonService)(nil).InsertKeyContext), arg0, arg1)
	return &MockMacaroonServiceInsertKeyContextCall{Call: call}
}

// MockMacaroonServiceInsertKeyContextCall wrap *gomock.Call
type MockMacaroonServiceInsertKeyContextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockMacaroonServiceInsertKeyContextCall) Return(arg0 error) *MockMacaroonServiceInsertKeyContextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockMacaroonServiceInsertKeyContextCall) Do(f func(context.Context, dbrootkeystore.RootKey) error) *MockMacaroonServiceInsertKeyContextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockMacaroonServiceInsertKeyContextCall) DoAndReturn(f func(context.Context, dbrootkeystore.RootKey) error) *MockMacaroonServiceInsertKeyContextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockAgentAuthenticatorFactory is a mock of AgentAuthenticatorFactory interface.
type MockAgentAuthenticatorFactory struct {
	ctrl     *gomock.Controller
	recorder *MockAgentAuthenticatorFactoryMockRecorder
}

// MockAgentAuthenticatorFactoryMockRecorder is the mock recorder for MockAgentAuthenticatorFactory.
type MockAgentAuthenticatorFactoryMockRecorder struct {
	mock *MockAgentAuthenticatorFactory
}

// NewMockAgentAuthenticatorFactory creates a new mock instance.
func NewMockAgentAuthenticatorFactory(ctrl *gomock.Controller) *MockAgentAuthenticatorFactory {
	mock := &MockAgentAuthenticatorFactory{ctrl: ctrl}
	mock.recorder = &MockAgentAuthenticatorFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentAuthenticatorFactory) EXPECT() *MockAgentAuthenticatorFactoryMockRecorder {
	return m.recorder
}

// Authenticator mocks base method.
func (m *MockAgentAuthenticatorFactory) Authenticator() authentication.EntityAuthenticator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticator")
	ret0, _ := ret[0].(authentication.EntityAuthenticator)
	return ret0
}

// Authenticator indicates an expected call of Authenticator.
func (mr *MockAgentAuthenticatorFactoryMockRecorder) Authenticator() *MockAgentAuthenticatorFactoryAuthenticatorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticator", reflect.TypeOf((*MockAgentAuthenticatorFactory)(nil).Authenticator))
	return &MockAgentAuthenticatorFactoryAuthenticatorCall{Call: call}
}

// MockAgentAuthenticatorFactoryAuthenticatorCall wrap *gomock.Call
type MockAgentAuthenticatorFactoryAuthenticatorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAgentAuthenticatorFactoryAuthenticatorCall) Return(arg0 authentication.EntityAuthenticator) *MockAgentAuthenticatorFactoryAuthenticatorCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAgentAuthenticatorFactoryAuthenticatorCall) Do(f func() authentication.EntityAuthenticator) *MockAgentAuthenticatorFactoryAuthenticatorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAgentAuthenticatorFactoryAuthenticatorCall) DoAndReturn(f func() authentication.EntityAuthenticator) *MockAgentAuthenticatorFactoryAuthenticatorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// AuthenticatorForState mocks base method.
func (m *MockAgentAuthenticatorFactory) AuthenticatorForState(arg0 *state.State) authentication.EntityAuthenticator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticatorForState", arg0)
	ret0, _ := ret[0].(authentication.EntityAuthenticator)
	return ret0
}

// AuthenticatorForState indicates an expected call of AuthenticatorForState.
func (mr *MockAgentAuthenticatorFactoryMockRecorder) AuthenticatorForState(arg0 any) *MockAgentAuthenticatorFactoryAuthenticatorForStateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticatorForState", reflect.TypeOf((*MockAgentAuthenticatorFactory)(nil).AuthenticatorForState), arg0)
	return &MockAgentAuthenticatorFactoryAuthenticatorForStateCall{Call: call}
}

// MockAgentAuthenticatorFactoryAuthenticatorForStateCall wrap *gomock.Call
type MockAgentAuthenticatorFactoryAuthenticatorForStateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockAgentAuthenticatorFactoryAuthenticatorForStateCall) Return(arg0 authentication.EntityAuthenticator) *MockAgentAuthenticatorFactoryAuthenticatorForStateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockAgentAuthenticatorFactoryAuthenticatorForStateCall) Do(f func(*state.State) authentication.EntityAuthenticator) *MockAgentAuthenticatorFactoryAuthenticatorForStateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockAgentAuthenticatorFactoryAuthenticatorForStateCall) DoAndReturn(f func(*state.State) authentication.EntityAuthenticator) *MockAgentAuthenticatorFactoryAuthenticatorForStateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
