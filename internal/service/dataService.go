package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type DataService interface {
	SaveCredentialsData(ctx context.Context, data *model.CredentialsData, userID int) error
	SaveBankingCardData(ctx context.Context, data *model.BankingCardData, userID int) error
	SaveTextData(ctx context.Context, data *model.TextData, userID int) error
	SaveBinaryData(ctx context.Context, data *model.BinaryData, userID int) error
	GetCredentialsByName(ctx context.Context, credName string, userID int) ([]model.CredentialsData, error)
	GetAllUserCredentials(ctx context.Context, userID int) ([]model.Credentials, error)
}
