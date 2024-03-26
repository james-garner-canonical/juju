// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/deployer (interfaces: DeployerAPI,CharmReader,DeployConfigFlag)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/deploy_mock.go github.com/juju/juju/cmd/juju/application/deployer DeployerAPI,CharmReader,DeployConfigFlag
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	charm "github.com/juju/charm/v13"
	resource "github.com/juju/charm/v13/resource"
	cmd "github.com/juju/cmd/v4"
	api "github.com/juju/juju/api"
	base "github.com/juju/juju/api/base"
	application "github.com/juju/juju/api/client/application"
	client "github.com/juju/juju/api/client/client"
	charm0 "github.com/juju/juju/api/common/charm"
	charms "github.com/juju/juju/api/common/charms"
	constraints "github.com/juju/juju/core/constraints"
	crossmodel "github.com/juju/juju/core/crossmodel"
	params "github.com/juju/juju/rpc/params"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
	httprequest "gopkg.in/httprequest.v1"
)

// MockDeployerAPI is a mock of DeployerAPI interface.
type MockDeployerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerAPIMockRecorder
}

// MockDeployerAPIMockRecorder is the mock recorder for MockDeployerAPI.
type MockDeployerAPIMockRecorder struct {
	mock *MockDeployerAPI
}

// NewMockDeployerAPI creates a new mock instance.
func NewMockDeployerAPI(ctrl *gomock.Controller) *MockDeployerAPI {
	mock := &MockDeployerAPI{ctrl: ctrl}
	mock.recorder = &MockDeployerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployerAPI) EXPECT() *MockDeployerAPIMockRecorder {
	return m.recorder
}

// APICall mocks base method.
func (m *MockDeployerAPI) APICall(arg0 context.Context, arg1 string, arg2 int, arg3, arg4 string, arg5, arg6 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall.
func (mr *MockDeployerAPIMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockDeployerAPI)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AddCharm mocks base method.
func (m *MockDeployerAPI) AddCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 bool) (charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm0.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharm indicates an expected call of AddCharm.
func (mr *MockDeployerAPIMockRecorder) AddCharm(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharm", reflect.TypeOf((*MockDeployerAPI)(nil).AddCharm), arg0, arg1, arg2)
}

// AddLocalCharm mocks base method.
func (m *MockDeployerAPI) AddLocalCharm(arg0 *charm.URL, arg1 charm.Charm, arg2 bool) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocalCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocalCharm indicates an expected call of AddLocalCharm.
func (mr *MockDeployerAPIMockRecorder) AddLocalCharm(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocalCharm", reflect.TypeOf((*MockDeployerAPI)(nil).AddLocalCharm), arg0, arg1, arg2)
}

// AddMachines mocks base method.
func (m *MockDeployerAPI) AddMachines(arg0 []params.AddMachineParams) ([]params.AddMachinesResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMachines", arg0)
	ret0, _ := ret[0].([]params.AddMachinesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMachines indicates an expected call of AddMachines.
func (mr *MockDeployerAPIMockRecorder) AddMachines(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMachines", reflect.TypeOf((*MockDeployerAPI)(nil).AddMachines), arg0)
}

// AddRelation mocks base method.
func (m *MockDeployerAPI) AddRelation(arg0, arg1 []string) (*params.AddRelationResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRelation", arg0, arg1)
	ret0, _ := ret[0].(*params.AddRelationResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRelation indicates an expected call of AddRelation.
func (mr *MockDeployerAPIMockRecorder) AddRelation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRelation", reflect.TypeOf((*MockDeployerAPI)(nil).AddRelation), arg0, arg1)
}

// AddUnits mocks base method.
func (m *MockDeployerAPI) AddUnits(arg0 application.AddUnitsParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUnits", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUnits indicates an expected call of AddUnits.
func (mr *MockDeployerAPIMockRecorder) AddUnits(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnits", reflect.TypeOf((*MockDeployerAPI)(nil).AddUnits), arg0)
}

// ApplicationsInfo mocks base method.
func (m *MockDeployerAPI) ApplicationsInfo(arg0 []names.ApplicationTag) ([]params.ApplicationInfoResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationsInfo", arg0)
	ret0, _ := ret[0].([]params.ApplicationInfoResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationsInfo indicates an expected call of ApplicationsInfo.
func (mr *MockDeployerAPIMockRecorder) ApplicationsInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationsInfo", reflect.TypeOf((*MockDeployerAPI)(nil).ApplicationsInfo), arg0)
}

// BakeryClient mocks base method.
func (m *MockDeployerAPI) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockDeployerAPIMockRecorder) BakeryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockDeployerAPI)(nil).BakeryClient))
}

// BestFacadeVersion mocks base method.
func (m *MockDeployerAPI) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion.
func (mr *MockDeployerAPIMockRecorder) BestFacadeVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockDeployerAPI)(nil).BestFacadeVersion), arg0)
}

