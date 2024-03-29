// Code generated by MockGen. DO NOT EDIT.
// Source: .\internal\service\auth.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/yurchenkosv/credential_storage/internal/model"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// AuthenticateUser mocks base method.
func (m *MockAuth) AuthenticateUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthenticateUser indicates an expected call of AuthenticateUser.
func (mr *MockAuthMockRecorder) AuthenticateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateUser", reflect.TypeOf((*MockAuth)(nil).AuthenticateUser), ctx, user)
}

// GenerateToken mocks base method.
func (m *MockAuth) GenerateToken(user *model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuth)(nil).GenerateToken), user)
}

// GetJWTTokenFromGRPCContext mocks base method.
func (m *MockAuth) GetJWTTokenFromGRPCContext(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJWTTokenFromGRPCContext", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJWTTokenFromGRPCContext indicates an expected call of GetJWTTokenFromGRPCContext.
func (mr *MockAuthMockRecorder) GetJWTTokenFromGRPCContext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJWTTokenFromGRPCContext", reflect.TypeOf((*MockAuth)(nil).GetJWTTokenFromGRPCContext), ctx)
}

// GetUserFromToken mocks base method.
func (m *MockAuth) GetUserFromToken(token string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFromToken", token)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFromToken indicates an expected call of GetUserFromToken.
func (mr *MockAuthMockRecorder) GetUserFromToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFromToken", reflect.TypeOf((*MockAuth)(nil).GetUserFromToken), token)
}

// RegisterUser mocks base method.
func (m *MockAuth) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockAuthMockRecorder) RegisterUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockAuth)(nil).RegisterUser), ctx, user)
}
