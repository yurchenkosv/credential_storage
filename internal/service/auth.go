package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/credStorageErrors"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
)

type Auth interface {
	RegisterUser(ctx context.Context, user *model.User) (*model.User, error)
	AuthenticateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) Auth {
	return AuthService{repo: repo}
}

func (auth AuthService) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Password = hashPW(user.Password)
	savedUser, err := auth.repo.GetUser(ctx, user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if savedUser.ID != nil {
		err := credStorageErrors.UserAlreadyExistsError{User: user.Username}
		return nil, &err
	}
	err = auth.repo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	savedUser, _ = auth.repo.GetUser(ctx, user)
	return savedUser, nil
}

func (auth AuthService) AuthenticateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Password = hashPW(user.Password)
	userFromRepo, _ := auth.repo.GetUser(ctx, user)
	if userFromRepo.ID == nil {
		err := credStorageErrors.InvalidUserError{User: user.Username}
		return nil, &err
	}
	return user, nil
}

func hashPW(pw string) string {
	pwHash := sha256.Sum256([]byte(pw))
	return base64.StdEncoding.EncodeToString(pwHash[:])
}