// CharmInfo mocks base method.
func (m *MockDeployerAPI) CharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo.
func (mr *MockDeployerAPIMockRecorder) CharmInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockDeployerAPI)(nil).CharmInfo), arg0)
}

// CheckCharmPlacement mocks base method.
func (m *MockDeployerAPI) CheckCharmPlacement(arg0 string, arg1 *charm.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCharmPlacement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCharmPlacement indicates an expected call of CheckCharmPlacement.
func (mr *MockDeployerAPIMockRecorder) CheckCharmPlacement(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCharmPlacement", reflect.TypeOf((*MockDeployerAPI)(nil).CheckCharmPlacement), arg0, arg1)
}

// Close mocks base method.
func (m *MockDeployerAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDeployerAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDeployerAPI)(nil).Close))
}

// ConnectControllerStream mocks base method.
func (m *MockDeployerAPI) ConnectControllerStream(arg0 context.Context, arg1 string, arg2 url.Values, arg3 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream.
func (mr *MockDeployerAPIMockRecorder) ConnectControllerStream(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockDeployerAPI)(nil).ConnectControllerStream), arg0, arg1, arg2, arg3)
}

// ConnectStream mocks base method.
func (m *MockDeployerAPI) ConnectStream(arg0 context.Context, arg1 string, arg2 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream.
func (mr *MockDeployerAPIMockRecorder) ConnectStream(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockDeployerAPI)(nil).ConnectStream), arg0, arg1, arg2)
}

// Consume mocks base method.
func (m *MockDeployerAPI) Consume(arg0 crossmodel.ConsumeApplicationArgs) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockDeployerAPIMockRecorder) Consume(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockDeployerAPI)(nil).Consume), arg0)
}

// Deploy mocks base method.
func (m *MockDeployerAPI) Deploy(arg0 application.DeployArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockDeployerAPIMockRecorder) Deploy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployerAPI)(nil).Deploy), arg0)
}

// DeployFromRepository mocks base method.
func (m *MockDeployerAPI) DeployFromRepository(arg0 application.DeployFromRepositoryArg) (application.DeployInfo, []application.PendingResourceUpload, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployFromRepository", arg0)
	ret0, _ := ret[0].(application.DeployInfo)
	ret1, _ := ret[1].([]application.PendingResourceUpload)
	ret2, _ := ret[2].([]error)
	return ret0, ret1, ret2
}

// DeployFromRepository indicates an expected call of DeployFromRepository.
func (mr *MockDeployerAPIMockRecorder) DeployFromRepository(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployFromRepository", reflect.TypeOf((*MockDeployerAPI)(nil).DeployFromRepository), arg0)
}

// Expose mocks base method.
func (m *MockDeployerAPI) Expose(arg0 string, arg1 map[string]params.ExposedEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expose", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Expose indicates an expected call of Expose.
func (mr *MockDeployerAPIMockRecorder) Expose(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expose", reflect.TypeOf((*MockDeployerAPI)(nil).Expose), arg0, arg1)
}

// GetAnnotations mocks base method.
func (m *MockDeployerAPI) GetAnnotations(arg0 []string) ([]params.AnnotationsGetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations", arg0)
	ret0, _ := ret[0].([]params.AnnotationsGetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnnotations indicates an expected call of GetAnnotations.
func (mr *MockDeployerAPIMockRecorder) GetAnnotations(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockDeployerAPI)(nil).GetAnnotations), arg0)
}

// GetCharmURLOrigin mocks base method.
func (m *MockDeployerAPI) GetCharmURLOrigin(arg0, arg1 string) (*charm.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmURLOrigin", arg0, arg1)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCharmURLOrigin indicates an expected call of GetCharmURLOrigin.
func (mr *MockDeployerAPIMockRecorder) GetCharmURLOrigin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmURLOrigin", reflect.TypeOf((*MockDeployerAPI)(nil).GetCharmURLOrigin), arg0, arg1)
}

// GetConfig mocks base method.
func (m *MockDeployerAPI) GetConfig(arg0 string, arg1 ...string) ([]map[string]any, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfig", varargs...)
	ret0, _ := ret[0].([]map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockDeployerAPIMockRecorder) GetConfig(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockDeployerAPI)(nil).GetConfig), varargs...)
}

