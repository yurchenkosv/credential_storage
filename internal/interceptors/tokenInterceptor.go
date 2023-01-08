package interceptors

import (
	"context"
	"errors"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get metadata from context")
	}
	tokens := md.Get("jwt")
	if len(tokens) == 0 {
		return nil, errors.New("no authorization token found in metadata")
	}
	token := tokens[0]
	_, err := i.authSvc.GetUserFromToken(token)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
