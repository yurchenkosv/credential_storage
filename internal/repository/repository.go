package repository

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type Repository interface {
	GetUser(ctx context.Context, user *model.User) (*model.User, error)
	SaveUser(ctx context.Context, user *model.User) error

	SaveCredentialsData(ctx context.Context, creds *model.CredentialsData, userID int) error
	SaveBankingCardData(ctx context.Context, data *model.BankingCardData, userID int) error
	SaveTextData(ctx context.Context, data *model.TextData, userID int) error
	SaveBinaryData(ctx context.Context, data *model.BinaryData, userID int, link string) error

	UpdateBankingCardData(ctx context.Context, data model.BankingCardData, userID int) error
	UpdateCredentialsData(ctx context.Context, data model.CredentialsData, userID int) error
	UpdateTextData(ctx context.Context, data model.TextData, userID int) error
	UpdateBinaryData(ctx context.Context, data model.BinaryData, userID int) error

	DeleteData(ctx context.Context, data model.Credentials, userID int) error

	GetCredentialsByUserID(ctx context.Context, userID int) ([]model.Credentials, error)
	GetCredentialsByName(ctx context.Context, name string, userID int) ([]model.CredentialsData, error)
}
