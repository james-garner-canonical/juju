// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/deployer (interfaces: ModelCommand,ConsumeDetails,CharmDeployAPI)
//
// Generated by this command:
//
//	mockgen -typed -package mocks -destination mocks/deployer_mock.go github.com/juju/juju/cmd/juju/application/deployer ModelCommand,ConsumeDetails,CharmDeployAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	httpbakery "github.com/go-macaroon-bakery/macaroon-bakery/v3/httpbakery"
	charms "github.com/juju/juju/api/common/charms"
	modelcmd "github.com/juju/juju/cmd/modelcmd"
	model "github.com/juju/juju/core/model"
	jujuclient "github.com/juju/juju/jujuclient"
	params "github.com/juju/juju/rpc/params"
	gomock "go.uber.org/mock/gomock"
)

// MockModelCommand is a mock of ModelCommand interface.
type MockModelCommand struct {
	ctrl     *gomock.Controller
	recorder *MockModelCommandMockRecorder
}

// MockModelCommandMockRecorder is the mock recorder for MockModelCommand.
type MockModelCommandMockRecorder struct {
	mock *MockModelCommand
}

// NewMockModelCommand creates a new mock instance.
func NewMockModelCommand(ctrl *gomock.Controller) *MockModelCommand {
	mock := &MockModelCommand{ctrl: ctrl}
	mock.recorder = &MockModelCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelCommand) EXPECT() *MockModelCommandMockRecorder {
	return m.recorder
}

// BakeryClient mocks base method.
func (m *MockModelCommand) BakeryClient() (*httpbakery.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(*httpbakery.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockModelCommandMockRecorder) BakeryClient() *MockModelCommandBakeryClientCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockModelCommand)(nil).BakeryClient))
	return &MockModelCommandBakeryClientCall{Call: call}
}

// MockModelCommandBakeryClientCall wrap *gomock.Call
type MockModelCommandBakeryClientCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelCommandBakeryClientCall) Return(arg0 *httpbakery.Client, arg1 error) *MockModelCommandBakeryClientCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelCommandBakeryClientCall) Do(f func() (*httpbakery.Client, error)) *MockModelCommandBakeryClientCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelCommandBakeryClientCall) DoAndReturn(f func() (*httpbakery.Client, error)) *MockModelCommandBakeryClientCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ControllerName mocks base method.
func (m *MockModelCommand) ControllerName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControllerName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ControllerName indicates an expected call of ControllerName.
func (mr *MockModelCommandMockRecorder) ControllerName() *MockModelCommandControllerNameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControllerName", reflect.TypeOf((*MockModelCommand)(nil).ControllerName))
	return &MockModelCommandControllerNameCall{Call: call}
}

// MockModelCommandControllerNameCall wrap *gomock.Call
type MockModelCommandControllerNameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelCommandControllerNameCall) Return(arg0 string, arg1 error) *MockModelCommandControllerNameCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelCommandControllerNameCall) Do(f func() (string, error)) *MockModelCommandControllerNameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelCommandControllerNameCall) DoAndReturn(f func() (string, error)) *MockModelCommandControllerNameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CurrentAccountDetails mocks base method.
func (m *MockModelCommand) CurrentAccountDetails() (*jujuclient.AccountDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentAccountDetails")
	ret0, _ := ret[0].(*jujuclient.AccountDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentAccountDetails indicates an expected call of CurrentAccountDetails.
func (mr *MockModelCommandMockRecorder) CurrentAccountDetails() *MockModelCommandCurrentAccountDetailsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentAccountDetails", reflect.TypeOf((*MockModelCommand)(nil).CurrentAccountDetails))
	return &MockModelCommandCurrentAccountDetailsCall{Call: call}
}

// MockModelCommandCurrentAccountDetailsCall wrap *gomock.Call
type MockModelCommandCurrentAccountDetailsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelCommandCurrentAccountDetailsCall) Return(arg0 *jujuclient.AccountDetails, arg1 error) *MockModelCommandCurrentAccountDetailsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelCommandCurrentAccountDetailsCall) Do(f func() (*jujuclient.AccountDetails, error)) *MockModelCommandCurrentAccountDetailsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelCommandCurrentAccountDetailsCall) DoAndReturn(f func() (*jujuclient.AccountDetails, error)) *MockModelCommandCurrentAccountDetailsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filesystem mocks base method.
func (m *MockModelCommand) Filesystem() modelcmd.Filesystem {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filesystem")
	ret0, _ := ret[0].(modelcmd.Filesystem)
	return ret0
}

