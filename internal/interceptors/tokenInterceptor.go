package interceptors

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/contextKeys"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"google.golang.org/grpc"
)

type AuthInterceptor struct {
	authSvc service.Auth
}

func NewAuthInterceptor(svc service.Auth) *AuthInterceptor {
	return &AuthInterceptor{authSvc: svc}
}

func (i *AuthInterceptor) JWTInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod == "/api.AuthService/AuthenticateUser" ||
		info.FullMethod == "/api.AuthService/RegisterUser" {
		return handler(ctx, req)
	}
	token, err := i.authSvc.GetJWTTokenFromGRPCContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := i.authSvc.GetUserFromToken(token)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, contextKeys.UserIDContexKey("user_id"), *user.ID)
	return handler(ctx, req)
}
