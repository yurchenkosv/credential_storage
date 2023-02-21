package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"os"
)

func addJWTToContext(jwt string) context.Context {
	meta := metadata.New(map[string]string{"jwt": jwt})
	return metadata.NewOutgoingContext(context.Background(), meta)
}

func loadTLSCredentials(filename string) (credentials.TransportCredentials, error) {
	pemServerCA, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
