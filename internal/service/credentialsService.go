package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"io"
)

type CredentialsService struct {
	repo       repository.Repository
	binaryRepo repository.BinaryRepository
}

func NewCredentialsService(repo repository.Repository, binaryRepo repository.BinaryRepository) *CredentialsService {
	return &CredentialsService{
		repo:       repo,
		binaryRepo: binaryRepo,
	}
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

func (s *CredentialsService) SaveBinaryData(ctx context.Context, reader io.Reader, data *model.BinaryData, userID int) error {
	fileID := uuid.New()
	link, err := s.binaryRepo.Save(reader, fileID.String())
	if err != nil {
		return err
	}
	err = s.repo.SaveBinaryData(ctx, data, userID, link)
	if err != nil {
		return err
	}
	return nil
}

func (s *CredentialsService) GetCredentialsByName(ctx context.Context, credName string, userID int) ([]model.CredentialsData, error) {
	creds, err := s.repo.GetCredentialsByName(ctx, credName, userID)
	return creds, err
}

func (s *CredentialsService) GetAllUserCredentials(ctx context.Context, userID int) ([]model.Credentials, error) {
	data, err := s.repo.GetCredentialsByUserID(ctx, userID)
	for _, dt := range data {
		if dt.BinaryData != nil {
			bindt, binErr := s.binaryRepo.Load(dt.BinaryData.Link)
			if binErr != nil {
				return nil, binErr
			}
			dt.BinaryData.Data = bindt
		}
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
