package interceptors

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Interceptor struct {
	auth *jwtauth.JWTAuth
}

func NewInterceptor(auth *jwtauth.JWTAuth) *Interceptor {
	return &Interceptor{auth: auth}
}

func (i *Interceptor) TokenInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	meta, _ := metadata.FromIncomingContext(ctx)
	token := meta.Get("jwt")[0]
	_, err := jwtauth.VerifyToken(i.auth, token)
	if err != nil {
		return err
	}
	err = invoker(ctx, method, req, reply, cc, opts...)

	return err
}
