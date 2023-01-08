// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\repository\repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yurchenkosv/credential_storage/internal/model"
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

// GetCredentialsByName mocks base method.
func (m *MockRepository) GetCredentialsByName(ctx context.Context, name string, userID int) ([]*model.CredentialsData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsByName", ctx, name, userID)
	ret0, _ := ret[0].([]*model.CredentialsData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsByName indicates an expected call of GetCredentialsByName.
func (mr *MockRepositoryMockRecorder) GetCredentialsByName(ctx, name, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsByName", reflect.TypeOf((*MockRepository)(nil).GetCredentialsByName), ctx, name, userID)
}

// GetCredentialsByUserID mocks base method.
func (m *MockRepository) GetCredentialsByUserID(ctx context.Context, userID int) ([]*model.CredentialsData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsByUserID", ctx, userID)
	ret0, _ := ret[0].([]*model.CredentialsData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsByUserID indicates an expected call of GetCredentialsByUserID.
func (mr *MockRepositoryMockRecorder) GetCredentialsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsByUserID", reflect.TypeOf((*MockRepository)(nil).GetCredentialsByUserID), ctx, userID)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), ctx, user)
}

// SaveBankingCardData mocks base method.
func (m *MockRepository) SaveBankingCardData(ctx context.Context, creds *model.BankingCardData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBankingCardData", ctx, creds, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBankingCardData indicates an expected call of SaveBankingCardData.
func (mr *MockRepositoryMockRecorder) SaveBankingCardData(ctx, creds, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBankingCardData", reflect.TypeOf((*MockRepository)(nil).SaveBankingCardData), ctx, creds, userID)
}

// SaveBinaryData mocks base method.
func (m *MockRepository) SaveBinaryData(ctx context.Context, creds *model.BinaryData, userID int, link string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveBinaryData", ctx, creds, userID, link)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveBinaryData indicates an expected call of SaveBinaryData.
func (mr *MockRepositoryMockRecorder) SaveBinaryData(ctx, creds, userID, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveBinaryData", reflect.TypeOf((*MockRepository)(nil).SaveBinaryData), ctx, creds, userID, link)
}

// SaveCredentialsData mocks base method.
func (m *MockRepository) SaveCredentialsData(ctx context.Context, creds *model.CredentialsData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCredentialsData", ctx, creds, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCredentialsData indicates an expected call of SaveCredentialsData.
func (mr *MockRepositoryMockRecorder) SaveCredentialsData(ctx, creds, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCredentialsData", reflect.TypeOf((*MockRepository)(nil).SaveCredentialsData), ctx, creds, userID)
}

// SaveTextData mocks base method.
func (m *MockRepository) SaveTextData(ctx context.Context, creds *model.TextData, userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTextData", ctx, creds, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTextData indicates an expected call of SaveTextData.
func (mr *MockRepositoryMockRecorder) SaveTextData(ctx, creds, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTextData", reflect.TypeOf((*MockRepository)(nil).SaveTextData), ctx, creds, userID)
}

// SaveUser mocks base method.
func (m *MockRepository) SaveUser(ctx context.Context, user *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser.
func (mr *MockRepositoryMockRecorder) SaveUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockRepository)(nil).SaveUser), ctx, user)
}