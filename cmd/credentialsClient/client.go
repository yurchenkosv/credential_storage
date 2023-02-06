package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/configProvider"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"github.com/yurchenkosv/credential_storage/internal/view"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := configProvider.NewClientConfigProvider()
	//ctx := context.Background()
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
	//jwt, err := authSvc.Authenticate(ctx, cfg.GetConfig().Login, cfg.GetConfig().Password)
	//ctx = context.WithValue(ctx, "jwt", jwt)
	//if err != nil {
	//	log.Fatal("cannot authenticate on server ", err)
	//}
	//credSvc := service.NewClientCredentialsService(ctx, client)
	//_, err = credSvc.GetData()
	//creds, err := credSvc.GetData()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for idx, cred := range creds {
	//	log.Infof("cred %d, value %v\n", idx, cred.Name)
	//}
	//log.Infof("returned creds %v", creds)

	tuiView := view.NewTuiView(authSvc)
	mainApp := tuiView.DrawMainApp
	app := tuiView.GetApp()
	err = app.SetRoot(mainApp(), true).Run()
	if err != nil {
		log.Fatal(err)
	}
}
