package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/go-chi/jwtauth/v5"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/credStorageErrors"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"time"
)

type Auth interface {
	RegisterUser(ctx context.Context, user *model.User) (*model.User, error)
	AuthenticateUser(ctx context.Context, user *model.User) (*model.User, error)
	GenerateToken(user *model.User) (string, error)
	GetUserFromToken(token string) (*model.User, error)
}

type AuthService struct {
	repo repository.Repository
	auth *jwtauth.JWTAuth
}

func NewAuthService(repo repository.Repository, auth *jwtauth.JWTAuth) *AuthService {
	return &AuthService{
		repo: repo,
		auth: auth,
	}
}

func (authService *AuthService) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Password = hashPW(user.Password)
	savedUser, err := authService.repo.GetUser(ctx, user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if savedUser.ID != nil {
		err := credStorageErrors.UserAlreadyExistsError{User: user.Username}
		return nil, &err
	}
	err = authService.repo.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	savedUser, _ = authService.repo.GetUser(ctx, user)
	return savedUser, nil
}

func (authService *AuthService) AuthenticateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user.Password = hashPW(user.Password)
	userFromRepo, _ := authService.repo.GetUser(ctx, user)
	if userFromRepo.ID == nil {
		err := credStorageErrors.InvalidUserError{User: user.Username}
		return nil, &err
	}
	return user, nil
}
func (authService *AuthService) GenerateToken(user *model.User) (string, error) {
	claims := map[string]interface{}{
		"user_id": *user.ID,
	}
	currentTime := time.Now()

	jwtauth.SetIssuedAt(claims, currentTime)
	jwtauth.SetExpiry(claims, currentTime.Add(5*time.Minute))
	_, token, err := authService.auth.Encode(claims)

	if err != nil {
		return "", err
	}
	return token, nil
}

func (authService *AuthService) GetUserFromToken(token string) (*model.User, error) {
	user := model.User{}
	jwt, err := jwtauth.VerifyToken(authService.auth, token)
	if err != nil {
		return nil, err
	}
	id, ok := jwt.Get("user_id")
	userIDExtracted := id.(float64)
	userID := int(userIDExtracted)
	if !ok {
		return nil, credStorageErrors.InvalidUserError{User: ""}
	}
	user.ID = &userID
	return &user, nil
}

func hashPW(pw string) string {
	pwHash := sha256.Sum256([]byte(pw))
	return base64.StdEncoding.EncodeToString(pwHash[:])
}
