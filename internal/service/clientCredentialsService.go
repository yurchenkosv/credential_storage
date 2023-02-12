package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"io"
)

type ClientCredentialsService struct {
	client           clients.CredentialsStorageClient
	binaryRepository repository.BinaryRepository
	ctx              context.Context
}

func NewClientCredentialsService(ctx context.Context,
	client clients.CredentialsStorageClient,
	binaryRepository repository.BinaryRepository,
) *ClientCredentialsService {
	return &ClientCredentialsService{
		client:           client,
		ctx:              ctx,
		binaryRepository: binaryRepository,
	}
}

func (s *ClientCredentialsService) GetData() ([]model.Credentials, error) {
	return s.client.GetData(s.ctx)
}

func (s *ClientCredentialsService) SaveBinary(reader io.Reader, filename string) error {
	_, err := s.binaryRepository.Save(reader, filename)
	return err
}

func (s *ClientCredentialsService) SendBankCard(card model.BankingCardData) error {
	return s.client.SendBankCard(s.ctx, card)
}

func (s *ClientCredentialsService) SendCredentials(credentials model.CredentialsData) error {
	return s.client.SendCredentials(s.ctx, credentials)
}

func (s *ClientCredentialsService) SendText(data model.TextData) error {
	return s.client.SendText(s.ctx, data)
}

func (s *ClientCredentialsService) SendBinary(data model.BinaryData) error {
	return s.client.SendBinary(s.ctx, data)
}
