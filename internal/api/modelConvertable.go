package api

import (
	"github.com/yurchenkosv/credential_storage/internal/model"
	"time"
)

func (x *CredentialsData) ToModel() *model.CredentialsData {
	modelData := model.CredentialsData{}
	modelData.Name = x.GetName()
	modelData.Login = x.GetLogin()
	modelData.Password = x.GetPassword()
	for _, meta := range x.GetMetadata() {
		metadata := model.Metadata{
			Value: meta,
		}
		modelData.Metadata = append(modelData.Metadata, metadata)
	}
	return &modelData
}

func (x *BankingCardData) ToModel() (*model.BankingCardData, error) {
	modelData := model.BankingCardData{}

	timeValidTill, err := time.Parse(time.RFC3339, x.GetValidTill())
	if err != nil {
		return nil, err
	}

	modelData.Name = x.GetName()
	modelData.CVV = int(x.GetCvv())
	modelData.ValidUntil = timeValidTill
	modelData.Number = int(x.GetNumber())
	modelData.CardholderName = x.GetCardholderName()
	modelData.Metadata = convertMetadata(x.Metadata)
	return &modelData, nil
}

func (x *TextData) ToModel() *model.TextData {
	modelData := model.TextData{}
	modelData.Name = x.GetName()
	modelData.Data = x.GetData()
	modelData.Metadata = convertMetadata(x.Metadata)
	return &modelData
}

func convertMetadata(m []string) []model.Metadata {
	var metadata []model.Metadata
	for _, meta := range m {
		data := model.Metadata{}
		data.Value = meta
		metadata = append(metadata, data)
	}
	return metadata
}