// Filesystem indicates an expected call of Filesystem.
func (mr *MockModelCommandMockRecorder) Filesystem() *MockModelCommandFilesystemCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filesystem", reflect.TypeOf((*MockModelCommand)(nil).Filesystem))
	return &MockModelCommandFilesystemCall{Call: call}
}

// MockModelCommandFilesystemCall wrap *gomock.Call
type MockModelCommandFilesystemCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelCommandFilesystemCall) Return(arg0 modelcmd.Filesystem) *MockModelCommandFilesystemCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelCommandFilesystemCall) Do(f func() modelcmd.Filesystem) *MockModelCommandFilesystemCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelCommandFilesystemCall) DoAndReturn(f func() modelcmd.Filesystem) *MockModelCommandFilesystemCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelDetails mocks base method.
func (m *MockModelCommand) ModelDetails() (string, *jujuclient.ModelDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelDetails")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*jujuclient.ModelDetails)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ModelDetails indicates an expected call of ModelDetails.
func (mr *MockModelCommandMockRecorder) ModelDetails() *MockModelCommandModelDetailsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelDetails", reflect.TypeOf((*MockModelCommand)(nil).ModelDetails))
	return &MockModelCommandModelDetailsCall{Call: call}
}

// MockModelCommandModelDetailsCall wrap *gomock.Call
type MockModelCommandModelDetailsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelCommandModelDetailsCall) Return(arg0 string, arg1 *jujuclient.ModelDetails, arg2 error) *MockModelCommandModelDetailsCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelCommandModelDetailsCall) Do(f func() (string, *jujuclient.ModelDetails, error)) *MockModelCommandModelDetailsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelCommandModelDetailsCall) DoAndReturn(f func() (string, *jujuclient.ModelDetails, error)) *MockModelCommandModelDetailsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelType mocks base method.
func (m *MockModelCommand) ModelType() (model.ModelType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelType")
	ret0, _ := ret[0].(model.ModelType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelType indicates an expected call of ModelType.
func (mr *MockModelCommandMockRecorder) ModelType() *MockModelCommandModelTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelType", reflect.TypeOf((*MockModelCommand)(nil).ModelType))
	return &MockModelCommandModelTypeCall{Call: call}
}

// MockModelCommandModelTypeCall wrap *gomock.Call
type MockModelCommandModelTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockModelCommandModelTypeCall) Return(arg0 model.ModelType, arg1 error) *MockModelCommandModelTypeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockModelCommandModelTypeCall) Do(f func() (model.ModelType, error)) *MockModelCommandModelTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockModelCommandModelTypeCall) DoAndReturn(f func() (model.ModelType, error)) *MockModelCommandModelTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockConsumeDetails is a mock of ConsumeDetails interface.
type MockConsumeDetails struct {
	ctrl     *gomock.Controller
	recorder *MockConsumeDetailsMockRecorder
}

// MockConsumeDetailsMockRecorder is the mock recorder for MockConsumeDetails.
type MockConsumeDetailsMockRecorder struct {
	mock *MockConsumeDetails
}

// NewMockConsumeDetails creates a new mock instance.
func NewMockConsumeDetails(ctrl *gomock.Controller) *MockConsumeDetails {
	mock := &MockConsumeDetails{ctrl: ctrl}
	mock.recorder = &MockConsumeDetailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumeDetails) EXPECT() *MockConsumeDetailsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockConsumeDetails) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConsumeDetailsMockRecorder) Close() *MockConsumeDetailsCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConsumeDetails)(nil).Close))
	return &MockConsumeDetailsCloseCall{Call: call}
}

