package clients

import "github.com/yurchenkosv/credential_storage/internal/model"

type CredentialsStorageClient interface {
	GetData() (model.Credentials, error)
	SendCredentials(credentials model.CredentialsData) error
	SendBankCard(card model.BankingCardData) error
	SendBinary(binary model.BinaryData) error
	SendText(text model.TextData) error
}
