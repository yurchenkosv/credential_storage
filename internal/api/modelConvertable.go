package api

import (
	"github.com/yurchenkosv/credential_storage/internal/model"
	"strconv"
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
	modelData.Name = x.GetName()
	modelData.CVV = strconv.Itoa(int(x.GetCvv()))
	modelData.ValidUntil = x.GetValidTill()
	modelData.Number = strconv.Itoa(int(x.GetNumber()))
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

func (x *BinaryData) ToModel() *model.BinaryData {
	modelData := model.BinaryData{
		Data:     x.GetData(),
		Name:     x.GetName(),
		Metadata: convertMetadata(x.GetMetadata()),
	}
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
