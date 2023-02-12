package controllers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type AuthGRPCController struct {
	authService service.Auth
}

func NewAuthGRPCController(svc service.Auth) *AuthGRPCController {
	return &AuthGRPCController{
		authService: svc,
	}
}

func (c *AuthGRPCController) RegisterUser(ctx context.Context,
	in *api.UserRegistration,
) (*api.ServerAuthResponse, error) {
	user := model.User{
		Username: in.Login,
		Password: in.Password,
		Name:     in.Name,
	}
	registeredUser, err := c.authService.RegisterUser(ctx, &user)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	token, err := c.authService.GenerateToken(registeredUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	header := metadata.Pairs("jwt", token)
	grpc.SendHeader(ctx, header)
	response := api.ServerAuthResponse{
		Message: "Successfully registered",
		Code:    http.StatusOK,
	}
	return &response, nil
}

func (c *AuthGRPCController) AuthenticateUser(ctx context.Context,
	in *api.UserAuthentication,
) (*api.ServerAuthResponse, error) {
	user := model.User{
		Username: in.Login,
		Password: in.Password,
	}
	authUser, err := c.authService.AuthenticateUser(ctx, &user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	token, err := c.authService.GenerateToken(authUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	header := metadata.Pairs("jwt", token)
	grpc.SendHeader(ctx, header)
	return &api.ServerAuthResponse{
		Message: "Successfully authorized",
		Code:    http.StatusOK,
	}, nil
}
