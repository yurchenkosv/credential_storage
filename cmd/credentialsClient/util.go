package main

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func addJWTToContext(jwt string) context.Context {
	meta := metadata.New(map[string]string{"jwt": jwt})
	return metadata.NewOutgoingContext(context.Background(), meta)
}
