package interceptors

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

func TestAuthInterceptor_JWTInterceptor(t *testing.T) {
	type fields struct {
		authSvc service.Auth
	}
	type args struct {
		ctx     context.Context
		req     interface{}
		info    *grpc.UnaryServerInfo
		handler grpc.UnaryHandler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AuthInterceptor{
				authSvc: tt.fields.authSvc,
			}
			got, err := i.JWTInterceptor(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("JWTInterceptor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWTInterceptor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthInterceptor(t *testing.T) {
	type args struct {
		svc service.Auth
	}
	tests := []struct {
		name string
		args args
		want *AuthInterceptor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthInterceptor(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthInterceptor() = %v, want %v", got, tt.want)
			}
		})
	}
}
