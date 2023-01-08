package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/configProvider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ()

func main() {

	cfg, err := configProvider.NewClientConfigProvider()
	if err != nil {
		log.Fatal("cannot get config ", err)
	}
	dialOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(cfg.GetConfig().ServerAddress, dialOption)
	if err != nil {
		log.Fatal("cannot create connection to server ", err)
	}
	clients.NewCredentialsStorageGRPCClient(conn)

}
