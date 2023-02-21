package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type ClientAuthService struct {
	client clients.CredentialsStorageClient
}

func NewClientAuthService(client clients.CredentialsStorageClient) *ClientAuthService {
	return &ClientAuthService{client: client}
}

func (s *ClientAuthService) Authenticate(ctx context.Context, login string, pwd string) (string, error) {
	jwt, err := s.client.AuthenticateUser(ctx, login, pwd)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (s *ClientAuthService) Register(ctx context.Context, user model.User) (string, error) {
	jwt, err := s.client.RegisterUser(ctx, user)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
