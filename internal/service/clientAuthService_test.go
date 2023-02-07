package service

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/clients"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"testing"
)

func TestClientAuthService_Authenticate(t *testing.T) {
	type fields struct {
		client clients.CredentialsStorageClient
	}
	type args struct {
		ctx   context.Context
		login string
		pwd   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientAuthService{
				client: tt.fields.client,
			}
			got, err := s.Authenticate(tt.args.ctx, tt.args.login, tt.args.pwd)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientAuthService_Register(t *testing.T) {
	type fields struct {
		client clients.CredentialsStorageClient
	}
	type args struct {
		ctx  context.Context
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ClientAuthService{
				client: tt.fields.client,
			}
			got, err := s.Register(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
