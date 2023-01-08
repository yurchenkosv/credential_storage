package service

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_repository "github.com/yurchenkosv/credential_storage/internal/mockRepo"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"reflect"
	"testing"
)

func intPointer(val int) *int {
	return &val
}

func TestAuthService_AuthenticateUser(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, user *model.User)

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name     string
		args     args
		want     *model.User
		behavior mockBehavior
		wantErr  bool
	}{
		{
			name: "should successfully authenticate user",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:       nil,
					Username: "test",
					Password: "testPWD",
					Name:     "test",
				},
			},
			want: &model.User{
				ID:       intPointer(1),
				Username: "test",
				Password: hashPW("testPWD"),
				Name:     "test",
			},
			behavior: func(ctx context.Context, s *mock_repository.MockRepository, user *model.User) {
				returnUser := user
				returnUser.ID = intPointer(1)
				s.EXPECT().GetUser(ctx, user).Return(returnUser, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockRepository(ctrl)
			tt.behavior(tt.args.ctx, repo, tt.args.user)

			auth := AuthService{
				repo: repo,
			}
			got, err := auth.AuthenticateUser(tt.args.ctx, tt.args.user)
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

func TestAuthService_RegisterUser(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, user *model.User)

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name     string
		args     args
		want     *model.User
		behavior mockBehavior
		wantErr  bool
	}{
		{
			name: "should successfuly register user",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username: "test",
					Password: "testPWD",
					Name:     "test",
				},
			},
			want: &model.User{
				ID:       intPointer(1),
				Username: "test",
				Password: hashPW("testPWD"),
				Name:     "test",
			},
			behavior: func(ctx context.Context, s *mock_repository.MockRepository, user *model.User) {
				s.EXPECT().GetUser(ctx, user).Return(user, nil)
				s.EXPECT().SaveUser(ctx, user).Return(nil)

				returnedUser := *user
				returnedUser.ID = intPointer(1)
				returnedUser.Password = hashPW(user.Password)
				s.EXPECT().GetUser(ctx, user).Return(&returnedUser, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockRepository(ctrl)
			tt.behavior(tt.args.ctx, repo, tt.args.user)

			auth := AuthService{
				repo: repo,
			}
			got, err := auth.RegisterUser(tt.args.ctx, tt.args.user)
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

func TestNewAuthService(t *testing.T) {
	type args struct {
		repo repository.Repository
	}
	tests := []struct {
		name string
		args args
		want Auth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashPW(t *testing.T) {
	type args struct {
		pw string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashPW(tt.args.pw); got != tt.want {
				t.Errorf("hashPW() = %v, want %v", got, tt.want)
			}
		})
	}
}
