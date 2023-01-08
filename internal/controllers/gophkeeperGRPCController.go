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

func NewGophkeeperController(svc *service.CredentialsService) *CredentialsGRPCCOntroller {
	return &CredentialsGRPCCOntroller{svc: svc}
}

func (c *CredentialsGRPCCOntroller) SaveCredentialsData(ctx context.Context, data *api.CredentialsData) (*api.ServerResponse, error) {
	modelData := GRPCToModel(data)
	id := 1
	err := c.svc.SaveCredentialsData(ctx, modelData, id)
	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  0,
		Message: "Successfully saved data",
	}, nil
}

func (c CredentialsGRPCCOntroller) SaveBankingData(ctx context.Context, data *api.BankingCardData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")
}

func (c CredentialsGRPCCOntroller) SaveTextData(ctx context.Context, data *api.TextData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) SaveBinaryData(ctx context.Context, data *api.BinaryData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetCredentialsData(ctx context.Context, data *api.CredentialsDataRequest) (*api.CredentialsData, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetBankingCardData(ctx context.Context, data *api.BankingCardDataRequest) (*api.BankingCardData, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetTextData(ctx context.Context, data *api.TextDataRequest) (*api.TextData, error) {
	return nil, errors.New("not implemented")

}

func (c CredentialsGRPCCOntroller) GetBinaryData(ctx context.Context, data *api.BinaryDataRequest) (*api.BinaryData, error) {
	return nil, errors.New("not implemented")

}
