package controllers

import (
	"bytes"
	"context"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/contextKeys"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"net/http"
	"strconv"
)

type CredentialsGRPCController struct {
	svc service.DataService
}

func NewCredentialsGRPCController(svc service.DataService) *CredentialsGRPCController {
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
		Status:  http.StatusOK,
		Message: "Successfully saved data",
	}, nil
}

func (c *CredentialsGRPCController) SaveBankingData(ctx context.Context, data *api.BankingCardData) (*api.ServerResponse, error) {
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
		Status:  http.StatusOK,
		Message: "Successfully saved data",
	}, nil
}

func (c *CredentialsGRPCController) SaveTextData(ctx context.Context, data *api.TextData) (*api.ServerResponse, error) {
	modelData := data.ToModel()
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	err := c.svc.SaveTextData(ctx, modelData, id)
	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  200,
		Message: "Successfully saved data",
	}, nil
}

func (c *CredentialsGRPCController) SaveBinaryData(ctx context.Context, data *api.BinaryData) (*api.ServerResponse, error) {
	modelData := data.ToModel()
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	reader := bytes.NewReader(data.Data)
	err := c.svc.SaveBinaryData(ctx, reader, modelData, id)
	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  http.StatusOK,
		Message: "Successfully saved data",
	}, nil
}

func (c *CredentialsGRPCController) GetData(ctx context.Context, data *api.AllDataRequest) (*api.SecretDataList, error) {
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	creds, err := c.svc.GetAllUserCredentials(ctx, id)
	secretDataList := &api.SecretDataList{}
	if err != nil {
		return nil, err
	}
	for _, secret := range creds {
		msg := api.SecretsData{Name: secret.Name, Id: int32(secret.ID)}
		if secret.BankingCardData != nil {
			num, _ := strconv.ParseInt(secret.BankingCardData.Number, 10, 64)
			cvv, _ := strconv.ParseInt(secret.BankingCardData.CVV, 10, 64)
			protoBank := &api.BankingCardData{
				Id:             int32(secret.BankingCardData.ID),
				Number:         int32(num),
				ValidTill:      secret.BankingCardData.ValidUntil,
				CardholderName: secret.BankingCardData.CardholderName,
				Cvv:            int32(cvv),
				Metadata:       nil,
			}
			msg.BankingData = protoBank
		}
		if secret.CredentialsData != nil {
			protoCred := &api.CredentialsData{
				Id:       int32(secret.CredentialsData.ID),
				Login:    secret.CredentialsData.Login,
				Password: secret.CredentialsData.Password,
				Metadata: nil,
			}
			msg.CredentialsData = protoCred
		}
		if secret.TextData != nil {
			protoText := &api.TextData{
				Id:       int32(secret.TextData.ID),
				Data:     secret.TextData.Data,
				Metadata: nil,
			}
			msg.TextData = protoText
		}
		if secret.BinaryData != nil {
			protoBinary := &api.BinaryData{
				Id:   int32(secret.BinaryData.ID),
				Data: secret.BinaryData.Data,
			}
			msg.BinaryData = protoBinary
		}
		for _, meta := range secret.Metadata {
			msg.Metadata = append(msg.Metadata, meta.Value)
		}
		secretDataList.Secrets = append(secretDataList.Secrets, &msg)
	}
	return secretDataList, nil
}

func (c *CredentialsGRPCController) DeleteData(ctx context.Context, data *api.SecretsData) (*api.ServerResponse, error) {
	id := ctx.Value(contextKeys.UserIDContexKey("user_id")).(int)
	credData := model.Credentials{}
	if data.BinaryData != nil {
		credData.BinaryData = &model.BinaryData{ID: int(data.BinaryData.Id)}
	}
	if data.CredentialsData != nil {

		credData.CredentialsData = &model.CredentialsData{ID: int(data.CredentialsData.Id)}
	}
	if data.BankingData != nil {

		credData.BankingCardData = &model.BankingCardData{ID: int(data.BankingData.Id)}
	}
	if data.TextData != nil {
		credData.TextData = &model.TextData{ID: int(data.TextData.Id)}

	}

	for _, meta := range data.Metadata {
		credData.Metadata = append(credData.Metadata, model.Metadata{Value: meta})
	}
	err := c.svc.DeleteCredential(ctx, credData, id)

	if err != nil {
		return nil, err
	}
	return &api.ServerResponse{
		Status:  http.StatusOK,
		Message: "Successfully saved data",
	}, nil
}
