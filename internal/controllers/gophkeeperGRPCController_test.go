package controllers

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/contextKeys"
	mock_service "github.com/yurchenkosv/credential_storage/internal/mockService"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestCredentialsGRPCController_GetData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data []model.Credentials)
	type args struct {
		ctx   context.Context
		data  *api.AllDataRequest
		creds []model.Credentials
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.SecretDataList
		wantErr      bool
	}{
		{
			name: "should return all user creds in grpc format",
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data []model.Credentials) {
				s.EXPECT().GetAllUserCredentials(ctx, 1).Return(data, nil)
			},
			args: args{
				ctx:  context.WithValue(context.Background(), contextKeys.UserIDContexKey("user_id"), 1),
				data: &api.AllDataRequest{},
				creds: []model.Credentials{
					{
						ID:   1,
						Name: "testCredentials",
						CredentialsData: &model.CredentialsData{
							ID:       1,
							Login:    "test_login",
							Password: "test_pwd",
						},
						BankingCardData: &model.BankingCardData{
							ID:             1,
							Number:         "382492873",
							ValidUntil:     "10/25",
							CardholderName: "test holder",
							CVV:            "423",
						},
						TextData: &model.TextData{
							ID:   1,
							Data: "text",
						},
						BinaryData: &model.BinaryData{
							ID:   1,
							Data: []byte("test"),
							Link: "/tmp/link",
						},
						Metadata: []model.Metadata{
							{
								Value: "test",
							},
						},
					},
				},
			},
			want: &api.SecretDataList{
				Secrets: []*api.SecretsData{
					{
						Name: "testCredentials",
						CredentialsData: &api.CredentialsData{
							Login:    "test_login",
							Password: "test_pwd",
							Id:       1,
						},
						BankingData: &api.BankingCardData{
							Number:         int32(382492873),
							ValidTill:      "10/25",
							CardholderName: "test holder",
							Cvv:            423,
							Id:             1,
						},
						TextData: &api.TextData{
							Data: "text",
							Id:   1,
						},
						BinaryData: &api.BinaryData{
							Id:   1,
							Data: []byte("test"),
						},
						Metadata: []string{"test"},
						Id:       1,
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
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.creds)
			c := NewCredentialsGRPCController(svc)
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
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.BankingCardData)
	type args struct {
		ctx      context.Context
		data     *api.BankingCardData
		bankData *model.BankingCardData
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerResponse
		wantErr      bool
	}{
		{
			name: "should save bank data",
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data *model.BankingCardData) {
				s.EXPECT().SaveBankingCardData(ctx, data, 1).Return(nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), contextKeys.UserIDContexKey("user_id"), 1),
				data: &api.BankingCardData{
					Number:         147433,
					ValidTill:      "10/26",
					CardholderName: "card holder",
					Cvv:            123,
					Name:           "test data",
				},
				bankData: &model.BankingCardData{
					Name:           "test data",
					Number:         "147433",
					ValidUntil:     "10/26",
					CardholderName: "card holder",
					CVV:            "123",
				},
			},
			want: &api.ServerResponse{
				Status:  200,
				Message: "Successfully saved data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.bankData)
			c := NewCredentialsGRPCController(svc)
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
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, reader io.Reader, data *model.BinaryData)
	type args struct {
		ctx        context.Context
		data       *api.BinaryData
		binaryData *model.BinaryData
		reader     io.Reader
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerResponse
		wantErr      bool
	}{
		{
			name: "should save binary",
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, reader io.Reader, data *model.BinaryData) {
				s.EXPECT().SaveBinaryData(ctx, reader, data, 1).Return(nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), contextKeys.UserIDContexKey("user_id"), 1),
				data: &api.BinaryData{
					Data: []byte("test"),
					Name: "test",
				},
				binaryData: &model.BinaryData{
					Data: []byte("test"),
					Name: "test",
				},
				reader: bytes.NewReader([]byte("test")),
			},
			want: &api.ServerResponse{
				Status:  200,
				Message: "Successfully saved data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.reader, tt.args.binaryData)
			c := NewCredentialsGRPCController(svc)
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
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.CredentialsData)
	type args struct {
		ctx      context.Context
		data     *api.CredentialsData
		credData *model.CredentialsData
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerResponse
		wantErr      bool
	}{
		{
			name: "should save credentials data",
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data *model.CredentialsData) {
				s.EXPECT().SaveCredentialsData(ctx, data, 1).Return(nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), contextKeys.UserIDContexKey("user_id"), 1),
				data: &api.CredentialsData{
					Login:    "test",
					Password: "test_pwd",
					Name:     "test",
				},
				credData: &model.CredentialsData{
					Name:     "test",
					Login:    "test",
					Password: "test_pwd",
				},
			},
			want: &api.ServerResponse{
				Status:  200,
				Message: "Successfully saved data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.credData)
			c := NewCredentialsGRPCController(svc)
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
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.TextData)
	type args struct {
		ctx      context.Context
		data     *api.TextData
		textData *model.TextData
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerResponse
		wantErr      bool
	}{
		{
			name: "should save text data",
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data *model.TextData) {
				s.EXPECT().SaveTextData(ctx, data, 1).Return(nil)
			},
			args: args{
				ctx: context.WithValue(context.Background(), contextKeys.UserIDContexKey("user_id"), 1),
				data: &api.TextData{
					Data: "text",
					Name: "test text",
				},
				textData: &model.TextData{
					Name: "test text",
					Data: "text",
				},
			},
			want: &api.ServerResponse{
				Status:  200,
				Message: "Successfully saved data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.textData)
			c := NewCredentialsGRPCController(svc)
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

func TestCredentialsGRPCController_DeleteData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data model.Credentials)
	type args struct {
		ctx      context.Context
		data     *api.SecretsData
		mockData model.Credentials
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		want         *api.ServerResponse
		wantErr      bool
	}{
		{
			name: "should delete data",
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data model.Credentials) {
				userID := 1
				s.EXPECT().DeleteCredential(ctx, data, userID).Return(nil)
			},
			args: args{
				ctx:      context.WithValue(context.Background(), contextKeys.UserIDContexKey("user_id"), 1),
				data:     &api.SecretsData{CredentialsData: &api.CredentialsData{Id: 1}},
				mockData: model.Credentials{CredentialsData: &model.CredentialsData{ID: 1}},
			},
			want: &api.ServerResponse{
				Status:  http.StatusOK,
				Message: "Successfully saved data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.mockData)
			c := &CredentialsGRPCController{
				svc: svc}
			got, err := c.DeleteData(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteData() got = %v, want %v", got, tt.want)
			}
		})
	}
}
