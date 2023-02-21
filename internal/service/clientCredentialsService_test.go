package service

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_clients "github.com/yurchenkosv/credential_storage/internal/mockCredStorageClient"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"reflect"
	"testing"
)

func TestClientCredentialsService_GetData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, credentials []model.Credentials)
	type fields struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		want         []model.Credentials
		wantErr      bool
	}{
		{
			name: "should success with get all user data",
			fields: fields{
				ctx: context.Background(),
			},
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, credentials []model.Credentials) {
				s.EXPECT().GetData(ctx).Return(credentials, nil)
			},
			want: []model.Credentials{
				{
					ID:   1,
					Name: "test_credentials",
					CredentialsData: &model.CredentialsData{
						ID:       2,
						Login:    "test",
						Password: "test",
					},
					BankingCardData: nil,
					TextData:        nil,
					BinaryData:      nil,
					Metadata:        nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.fields.ctx, client, tt.want)
			s := ClientCredentialsService{
				client: client,
				ctx:    tt.fields.ctx,
			}
			got, err := s.GetData()
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

func TestClientCredentialsService_SendBankCard(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.BankingCardData)
	type fields struct {
		ctx context.Context
	}
	type args struct {
		card model.BankingCardData
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should success with bank data",
			fields: fields{
				ctx: context.Background(),
			},
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.BankingCardData) {
				s.EXPECT().SendBankCard(ctx, data).Return(nil)
			},
			args: args{
				card: model.BankingCardData{
					ID:             2,
					Name:           "bankData",
					Number:         "918230913290813",
					ValidUntil:     "10/22",
					CardholderName: "test user",
					CVV:            "245",
					Metadata: []model.Metadata{
						{
							ID:    3,
							Value: "test",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.fields.ctx, client, tt.args.card)
			s := &ClientCredentialsService{
				client: client,
				ctx:    tt.fields.ctx,
			}
			if err := s.SendBankCard(tt.args.card); (err != nil) != tt.wantErr {
				t.Errorf("SendBankCard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientCredentialsService_SendBinary(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.BinaryData)
	type fields struct {
		ctx context.Context
	}
	type args struct {
		data model.BinaryData
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name:   "should success with sending binary data",
			fields: fields{},
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.BinaryData) {
				s.EXPECT().SendBinary(ctx, data).Return(nil)
			},
			args: args{
				data: model.BinaryData{
					ID:   1,
					Data: []byte("test"),
					Name: "test binary",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.fields.ctx, client, tt.args.data)
			s := &ClientCredentialsService{
				client: client,
				ctx:    tt.fields.ctx,
			}
			if err := s.SendBinary(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientCredentialsService_SendCredentials(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.CredentialsData)
	type fields struct {
		ctx context.Context
	}
	type args struct {
		credentials model.CredentialsData
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should success with sending credentials",
			fields: fields{
				ctx: context.Background(),
			},
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.CredentialsData) {
				s.EXPECT().SendCredentials(ctx, data).Return(nil)
			},
			args: args{
				credentials: model.CredentialsData{
					Name:     "test",
					Login:    "test_login",
					Password: "test_pwd",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.fields.ctx, client, tt.args.credentials)
			s := &ClientCredentialsService{
				client: client,
				ctx:    tt.fields.ctx,
			}
			if err := s.SendCredentials(tt.args.credentials); (err != nil) != tt.wantErr {
				t.Errorf("SendCredentials() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClientCredentialsService_SendText(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.TextData)
	type fields struct {
		ctx context.Context
	}
	type args struct {
		data model.TextData
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should success with sending text data",
			fields: fields{
				ctx: context.Background(),
			},
			mockBehavior: func(ctx context.Context, s *mock_clients.MockCredentialsStorageClient, data model.TextData) {
				s.EXPECT().SendText(ctx, data).Return(nil)
			},
			args: args{
				data: model.TextData{
					Name:     "test text data",
					Data:     "text",
					Metadata: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			client := mock_clients.NewMockCredentialsStorageClient(ctrl)
			tt.mockBehavior(tt.fields.ctx, client, tt.args.data)
			s := &ClientCredentialsService{
				client: client,
				ctx:    tt.fields.ctx,
			}
			if err := s.SendText(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SendText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
