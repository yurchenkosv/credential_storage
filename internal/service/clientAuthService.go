package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"google.golang.org/grpc/metadata"
)

type ClientAuthService struct {
	client clients.CredentialsStorageClient
}

func NewClientAuthService(client clients.CredentialsStorageClient) *ClientAuthService {
	return &ClientAuthService{client: client}
}

func (s *ClientAuthService) Authenticate(ctx context.Context, login string, pwd string) (string, error) {
	jwt, err := s.client.AuthenticateUser(ctx, login, pwd)
	addJWTToContext(ctx, jwt)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (s *ClientAuthService) Register(ctx context.Context, user model.User) (string, error) {
	jwt, err := s.client.RegisterUser(ctx, user)
	addJWTToContext(ctx, jwt)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func addJWTToContext(ctx context.Context, jwt string) {
	meta := metadata.New(map[string]string{"jwt": jwt})
	ctx = metadata.NewOutgoingContext(ctx, meta)
}
