// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\clients\credentialsStorageClient.go

// Package mock_clients is a generated GoMock package.
package mock_clients

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yurchenkosv/credential_storage/internal/model"
)

// MockCredentialsStorageClient is a mock of CredentialsStorageClient interface.
type MockCredentialsStorageClient struct {
	ctrl     *gomock.Controller
	recorder *MockCredentialsStorageClientMockRecorder
}

// MockCredentialsStorageClientMockRecorder is the mock recorder for MockCredentialsStorageClient.
type MockCredentialsStorageClientMockRecorder struct {
	mock *MockCredentialsStorageClient
}

// NewMockCredentialsStorageClient creates a new mock instance.
func NewMockCredentialsStorageClient(ctrl *gomock.Controller) *MockCredentialsStorageClient {
	mock := &MockCredentialsStorageClient{ctrl: ctrl}
	mock.recorder = &MockCredentialsStorageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCredentialsStorageClient) EXPECT() *MockCredentialsStorageClientMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockCredentialsStorageClient) AuthenticateUser(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockCredentialsStorageClientMockRecorder) AuthenticateUser(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockCredentialsStorageClient)(nil).AuthenticateUser), ctx, login, password)
}

// DeleteData mocks base method.
func (m *MockCredentialsStorageClient) DeleteData(ctx context.Context, data model.Credentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteData", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteData indicates an expected call of DeleteData.
func (mr *MockCredentialsStorageClientMockRecorder) DeleteData(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteData", reflect.TypeOf((*MockCredentialsStorageClient)(nil).DeleteData), ctx, data)
}

// GetData mocks base method.
func (m *MockCredentialsStorageClient) GetData(ctx context.Context) ([]model.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", ctx)
	ret0, _ := ret[0].([]model.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockCredentialsStorageClientMockRecorder) GetData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockCredentialsStorageClient)(nil).GetData), ctx)
}

// RegisterUser mocks base method.
func (m *MockCredentialsStorageClient) RegisterUser(ctx context.Context, user model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockCredentialsStorageClientMockRecorder) RegisterUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockCredentialsStorageClient)(nil).RegisterUser), ctx, user)
}

// SendBankCard mocks base method.
func (m *MockCredentialsStorageClient) SendBankCard(ctx context.Context, card model.BankingCardData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBankCard", ctx, card)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendBankCard indicates an expected call of SendBankCard.
func (mr *MockCredentialsStorageClientMockRecorder) SendBankCard(ctx, card interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBankCard", reflect.TypeOf((*MockCredentialsStorageClient)(nil).SendBankCard), ctx, card)
}

// SendBinary mocks base method.
func (m *MockCredentialsStorageClient) SendBinary(ctx context.Context, binary model.BinaryData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBinary", ctx, binary)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendBinary indicates an expected call of SendBinary.
func (mr *MockCredentialsStorageClientMockRecorder) SendBinary(ctx, binary interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBinary", reflect.TypeOf((*MockCredentialsStorageClient)(nil).SendBinary), ctx, binary)
}

// SendCredentials mocks base method.
func (m *MockCredentialsStorageClient) SendCredentials(ctx context.Context, credentials model.CredentialsData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCredentials", ctx, credentials)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCredentials indicates an expected call of SendCredentials.
func (mr *MockCredentialsStorageClientMockRecorder) SendCredentials(ctx, credentials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCredentials", reflect.TypeOf((*MockCredentialsStorageClient)(nil).SendCredentials), ctx, credentials)
}

// SendText mocks base method.
func (m *MockCredentialsStorageClient) SendText(ctx context.Context, text model.TextData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendText", ctx, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendText indicates an expected call of SendText.
func (mr *MockCredentialsStorageClientMockRecorder) SendText(ctx, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendText", reflect.TypeOf((*MockCredentialsStorageClient)(nil).SendText), ctx, text)
}
