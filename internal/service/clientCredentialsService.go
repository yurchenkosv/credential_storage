package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type ClientCredentialsService struct {
	client clients.CredentialsStorageClient
	ctx    context.Context
}

func NewClientCredentialsService(ctx context.Context, client clients.CredentialsStorageClient) *ClientCredentialsService {
	return &ClientCredentialsService{client: client, ctx: ctx}
}

func (s *ClientCredentialsService) GetData() ([]model.Credentials, error) {
	return s.client.GetData(s.ctx)
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
