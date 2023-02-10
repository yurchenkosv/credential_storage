package interceptors

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_service "github.com/yurchenkosv/credential_storage/internal/mockService"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

func intPtr(v int) *int {
	return &v
}

func TestAuthInterceptor_JWTInterceptor(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockAuth, data *model.User)
	type args struct {
		data    *model.User
		ctx     context.Context
		req     interface{}
		info    *grpc.UnaryServerInfo
		handler grpc.UnaryHandler
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         interface{}
		wantErr      bool
	}{
		{
			name: "should successfully authenticate via jwt",
			mockBehavior: func(ctx context.Context, s *mock_service.MockAuth, data *model.User) {
				s.EXPECT().GetJWTTokenFromGRPCContext(ctx).Return("token", nil)
				s.EXPECT().GetUserFromToken("token").Return(data, nil)
			},
			args: args{
				data: &model.User{
					ID:       intPtr(1),
					Username: "test",
					Password: "test",
				},
				ctx: context.Background(),
				req: nil,
				info: &grpc.UnaryServerInfo{
					Server:     nil,
					FullMethod: "/api",
				},
				handler: grpc.UnaryHandler(func(ctx context.Context, req interface{}) (interface{}, error) {
					return req, nil
				}),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockAuth(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.data)
			i := NewAuthInterceptor(svc)
			got, err := i.JWTInterceptor(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTInterceptor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWTInterceptor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
