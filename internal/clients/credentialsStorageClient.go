package clients

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type CredentialsStorageClient interface {
	GetData(ctx context.Context) ([]model.Credentials, error)
	SendCredentials(ctx context.Context, credentials model.CredentialsData) error
	SendBankCard(ctx context.Context, card model.BankingCardData) error
	SendBinary(ctx context.Context, binary model.BinaryData) error
	SendText(ctx context.Context, text model.TextData) error
	AuthenticateUser(ctx context.Context, login string, password string) (string, error)
	RegisterUser(ctx context.Context, user model.User) (string, error)
}
