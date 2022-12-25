package clients

import "github.com/yurchenkosv/credential_storage/internal/model"

type CredentialsStorageClient interface {
	GetCredentials() (model.Credentials, error)
	GetBankCard() (model.BankingCard, error)
	GetBinary() (model.BinaryData, error)
	GetText() (model.TextData, error)
	SendCredentials(credentials model.Credentials) error
	SendBankCard(card model.BankingCard) error
	SendBinary(binary model.BinaryData) error
	SendText(text model.TextData) error
}
