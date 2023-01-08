package service

import (
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type ClientCredentialsService struct {
	client clients.CredentialsStorageClient
}

func NewClientCredentialsService(client clients.CredentialsStorageClient) *ClientCredentialsService {
	return &ClientCredentialsService{client: client}
}

func (s *ClientCredentialsService) GetBankCard() (model.BankingCardData, error) {
	return s.client.GetBankCard()
}

func (s *ClientCredentialsService) GetCredentials() (model.CredentialsData, error) {
	return s.client.GetCredentials()
}

func (s *ClientCredentialsService) GetText() (model.TextData, error) {
	return s.client.GetText()
}

func (s *ClientCredentialsService) GetBinary() (model.BinaryData, error) {
	return s.client.GetBinary()
}

func (s *ClientCredentialsService) SendBankCard(card model.BankingCardData) error {
	return s.client.SendBankCard(card)
}

func (s *ClientCredentialsService) SendCredentials(credentials model.CredentialsData) error {
	return s.client.SendCredentials(credentials)
}

func (s *ClientCredentialsService) SendText(data model.TextData) error {
	return s.client.SendText(data)
}

func (s *ClientCredentialsService) SendBinary(data []byte) error {
	binary := model.BinaryData{Data: data}
	return s.client.SendBinary(binary)
}
