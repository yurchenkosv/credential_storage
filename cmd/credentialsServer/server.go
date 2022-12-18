package main

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/yurchenkosv/credential_storage/internal/configProvider"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"github.com/yurchenkosv/credential_storage/internal/routers"
	"log"
	"net/http"
)

var (
	tokenAuth *jwtauth.JWTAuth
)

func main() {
	config, err := configProvider.NewServerConfigProvider()
	if err != nil {
		log.Fatal(err)
	}
	repo, err := repository.NewPostgresRepo(config.GetConfig().DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	err = repo.MigrateDB("internal/migrations")
	if err != nil {
		log.Fatal(err)
	}

	tokenAuth = jwtauth.New("HS256", []byte(config.GetConfig().JWTSecret), nil)

	router := routers.NewRouter(repo, tokenAuth)

	log.Fatal(http.ListenAndServe(config.GetConfig().Listen, router))

}
