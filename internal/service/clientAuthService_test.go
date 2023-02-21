package service

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_clients "github.com/yurchenkosv/credential_storage/internal/mockCredStorageClient"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"testing"
)

func TestClientAuthService_Authenticate(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, login string, password string)
	type args struct {
		ctx   context.Context
		login string
		pwd   string
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         string
		wantErr      bool
	}{
		{
			name: "should successfully authenticate user",
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, login string, password string) {
				s.EXPECT().AuthenticateUser(ctx, login, password).Return("token", nil)
			},
			args: args{
				ctx:   context.Background(),
				login: "test_user",
				pwd:   "test_pwd",
			},
			want:    "token",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.args.ctx, client, tt.args.login, tt.args.pwd)
			s := &ClientAuthService{
				client: client,
			}
			got, err := s.Authenticate(tt.args.ctx, tt.args.login, tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientAuthService_Register(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, user model.User)
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         string
		wantErr      bool
	}{
		{
			name: "should successfully register user",
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, user model.User) {
				s.EXPECT().RegisterUser(ctx, user).Return("token", nil)
			},
			args: args{
				ctx: context.Background(),
				user: model.User{
					Username: "test_username",
					Password: "test_pwd",
					Name:     "TestUserName",
				},
			},
			want:    "token",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.args.ctx, client, tt.args.user)
			s := &ClientAuthService{
				client: client,
			}
			got, err := s.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
