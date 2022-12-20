package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
)

type CredentialsService struct {
	repo repository.Repository
}

func NewCredentialsService(repo repository.Repository) *CredentialsService {
	return &CredentialsService{repo: repo}
}

func (s *CredentialsService) SaveCredentials(ctx context.Context, creds *model.Credentials) error {
	return nil
}

func (s *CredentialsService) GetCredentialsByName(ctx context.Context,
	credName string,
	userID int) (*model.Credentials, error) {
	return nil, nil
}
func (s *CredentialsService) GetAllUserCredentials(ctx context.Context, userID int) (*model.Credentials, error) {
	return nil, nil
}
