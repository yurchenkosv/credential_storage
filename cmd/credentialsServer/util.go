package main

import (
	"crypto/tls"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials(key string, cert string) (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
