package controllers

import (
	"context"
	"errors"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/service"
)

type CredentialsGRPCCOntroller struct {
	svc *service.CredentialsService
}

func NewCredentialsGRPCController(svc *service.CredentialsService) *CredentialsGRPCCOntroller {
	return &CredentialsGRPCCOntroller{svc: svc}
}

func (c *CredentialsGRPCCOntroller) SaveCredentialsData(context.Context, *api.CredentialsData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")
}

func (c CredentialsGRPCCOntroller) SaveBankingData(context.Context, *api.BankingCardData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")
}

func (c CredentialsGRPCCOntroller) SaveTextData(context.Context, *api.TextData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) SaveBinaryData(context.Context, *api.BinaryData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetCredentialsData(context.Context, *api.CredentialsDataRequest) (*api.CredentialsData, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetBankingCardData(context.Context, *api.BankingCardDataRequest) (*api.BankingCardData, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetTextData(context.Context, *api.TextDataRequest) (*api.TextData, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetBinaryData(context.Context, *api.BinaryDataRequest) (*api.BinaryData, error) {
	return nil, errors.New("not implemented")

}
