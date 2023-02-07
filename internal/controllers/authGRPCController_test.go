package controllers

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"reflect"
	"testing"
)

func TestAuthGRPCController_AuthenticateUser(t *testing.T) {
	type fields struct {
		authService service.Auth
	}
	type args struct {
		ctx context.Context
		in  *api.UserAuthentication
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServerAuthResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AuthGRPCController{
				authService: tt.fields.authService,
			}
			got, err := c.AuthenticateUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthGRPCController_RegisterUser(t *testing.T) {
	type fields struct {
		authService service.Auth
	}
	type args struct {
		ctx context.Context
		in  *api.UserRegistration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServerAuthResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AuthGRPCController{
				authService: tt.fields.authService,
			}
			got, err := c.RegisterUser(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthGRPCController(t *testing.T) {
	type args struct {
		svc service.Auth
	}
	tests := []struct {
		name string
		args args
		want *AuthGRPCController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthGRPCController(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthGRPCController() = %v, want %v", got, tt.want)
			}
		})
	}
}
