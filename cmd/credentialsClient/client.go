package main

import (
	"context"
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/configProvider"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"github.com/yurchenkosv/credential_storage/internal/view"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cfg, err := configProvider.NewClientConfigProvider()
	ctx := context.Background()
	if err != nil {
		log.Fatal("config processing error: ", err)
	}
	dialOption := grpc.WithTransportCredentials(credentials.NewTLS(
		&tls.Config{
			InsecureSkipVerify: true,
		}),
	)
	conn, err := grpc.Dial(cfg.GetConfig().ServerAddress, dialOption)
	if err != nil {
		log.Fatal("cannot create connection to server ", err)
	}
	client := clients.NewCredentialsStorageGRPCClient(conn)
	authSvc := service.NewClientAuthService(client)

	if cfg.GetConfig().RegisterUser {
		user := model.User{
			Username: cfg.GetConfig().Login,
			Password: cfg.GetConfig().Password,
			Name:     cfg.GetConfig().Name,
		}
		_, err = authSvc.Register(ctx, user)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = authSvc.Authenticate(ctx, cfg.GetConfig().Login, cfg.GetConfig().Password)
		if err != nil {
			log.Fatal("cannot authenticate on server ", err)
		}
	}

	credSvc := service.NewClientCredentialsService(ctx, client)
	tui := view.NewTUI(credSvc, ctx)
	if err = tui.RunApp(); err != nil {
		log.Fatal(err)
	}
}
