package controllers

import (
	"context"
	"errors"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/contextKeys"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"time"
)

type CredentialsGRPCController struct {
	svc     *service.CredentialsService
	authSvc service.Auth
}

func NewGophkeeperController(svc *service.CredentialsService) *CredentialsGRPCController {
	return &CredentialsGRPCController{svc: svc}
}

func (c *CredentialsGRPCController) SaveCredentialsData(ctx context.Context, data *api.CredentialsData) (*api.ServerResponse, error) {
	modelData := data.ToModel()
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	err := c.svc.SaveCredentialsData(ctx, modelData, id)
	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  0,
		Message: "Successfully saved data",
	}, nil
}

func (c CredentialsGRPCController) SaveBankingData(ctx context.Context, data *api.BankingCardData) (*api.ServerResponse, error) {
	modelData, err := data.ToModel()
	if err != nil {
		return nil, err
	}
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	err = c.svc.SaveBankingCardData(ctx, modelData, id)
	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  0,
		Message: "Successfully saved data",
	}, nil
}

func (c CredentialsGRPCController) SaveTextData(ctx context.Context, data *api.TextData) (*api.ServerResponse, error) {
	modelData := data.ToModel()
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	err := c.svc.SaveTextData(ctx, modelData, id)
	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  0,
		Message: "Successfully saved data",
	}, nil
}

func (c CredentialsGRPCController) SaveBinaryData(ctx context.Context, data *api.BinaryData) (*api.ServerResponse, error) {
	return nil, errors.New("not implemented")

}
func (c CredentialsGRPCController) GetData(ctx context.Context, data *api.AllDataRequest) (*api.SecretDataList, error) {
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	secretDataList := api.SecretDataList{}
	creds, err := c.svc.GetAllUserCredentials(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, secret := range creds {
		msg := api.SecretsDataResponse{}
		protoBank := &api.BankingCardData{
			Number:         int32(secret.BankingCardData.Number),
			ValidTill:      secret.BankingCardData.ValidUntil.Format(time.RFC3339),
			CardholderName: secret.BankingCardData.CardholderName,
			Cvv:            int32(secret.BankingCardData.CVV),
			Metadata:       nil,
		}
		protoCred := &api.CredentialsData{
			Login:    secret.CredentialsData.Login,
			Password: secret.CredentialsData.Password,
			Metadata: nil,
		}
		protoText := &api.TextData{
			Data:     secret.TextData.Data,
			Metadata: nil,
		}
		msg.BankingData = protoBank
		msg.CredentialsData = protoCred
		msg.TextData = protoText
		secretDataList.Secrets = append(secretDataList.Secrets, &msg)
	}
	return nil, errors.New("not implemented")
}