// GetConstraints mocks base method.
func (m *MockDeployerAPI) GetConstraints(arg0 ...string) ([]constraints.Value, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConstraints", varargs...)
	ret0, _ := ret[0].([]constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConstraints indicates an expected call of GetConstraints.
func (mr *MockDeployerAPIMockRecorder) GetConstraints(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConstraints", reflect.TypeOf((*MockDeployerAPI)(nil).GetConstraints), arg0...)
}

// GetModelConstraints mocks base method.
func (m *MockDeployerAPI) GetModelConstraints() (constraints.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelConstraints")
	ret0, _ := ret[0].(constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelConstraints indicates an expected call of GetModelConstraints.
func (mr *MockDeployerAPIMockRecorder) GetModelConstraints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelConstraints", reflect.TypeOf((*MockDeployerAPI)(nil).GetModelConstraints))
}

// GrantOffer mocks base method.
func (m *MockDeployerAPI) GrantOffer(arg0, arg1 string, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GrantOffer", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantOffer indicates an expected call of GrantOffer.
func (mr *MockDeployerAPIMockRecorder) GrantOffer(arg0, arg1 any, arg2 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantOffer", reflect.TypeOf((*MockDeployerAPI)(nil).GrantOffer), varargs...)
}

// HTTPClient mocks base method.
func (m *MockDeployerAPI) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockDeployerAPIMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockDeployerAPI)(nil).HTTPClient))
}

// ListCharmResources mocks base method.
func (m *MockDeployerAPI) ListCharmResources(arg0 string, arg1 charm0.Origin) ([]resource.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCharmResources", arg0, arg1)
	ret0, _ := ret[0].([]resource.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCharmResources indicates an expected call of ListCharmResources.
func (mr *MockDeployerAPIMockRecorder) ListCharmResources(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCharmResources", reflect.TypeOf((*MockDeployerAPI)(nil).ListCharmResources), arg0, arg1)
}

// ListSpaces mocks base method.
func (m *MockDeployerAPI) ListSpaces() ([]params.Space, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSpaces")
	ret0, _ := ret[0].([]params.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSpaces indicates an expected call of ListSpaces.
func (mr *MockDeployerAPIMockRecorder) ListSpaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSpaces", reflect.TypeOf((*MockDeployerAPI)(nil).ListSpaces))
}

// ModelGet mocks base method.
func (m *MockDeployerAPI) ModelGet() (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet.
func (mr *MockDeployerAPIMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockDeployerAPI)(nil).ModelGet))
}

// ModelTag mocks base method.
func (m *MockDeployerAPI) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockDeployerAPIMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockDeployerAPI)(nil).ModelTag))
}

// ModelUUID mocks base method.
func (m *MockDeployerAPI) ModelUUID() (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelUUID indicates an expected call of ModelUUID.
func (mr *MockDeployerAPIMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockDeployerAPI)(nil).ModelUUID))
}

// Offer mocks base method.
func (m *MockDeployerAPI) Offer(arg0, arg1 string, arg2 []string, arg3, arg4, arg5 string) ([]params.ErrorResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offer", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]params.ErrorResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offer indicates an expected call of Offer.
func (mr *MockDeployerAPIMockRecorder) Offer(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offer", reflect.TypeOf((*MockDeployerAPI)(nil).Offer), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RootHTTPClient mocks base method.
func (m *MockDeployerAPI) RootHTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootHTTPClient indicates an expected call of RootHTTPClient.
func (mr *MockDeployerAPIMockRecorder) RootHTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHTTPClient", reflect.TypeOf((*MockDeployerAPI)(nil).RootHTTPClient))
}

// ScaleApplication mocks base method.
func (m *MockDeployerAPI) ScaleApplication(arg0 application.ScaleApplicationParams) (params.ScaleApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleApplication", arg0)
	ret0, _ := ret[0].(params.ScaleApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScaleApplication indicates an expected call of ScaleApplication.
func (mr *MockDeployerAPIMockRecorder) ScaleApplication(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleApplication", reflect.TypeOf((*MockDeployerAPI)(nil).ScaleApplication), arg0)
}

// Sequences mocks base method.
func (m *MockDeployerAPI) Sequences() (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sequences")
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sequences indicates an expected call of Sequences.
func (mr *MockDeployerAPIMockRecorder) Sequences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sequences", reflect.TypeOf((*MockDeployerAPI)(nil).Sequences))
}

// SetAnnotation mocks base method.
func (m *MockDeployerAPI) SetAnnotation(arg0 map[string]map[string]string) ([]params.ErrorResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAnnotation", arg0)
	ret0, _ := ret[0].([]params.ErrorResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAnnotation indicates an expected call of SetAnnotation.
func (mr *MockDeployerAPIMockRecorder) SetAnnotation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotation", reflect.TypeOf((*MockDeployerAPI)(nil).SetAnnotation), arg0)
}

// SetCharm mocks base method.
func (m *MockDeployerAPI) SetCharm(arg0 string, arg1 application.SetCharmConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharm indicates an expected call of SetCharm.
func (mr *MockDeployerAPIMockRecorder) SetCharm(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockDeployerAPI)(nil).SetCharm), arg0, arg1)
}

// SetConfig mocks base method.
func (m *MockDeployerAPI) SetConfig(arg0, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig.
func (mr *MockDeployerAPIMockRecorder) SetConfig(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockDeployerAPI)(nil).SetConfig), arg0, arg1, arg2, arg3)
}

