package main

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/configProvider"
	"github.com/yurchenkosv/credential_storage/internal/controllers"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	tokenAuth *jwtauth.JWTAuth
)

func main() {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

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

	authSvc := service.NewAuthService(repo)

	grpcAuthController := controllers.NewAuthGRPCController(authSvc, tokenAuth)
	grpcServer := grpc.NewServer()
	api.RegisterAuthServiceServer(grpcServer, grpcAuthController)

	listener, err := net.Listen("tcp", config.GetConfig().ListenGRPC)
	if err != nil {
		log.Fatal(err)
	}

	go func(listener net.Listener) {
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Error(err)
		}
	}(listener)

	<-osSignal
	grpcServer.GracefulStop()

	//router := routers.NewRouter(repo, tokenAuth)
	//
	//log.Fatal(http.ListenAndServe(config.GetConfig().Listen, router))

}
