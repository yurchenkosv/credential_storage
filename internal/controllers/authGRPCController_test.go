package controllers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/yurchenkosv/credential_storage/internal/api"
	mock_service "github.com/yurchenkosv/credential_storage/internal/mockService"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"reflect"
	"testing"
)

func TestAuthGRPCController_AuthenticateUser(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockAuth, userAuth *api.UserAuthentication)
	type args struct {
		ctx context.Context
		in  *api.UserAuthentication
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerAuthResponse
		wantErr      bool
	}{
		{
			name: "should authenticate user",
			mockBehavior: func(ctx context.Context, s *mock_service.MockAuth, userAuth *api.UserAuthentication) {
				user := model.User{
					Username: userAuth.Login,
					Password: userAuth.Password,
				}
				s.EXPECT().AuthenticateUser(ctx, &user).Return(&user, nil)
				s.EXPECT().GenerateToken(&user).Return("token", nil)
			},
			args: args{
				ctx: context.Background(),
				in: &api.UserAuthentication{
					Login:    "test",
					Password: "test",
				},
			},
			want: &api.ServerAuthResponse{
				Message: "Successfully authorized",
				Code:    200,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockAuth(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.in)
			c := NewAuthGRPCController(svc)
			got, err := c.AuthenticateUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthGRPCController_RegisterUser(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockAuth, userAuth *api.UserRegistration)
	type args struct {
		ctx context.Context
		in  *api.UserRegistration
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerAuthResponse
		wantErr      bool
	}{
		{
			name: "should register user",
			mockBehavior: func(ctx context.Context, s *mock_service.MockAuth, userAuth *api.UserRegistration) {
				user := model.User{
					Username: userAuth.Name,
					Password: userAuth.Password,
					Name:     userAuth.Name,
				}
				s.EXPECT().RegisterUser(ctx, &user).Return(&user, nil)
				s.EXPECT().GenerateToken(&user).Return("token", nil)
			},
			args: args{
				ctx: context.Background(),
				in: &api.UserRegistration{
					Login:    "test",
					Password: "test",
					Name:     "test",
				},
			},
			want: &api.ServerAuthResponse{
				Message: "Successfully registered",
				Code:    200,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockAuth(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.in)
			c := NewAuthGRPCController(svc)
			got, err := c.RegisterUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
