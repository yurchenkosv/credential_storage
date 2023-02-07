package service

import (
	"context"
	"github.com/golang/mock/gomock"
	mock_repository "github.com/yurchenkosv/credential_storage/internal/mockRepo"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"reflect"
	"testing"
)

func TestCredentialsService_GetAllUserCredentials(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, userID int, creds []model.Credentials)

	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		want         []model.Credentials
		wantErr      bool
	}{
		{
			name: "should successfully return all user credentials",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
			mockBehavior: func(ctx context.Context, s *mock_repository.MockRepository, userID int, creds []model.Credentials) {
				s.EXPECT().GetCredentialsByUserID(ctx, userID).Return(creds, nil)
			},
			want: []model.Credentials{
				{
					ID:   1,
					Name: "testCredentials",
					CredentialsData: &model.CredentialsData{
						ID:       1,
						Login:    "test_login",
						Password: "test_password",
					},
					BankingCardData: nil,
					TextData:        nil,
					BinaryData:      nil,
					Metadata: []model.Metadata{
						{
							ID:    1,
							Value: "lorem ipsum",
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
			repo := mock_repository.NewMockRepository(ctrl)
			tt.mockBehavior(tt.args.ctx, repo, tt.args.userID, tt.want)
			s := &CredentialsService{
				repo: repo,
			}
			got, err := s.GetAllUserCredentials(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUserCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUserCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsService_GetCredentialsByName(t *testing.T) {
	type mockBehavior func(ctx context.Context,
		s *mock_repository.MockRepository,
		userID int,
		credName string,
		creds []model.CredentialsData)
	type args struct {
		ctx      context.Context
		credName string
		userID   int
	}
	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		want         []model.CredentialsData
		wantErr      bool
	}{
		{
			name: "should successfuly return credentials",
			args: args{
				ctx:      context.Background(),
				credName: "testCreds",
				userID:   1,
			},
			mockBehavior: func(ctx context.Context, s *mock_repository.MockRepository, userID int,
				credName string, creds []model.CredentialsData) {
				s.EXPECT().GetCredentialsByName(ctx, credName, userID).Return(creds, nil)
			},
			want: []model.CredentialsData{
				{
					ID:       1,
					Name:     "testcreds",
					Login:    "testLogin",
					Password: "testPassword",
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
			repo := mock_repository.NewMockRepository(ctrl)
			tt.mockBehavior(tt.args.ctx, repo, tt.args.userID, tt.args.credName, tt.want)
			s := &CredentialsService{
				repo: repo,
			}
			got, err := s.GetCredentialsByName(tt.args.ctx, tt.args.credName, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentialsByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentialsByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsService_SaveBankingCardData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.BankingCardData)
	type args struct {
		ctx    context.Context
		data   *model.BankingCardData
		userID int
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should successfully save bank data",
			mockBehavior: func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.BankingCardData) {
				s.EXPECT().SaveBankingCardData(ctx, data, userID).Return(nil)
			},
			args: args{
				ctx: nil,
				data: &model.BankingCardData{
					ID:             1,
					Name:           "testBank card",
					Number:         "42208537328",
					ValidUntil:     "10/23",
					CardholderName: "test user",
					CVV:            "234",
					Metadata: []model.Metadata{
						{
							ID:    1,
							Value: "important meta1",
						},
						{
							ID:    2,
							Value: "important meta2",
						},
					},
				},
				userID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockRepository(ctrl)
			tt.mockBehavior(tt.args.ctx, repo, tt.args.userID, tt.args.data)
			s := &CredentialsService{
				repo: repo,
			}
			if err := s.SaveBankingCardData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveBankingCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialsService_SaveBinaryData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.BinaryData)
	type args struct {
		ctx    context.Context
		data   *model.BinaryData
		userID int
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "shuld return error",
			mockBehavior: func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.BinaryData) {

			},
			args: args{
				ctx: context.Background(),
				data: &model.BinaryData{
					ID:       1,
					Name:     "test",
					Link:     "/url",
					Metadata: nil,
				},
				userID: 2,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockRepository(ctrl)
			tt.mockBehavior(tt.args.ctx, repo, tt.args.userID, tt.args.data)
			s := &CredentialsService{
				repo: repo,
			}
			if err := s.SaveBinaryData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialsService_SaveCredentialsData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.CredentialsData)
	type args struct {
		ctx    context.Context
		data   *model.CredentialsData
		userID int
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should successfully save credentials",
			mockBehavior: func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.CredentialsData) {
				s.EXPECT().SaveCredentialsData(ctx, data, userID).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				data: &model.CredentialsData{
					ID:       1,
					Name:     "test credentials",
					Login:    "test_user",
					Password: "test_password",
					Metadata: nil,
				},
				userID: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockRepository(ctrl)
			tt.mockBehavior(tt.args.ctx, repo, tt.args.userID, tt.args.data)
			s := &CredentialsService{
				repo: repo,
			}
			if err := s.SaveCredentialsData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialsService_SaveTextData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.TextData)
	type args struct {
		ctx    context.Context
		data   *model.TextData
		userID int
	}
	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should successfully save text data",
			mockBehavior: func(ctx context.Context, s *mock_repository.MockRepository, userID int, data *model.TextData) {
				s.EXPECT().SaveTextData(ctx, data, userID).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				data: &model.TextData{
					ID:       1,
					Name:     "text data",
					Data:     "lorem ipsum",
					Metadata: nil,
				},
				userID: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repo := mock_repository.NewMockRepository(ctrl)
			tt.mockBehavior(tt.args.ctx, repo, tt.args.userID, tt.args.data)
			s := &CredentialsService{
				repo: repo,
			}
			if err := s.SaveTextData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
