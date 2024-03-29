// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\service\dataService.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yurchenkosv/credential_storage/internal/model"
)

// MockDataService is a mock of DataService interface.
type MockDataService struct {
	ctrl     *gomock.Controller
	recorder *MockDataServiceMockRecorder
}

// MockDataServiceMockRecorder is the mock recorder for MockDataService.
type MockDataServiceMockRecorder struct {
	mock *MockDataService
}

// NewMockDataService creates a new mock instance.
func NewMockDataService(ctrl *gomock.Controller) *MockDataService {
	mock := &MockDataService{ctrl: ctrl}
	mock.recorder = &MockDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataService) EXPECT() *MockDataServiceMockRecorder {
	return m.recorder
}

// DeleteCredential mocks base method.
func (m *MockDataService) DeleteCredential(ctx context.Context, data model.Credentials, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCredential", ctx, data, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCredential indicates an expected call of DeleteCredential.
func (mr *MockDataServiceMockRecorder) DeleteCredential(ctx, data, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCredential", reflect.TypeOf((*MockDataService)(nil).DeleteCredential), ctx, data, userID)
}

// GetAllUserCredentials mocks base method.
func (m *MockDataService) GetAllUserCredentials(ctx context.Context, userID int) ([]model.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserCredentials", ctx, userID)
	ret0, _ := ret[0].([]model.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserCredentials indicates an expected call of GetAllUserCredentials.
func (mr *MockDataServiceMockRecorder) GetAllUserCredentials(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserCredentials", reflect.TypeOf((*MockDataService)(nil).GetAllUserCredentials), ctx, userID)
}

// GetCredentialsByName mocks base method.
func (m *MockDataService) GetCredentialsByName(ctx context.Context, credName string, userID int) ([]model.CredentialsData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsByName", ctx, credName, userID)
	ret0, _ := ret[0].([]model.CredentialsData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsByName indicates an expected call of GetCredentialsByName.
func (mr *MockDataServiceMockRecorder) GetCredentialsByName(ctx, credName, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsByName", reflect.TypeOf((*MockDataService)(nil).GetCredentialsByName), ctx, credName, userID)
}

// SaveBankingCardData mocks base method.
func (m *MockDataService) SaveBankingCardData(ctx context.Context, data *model.BankingCardData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBankingCardData", ctx, data, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBankingCardData indicates an expected call of SaveBankingCardData.
func (mr *MockDataServiceMockRecorder) SaveBankingCardData(ctx, data, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBankingCardData", reflect.TypeOf((*MockDataService)(nil).SaveBankingCardData), ctx, data, userID)
}

// SaveBinaryData mocks base method.
func (m *MockDataService) SaveBinaryData(ctx context.Context, reader io.Reader, data *model.BinaryData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBinaryData", ctx, reader, data, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBinaryData indicates an expected call of SaveBinaryData.
func (mr *MockDataServiceMockRecorder) SaveBinaryData(ctx, reader, data, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBinaryData", reflect.TypeOf((*MockDataService)(nil).SaveBinaryData), ctx, reader, data, userID)
}

// SaveCredentialsData mocks base method.
func (m *MockDataService) SaveCredentialsData(ctx context.Context, data *model.CredentialsData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCredentialsData", ctx, data, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCredentialsData indicates an expected call of SaveCredentialsData.
func (mr *MockDataServiceMockRecorder) SaveCredentialsData(ctx, data, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCredentialsData", reflect.TypeOf((*MockDataService)(nil).SaveCredentialsData), ctx, data, userID)
}

// SaveTextData mocks base method.
func (m *MockDataService) SaveTextData(ctx context.Context, data *model.TextData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTextData", ctx, data, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTextData indicates an expected call of SaveTextData.
func (mr *MockDataServiceMockRecorder) SaveTextData(ctx, data, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTextData", reflect.TypeOf((*MockDataService)(nil).SaveTextData), ctx, data, userID)
}
