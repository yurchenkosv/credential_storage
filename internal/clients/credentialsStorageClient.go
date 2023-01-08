package clients

import "github.com/yurchenkosv/credential_storage/internal/model"

type CredentialsStorageClient interface {
	GetCredentials() (model.CredentialsData, error)
	GetBankCard() (model.BankingCardData, error)
	GetBinary() (model.BinaryData, error)
	GetText() (model.TextData, error)
	SendCredentials(credentials model.CredentialsData) error
	SendBankCard(card model.BankingCardData) error
	SendBinary(binary model.BinaryData) error
	SendText(text model.TextData) error
}
