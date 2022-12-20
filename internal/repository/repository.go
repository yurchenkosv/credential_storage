package repository

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type Repository interface {
	GetUser(ctx context.Context, user *model.User) (*model.User, error)
	SaveUser(ctx context.Context, user *model.User) error

	SaveCredentialsData(ctx context.Context, creds *model.Credentials, userID int) error
	SaveBankingCardData(ctx context.Context, creds *model.Credentials, userID int) error
	SaveTextData(ctx context.Context, creds *model.Credentials, userID int) error
	SaveBinaryData(ctx context.Context, creds *model.Credentials, userID int, link string) error

	GetCredentialsByUserID(ctx context.Context, userID int) (*model.Credentials, error)
	GetCredentialsByName(ctx context.Context, name string) (*model.Credentials, error)
}
