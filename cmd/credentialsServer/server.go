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
	"github.com/yurchenkosv/credential_storage/internal/interceptors"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	tokenAuth *jwtauth.JWTAuth
)

func main() {
	log.SetLevel(log.WarnLevel)
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

	tlsCredentials, err := loadTLSCredentials(config.GetConfig().PrivateKeyLocation, config.GetConfig().CertLocation)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	tokenAuth = jwtauth.New("HS256", []byte(config.GetConfig().JWTSecret), nil)
	authSvc := service.NewAuthService(repo, tokenAuth)
	binaryRepo := repository.NewLocalBinaryRepository(config.GetConfig().BinaryLocalStorageLocation)
	credentialsSvc, err := service.NewProxyEncryptedCredentialService(repo, binaryRepo, config.GetConfig().EncryptionSecret)
	if err != nil {
		log.Fatal(err)
	}
	authInterceptor := interceptors.NewAuthInterceptor(authSvc)

	grpcAuthController := controllers.NewAuthGRPCController(authSvc)
	credentialsController := controllers.NewCredentialsGRPCController(credentialsSvc)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor.JWTInterceptor), grpc.Creds(tlsCredentials))

	api.RegisterAuthServiceServer(grpcServer, grpcAuthController)
	api.RegisterCredentialServiceServer(grpcServer, credentialsController)

	listener, err := net.Listen("tcp", config.GetConfig().ListenGRPC)
	if err != nil {
		log.Fatal(err)
	}

	go func(listener net.Listener) {
		log.Info("credentials server started")
		err = grpcServer.Serve(listener)
		if err != nil {
			log.Error(err)
		}
	}(listener)

	<-osSignal
	grpcServer.GracefulStop()
}
