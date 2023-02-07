package controllers

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/service"
	"reflect"
	"testing"
)

func TestCredentialsGRPCController_GetData(t *testing.T) {
	type fields struct {
		svc     service.DataService
		authSvc service.Auth
	}
	type args struct {
		ctx  context.Context
		data *api.AllDataRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.SecretDataList
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CredentialsGRPCController{
				svc:     tt.fields.svc,
				authSvc: tt.fields.authSvc,
			}
			got, err := c.GetData(tt.args.ctx, tt.args.data)
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

func TestCredentialsGRPCController_SaveBankingData(t *testing.T) {
	type fields struct {
		svc     service.DataService
		authSvc service.Auth
	}
	type args struct {
		ctx  context.Context
		data *api.BankingCardData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CredentialsGRPCController{
				svc:     tt.fields.svc,
				authSvc: tt.fields.authSvc,
			}
			got, err := c.SaveBankingData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveBankingData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveBankingData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsGRPCController_SaveBinaryData(t *testing.T) {
	type fields struct {
		svc     service.DataService
		authSvc service.Auth
	}
	type args struct {
		ctx  context.Context
		data *api.BinaryData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CredentialsGRPCController{
				svc:     tt.fields.svc,
				authSvc: tt.fields.authSvc,
			}
			got, err := c.SaveBinaryData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveBinaryData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveBinaryData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsGRPCController_SaveCredentialsData(t *testing.T) {
	type fields struct {
		svc     service.DataService
		authSvc service.Auth
	}
	type args struct {
		ctx  context.Context
		data *api.CredentialsData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CredentialsGRPCController{
				svc:     tt.fields.svc,
				authSvc: tt.fields.authSvc,
			}
			got, err := c.SaveCredentialsData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveCredentialsData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsGRPCController_SaveTextData(t *testing.T) {
	type fields struct {
		svc     service.DataService
		authSvc service.Auth
	}
	type args struct {
		ctx  context.Context
		data *api.TextData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.ServerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CredentialsGRPCController{
				svc:     tt.fields.svc,
				authSvc: tt.fields.authSvc,
			}
			got, err := c.SaveTextData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveTextData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveTextData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGophkeeperController(t *testing.T) {
	type args struct {
		svc service.DataService
	}
	tests := []struct {
		name string
		args args
		want *CredentialsGRPCController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGophkeeperController(tt.args.svc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGophkeeperController() = %v, want %v", got, tt.want)
			}
		})
	}
}
