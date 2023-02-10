package clients

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

func TestCredentialsStorageGRPCClient_AuthenticateUser(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
	}
	type args struct {
		ctx      context.Context
		login    string
		password string
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
			c := CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			got, err := c.AuthenticateUser(tt.args.ctx, tt.args.login, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AuthenticateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsStorageGRPCClient_GetData(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Credentials
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			got, err := c.GetData(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsStorageGRPCClient_RegisterUser(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
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
			c := &CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			got, err := c.RegisterUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsStorageGRPCClient_SendBankCard(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
	}
	type args struct {
		ctx  context.Context
		data model.BankingCardData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			if err := c.SendBankCard(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendBankCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialsStorageGRPCClient_SendBinary(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
	}
	type args struct {
		ctx  context.Context
		data model.BinaryData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			if err := c.SendBinary(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialsStorageGRPCClient_SendCredentials(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
	}
	type args struct {
		ctx  context.Context
		data model.CredentialsData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			if err := c.SendCredentials(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendCredentials() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialsStorageGRPCClient_SendText(t *testing.T) {
	type fields struct {
		opts              []grpc.CallOption
		credServiceClient api.CredentialServiceClient
		authClient        api.AuthServiceClient
	}
	type args struct {
		ctx  context.Context
		data model.TextData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CredentialsStorageGRPCClient{
				opts:              tt.fields.opts,
				credServiceClient: tt.fields.credServiceClient,
				authClient:        tt.fields.authClient,
			}
			if err := c.SendText(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCredentialsStorageGRPCClient(t *testing.T) {
	type args struct {
		connect *grpc.ClientConn
		opts    []grpc.CallOption
	}
	tests := []struct {
		name string
		args args
		want *CredentialsStorageGRPCClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCredentialsStorageGRPCClient(tt.args.connect, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCredentialsStorageGRPCClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
