package service

import (
	"context"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	mock_repository "github.com/yurchenkosv/credential_storage/internal/mockRepo"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"google.golang.org/grpc/metadata"
	"reflect"
	"testing"
	"time"
)

func intPointer(val int) *int {
	return &val
}
func createToken(key string, claims map[string]interface{}) string {
	auth := jwtauth.New("HS256", []byte(key), nil)
	currentTime := time.Now()

	jwtauth.SetIssuedAt(claims, currentTime)
	jwtauth.SetExpiry(claims, currentTime.Add(5*time.Minute))
	_, token, err := auth.Encode(claims)
	if err != nil {
		log.Fatal(err)
	}
	return token
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
		{
			name: "should fail with repo error",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username: "test",
					Password: "testPWD",
					Name:     "test",
				},
			},
			want: nil,
			behavior: func(ctx context.Context, s *mock_repository.MockRepository, user *model.User) {
				s.EXPECT().GetUser(ctx, user).Return(nil, errors.New("database error"))
			},
			wantErr: true,
		},
		{
			name: "should fail with user already exists",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:       intPointer(1),
					Username: "test",
					Password: "testPWD",
					Name:     "test",
				},
			},
			want: nil,
			behavior: func(ctx context.Context, s *mock_repository.MockRepository, user *model.User) {
				s.EXPECT().GetUser(ctx, user).Return(user, nil)
			},
			wantErr: true,
		},
		{
			name: "should fail with save user error",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username: "test",
					Password: "testPWD",
					Name:     "test",
				},
			},
			want: nil,
			behavior: func(ctx context.Context, s *mock_repository.MockRepository, user *model.User) {
				s.EXPECT().GetUser(ctx, user).Return(user, nil)
				s.EXPECT().SaveUser(ctx, user).Return(errors.New("database error"))
			},
			wantErr: true,
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

func TestAuthService_GetJWTTokenFromGRPCContext(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should successfuly get token from context",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{"jwt": "token"})),
			},
			want:    "token",
			wantErr: false,
		},
		{
			name: "should fail with no metadata",
			args: args{
				ctx: context.Background(),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "should fail with wrong metadata",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{"test": "val"})),
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authService := &AuthService{}
			got, err := authService.GetJWTTokenFromGRPCContext(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWTTokenFromGRPCContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetJWTTokenFromGRPCContext() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_GetUserFromToken(t *testing.T) {
	type fields struct {
		auth *jwtauth.JWTAuth
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "should successfully get user from valid token",
			fields: fields{
				auth: jwtauth.New("HS256", []byte("test"), nil),
			},
			args: args{
				token: createToken("test", map[string]interface{}{"user_id": 1}),
			},
			want: &model.User{
				ID: intPointer(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authService := &AuthService{
				auth: tt.fields.auth,
			}
			got, err := authService.GetUserFromToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserFromToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserFromToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_GenerateToken(t *testing.T) {
	type fields struct {
		auth *jwtauth.JWTAuth
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should successfully generate token",
			fields: fields{
				auth: jwtauth.New("HS256", []byte("test"), nil),
			},
			args: args{
				user: &model.User{ID: intPointer(1)},
			},
			want:    createToken("test", map[string]interface{}{"user_id": 1}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authService := &AuthService{
				auth: tt.fields.auth,
			}
			got, err := authService.GenerateToken(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
