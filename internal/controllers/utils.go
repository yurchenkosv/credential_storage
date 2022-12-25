package controllers

import (
	"context"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"time"
)

func SetToken(user model.User, auth *jwtauth.JWTAuth) (string, error) {
	claims := map[string]interface{}{
		"user_id": *user.ID,
	}
	currentTime := time.Now()

	jwtauth.SetIssuedAt(claims, currentTime)
	jwtauth.SetExpiry(claims, currentTime.Add(5*time.Minute))
	_, token, err := auth.Encode(claims)

	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserIDFromTokenContext(ctx context.Context) int {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		log.Error(err)
	}
	userID := claims["user_id"].(float64)
	return int(userID)
}

func GetUserIDFromToken(token string, auth *jwtauth.JWTAuth) (int, error) {
	decodedToken, err := auth.Decode(token)
	if err != nil {
		return 0, err
	}
	userID, ok := decodedToken.Get("user_id")
	if !ok {
		return 0, errors.New("cannot get user id from token")
	}
	id := userID.(int)
	return id, nil
}
