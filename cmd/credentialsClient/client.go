package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/configProvider"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"github.com/yurchenkosv/credential_storage/internal/view"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	cfg, err := configProvider.NewClientConfigProvider()
	ctx := context.Background()
	if err != nil {
		log.Fatal("cannot get config ", err)
	}
	dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(cfg.GetConfig().ServerAddress, dialOption)
	if err != nil {
		log.Fatal("cannot create connection to server ", err)
	}
	client := clients.NewCredentialsStorageGRPCClient(conn)
	authSvc := service.NewClientAuthService(client)
	jwt, err := authSvc.Authenticate(ctx, cfg.GetConfig().Login, cfg.GetConfig().Password)
	meta := metadata.New(map[string]string{"jwt": jwt})
	ctx = metadata.NewOutgoingContext(ctx, meta)
	if err != nil {
		log.Fatal("cannot authenticate on server ", err)
	}
	credSvc := service.NewClientCredentialsService(ctx, client)
	tui := view.NewTUI(credSvc, ctx)
	if err = tui.RunApp(); err != nil {
		log.Fatal(err)
	}
}
