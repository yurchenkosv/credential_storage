package repository

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/model"
)

type Repository interface {
	Save() error
	GetUser(ctx context.Context, user *model.User) (*model.User, error)
	SaveUser(ctx context.Context, user *model.User) error
}
