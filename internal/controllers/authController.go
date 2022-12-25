package controllers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/credStorageErrors"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"io"
	"net/http"
)

type AuthController struct {
	authService service.Auth
	jwtAuth     *jwtauth.JWTAuth
}

func NewAuthController(authService *service.Auth, jwtAuth *jwtauth.JWTAuth) AuthController {
	return AuthController{
		authService: *authService,
		jwtAuth:     jwtAuth,
	}
}

func (h AuthController) HandleUserRegistration(writer http.ResponseWriter, request *http.Request) {
	user, err := h.parseForUser(request)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedUser, err := h.authService.RegisterUser(request.Context(), user)
	if err != nil {
		switch e := err.(type) {
		case *credStorageErrors.UserAlreadyExistsError:
			log.Error(err)
			writer.WriteHeader(http.StatusConflict)
			return
		default:
			log.Error("error creating user", e)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	token, err := SetToken(*updatedUser, h.jwtAuth)
	if err != nil {
		log.Error("error setting token for user:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Add("jwt", token)
	writer.Header().Add("Set-Cookie", "jwt="+token)
	writer.WriteHeader(http.StatusOK)
}

func (h AuthController) HanldeUserLogin(writer http.ResponseWriter, request *http.Request) {
	user, err := h.parseForUser(request)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	updatedUser, err := h.authService.AuthenticateUser(request.Context(), user)
	if err != nil {
		switch e := err.(type) {

		case *credStorageErrors.InvalidUserError:
			log.Error(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		default:
			log.Error("error during user authentication ", e)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	token, err := SetToken(*updatedUser, h.jwtAuth)
	if err != nil {
		log.Error("error setting token for user:", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Add("jwt", token)
	writer.Header().Add("Set-Cookie", "jwt="+token)
	writer.WriteHeader(http.StatusOK)
	writer.WriteHeader(http.StatusOK)
}

func (h AuthController) parseForUser(request *http.Request) (*model.User, error) {
	var user model.User

	data, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		log.Error(err)
	}
	return &user, nil
}