// MockConsumeDetailsCloseCall wrap *gomock.Call
type MockConsumeDetailsCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConsumeDetailsCloseCall) Return(arg0 error) *MockConsumeDetailsCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConsumeDetailsCloseCall) Do(f func() error) *MockConsumeDetailsCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConsumeDetailsCloseCall) DoAndReturn(f func() error) *MockConsumeDetailsCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetConsumeDetails mocks base method.
func (m *MockConsumeDetails) GetConsumeDetails(arg0 string) (params.ConsumeOfferDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumeDetails", arg0)
	ret0, _ := ret[0].(params.ConsumeOfferDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumeDetails indicates an expected call of GetConsumeDetails.
func (mr *MockConsumeDetailsMockRecorder) GetConsumeDetails(arg0 any) *MockConsumeDetailsGetConsumeDetailsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumeDetails", reflect.TypeOf((*MockConsumeDetails)(nil).GetConsumeDetails), arg0)
	return &MockConsumeDetailsGetConsumeDetailsCall{Call: call}
}

// MockConsumeDetailsGetConsumeDetailsCall wrap *gomock.Call
type MockConsumeDetailsGetConsumeDetailsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockConsumeDetailsGetConsumeDetailsCall) Return(arg0 params.ConsumeOfferDetails, arg1 error) *MockConsumeDetailsGetConsumeDetailsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockConsumeDetailsGetConsumeDetailsCall) Do(f func(string) (params.ConsumeOfferDetails, error)) *MockConsumeDetailsGetConsumeDetailsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockConsumeDetailsGetConsumeDetailsCall) DoAndReturn(f func(string) (params.ConsumeOfferDetails, error)) *MockConsumeDetailsGetConsumeDetailsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockCharmDeployAPI is a mock of CharmDeployAPI interface.
type MockCharmDeployAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCharmDeployAPIMockRecorder
}

// MockCharmDeployAPIMockRecorder is the mock recorder for MockCharmDeployAPI.
type MockCharmDeployAPIMockRecorder struct {
	mock *MockCharmDeployAPI
}

// NewMockCharmDeployAPI creates a new mock instance.
func NewMockCharmDeployAPI(ctrl *gomock.Controller) *MockCharmDeployAPI {
	mock := &MockCharmDeployAPI{ctrl: ctrl}
	mock.recorder = &MockCharmDeployAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmDeployAPI) EXPECT() *MockCharmDeployAPIMockRecorder {
	return m.recorder
}

// CharmInfo mocks base method.
func (m *MockCharmDeployAPI) CharmInfo(arg0 context.Context, arg1 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0, arg1)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo.
func (mr *MockCharmDeployAPIMockRecorder) CharmInfo(arg0, arg1 any) *MockCharmDeployAPICharmInfoCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockCharmDeployAPI)(nil).CharmInfo), arg0, arg1)
	return &MockCharmDeployAPICharmInfoCall{Call: call}
}

// MockCharmDeployAPICharmInfoCall wrap *gomock.Call
type MockCharmDeployAPICharmInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmDeployAPICharmInfoCall) Return(arg0 *charms.CharmInfo, arg1 error) *MockCharmDeployAPICharmInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmDeployAPICharmInfoCall) Do(f func(context.Context, string) (*charms.CharmInfo, error)) *MockCharmDeployAPICharmInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmDeployAPICharmInfoCall) DoAndReturn(f func(context.Context, string) (*charms.CharmInfo, error)) *MockCharmDeployAPICharmInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ModelGet mocks base method.
func (m *MockCharmDeployAPI) ModelGet() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet.
func (mr *MockCharmDeployAPIMockRecorder) ModelGet() *MockCharmDeployAPIModelGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockCharmDeployAPI)(nil).ModelGet))
	return &MockCharmDeployAPIModelGetCall{Call: call}
}

// MockCharmDeployAPIModelGetCall wrap *gomock.Call
type MockCharmDeployAPIModelGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockCharmDeployAPIModelGetCall) Return(arg0 map[string]any, arg1 error) *MockCharmDeployAPIModelGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockCharmDeployAPIModelGetCall) Do(f func() (map[string]any, error)) *MockCharmDeployAPIModelGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockCharmDeployAPIModelGetCall) DoAndReturn(f func() (map[string]any, error)) *MockCharmDeployAPIModelGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