// SetConstraints mocks base method.
func (m *MockDeployerAPI) SetConstraints(arg0 string, arg1 constraints.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConstraints", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints.
func (mr *MockDeployerAPIMockRecorder) SetConstraints(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockDeployerAPI)(nil).SetConstraints), arg0, arg1)
}

// Status mocks base method.
func (m *MockDeployerAPI) Status(arg0 *client.StatusArgs) (*params.FullStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(*params.FullStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockDeployerAPIMockRecorder) Status(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockDeployerAPI)(nil).Status), arg0)
}

// WatchAll mocks base method.
func (m *MockDeployerAPI) WatchAll() (api.AllWatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAll")
	ret0, _ := ret[0].(api.AllWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAll indicates an expected call of WatchAll.
func (mr *MockDeployerAPIMockRecorder) WatchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAll", reflect.TypeOf((*MockDeployerAPI)(nil).WatchAll))
}

// MockCharmReader is a mock of CharmReader interface.
type MockCharmReader struct {
	ctrl     *gomock.Controller
	recorder *MockCharmReaderMockRecorder
}

// MockCharmReaderMockRecorder is the mock recorder for MockCharmReader.
type MockCharmReaderMockRecorder struct {
	mock *MockCharmReader
}

// NewMockCharmReader creates a new mock instance.
func NewMockCharmReader(ctrl *gomock.Controller) *MockCharmReader {
	mock := &MockCharmReader{ctrl: ctrl}
	mock.recorder = &MockCharmReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmReader) EXPECT() *MockCharmReaderMockRecorder {
	return m.recorder
}

// NewCharmAtPath mocks base method.
func (m *MockCharmReader) NewCharmAtPath(arg0 string) (charm.Charm, *charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCharmAtPath", arg0)
	ret0, _ := ret[0].(charm.Charm)
	ret1, _ := ret[1].(*charm.URL)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// NewCharmAtPath indicates an expected call of NewCharmAtPath.
func (mr *MockCharmReaderMockRecorder) NewCharmAtPath(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCharmAtPath", reflect.TypeOf((*MockCharmReader)(nil).NewCharmAtPath), arg0)
}

// MockDeployConfigFlag is a mock of DeployConfigFlag interface.
type MockDeployConfigFlag struct {
	ctrl     *gomock.Controller
	recorder *MockDeployConfigFlagMockRecorder
}

// MockDeployConfigFlagMockRecorder is the mock recorder for MockDeployConfigFlag.
type MockDeployConfigFlagMockRecorder struct {
	mock *MockDeployConfigFlag
}

// NewMockDeployConfigFlag creates a new mock instance.
func NewMockDeployConfigFlag(ctrl *gomock.Controller) *MockDeployConfigFlag {
	mock := &MockDeployConfigFlag{ctrl: ctrl}
	mock.recorder = &MockDeployConfigFlagMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployConfigFlag) EXPECT() *MockDeployConfigFlagMockRecorder {
	return m.recorder
}

// AbsoluteFileNames mocks base method.
func (m *MockDeployConfigFlag) AbsoluteFileNames(arg0 *cmd.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbsoluteFileNames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbsoluteFileNames indicates an expected call of AbsoluteFileNames.
func (mr *MockDeployConfigFlagMockRecorder) AbsoluteFileNames(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbsoluteFileNames", reflect.TypeOf((*MockDeployConfigFlag)(nil).AbsoluteFileNames), arg0)
}

// ReadConfigPairs mocks base method.
func (m *MockDeployConfigFlag) ReadConfigPairs(arg0 *cmd.Context) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConfigPairs", arg0)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadConfigPairs indicates an expected call of ReadConfigPairs.
func (mr *MockDeployConfigFlagMockRecorder) ReadConfigPairs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConfigPairs", reflect.TypeOf((*MockDeployConfigFlag)(nil).ReadConfigPairs), arg0)
}
