package service

import (
	"context"
	"errors"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
)

type CredentialsService struct {
	repo repository.Repository
}

func NewCredentialsService(repo repository.Repository) *CredentialsService {
	return &CredentialsService{repo: repo}
}

func (s *CredentialsService) SaveCredentialsData(ctx context.Context, data *model.CredentialsData, userID int) error {
	err := s.repo.SaveCredentialsData(ctx, data, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CredentialsService) SaveBankingCardData(ctx context.Context, data *model.BankingCardData, userID int) error {
	err := s.repo.SaveBankingCardData(ctx, data, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CredentialsService) SaveTextData(ctx context.Context, data *model.TextData, userID int) error {
	err := s.repo.SaveTextData(ctx, data, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CredentialsService) SaveBinaryData(ctx context.Context, data *model.CredentialsData, userID int) error {
	return errors.New("Not implemented")
}

func (s *CredentialsService) GetCredentialsByName(ctx context.Context, credName string, userID int) ([]*model.CredentialsData, error) {
	creds, err := s.repo.GetCredentialsByName(ctx, credName, userID)
	return creds, err
}

func (s *CredentialsService) GetAllUserCredentials(ctx context.Context, userID int) ([]*model.CredentialsData, error) {
	return nil, errors.New("Not implemented")
}
