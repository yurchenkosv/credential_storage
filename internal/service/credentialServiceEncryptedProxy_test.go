package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"github.com/golang/mock/gomock"
	mock_service "github.com/yurchenkosv/credential_storage/internal/mockService"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"io"
	"reflect"
	"testing"
)

func createCipherBlock(key string) cipher.Block {
	initKey := sha256.Sum256([]byte(key))
	block, _ := aes.NewCipher(initKey[:])
	return block
}

func Test_encryptMetadata(t *testing.T) {
	type args struct {
		block    cipher.Block
		metadata []model.Metadata
	}
	tests := []struct {
		name string
		args args
		want []model.Metadata
	}{
		{
			name: "should successfully encrypt / decrypt metadata",
			args: args{
				metadata: []model.Metadata{
					{
						ID:    1,
						Value: "test_meta",
					},
				},
				block: createCipherBlock("test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(encryptMetadata(tt.args.metadata, tt.args.block), decryptMetadata(tt.args.metadata, tt.args.block)) {
				t.Error("encryptMetadata() != decryptMetadata")
			}
		})
	}
}

func Test_initCypher(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    cipher.Block
		wantErr bool
	}{
		{
			name: "should sucessfully create cipher block",
			args: args{
				key: "test",
			},
			want:    createCipherBlock("test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initCypher(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("initCypher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initCypher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialServiceEncryptedProxy_GetAllUserCredentials(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.TextData)
	type fields struct {
		cypherBlock cipher.Block
	}
	type args struct {
		ctx    context.Context
		userID int
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
			s := &CredentialServiceEncryptedProxy{
				cypherBlock: tt.fields.cypherBlock,
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

func TestCredentialServiceEncryptedProxy_GetCredentialsByName(t *testing.T) {
	type fields struct {
		cypherBlock cipher.Block
		svc         DataService
	}
	type args struct {
		ctx      context.Context
		credName string
		userID   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.CredentialsData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CredentialServiceEncryptedProxy{
				cypherBlock: tt.fields.cypherBlock,
				svc:         tt.fields.svc,
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

func TestCredentialServiceEncryptedProxy_SaveBankingCardData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.BankingCardData, userID int)
	type fields struct {
		cypherBlock cipher.Block
	}
	type args struct {
		ctx    context.Context
		data   *model.BankingCardData
		userID int
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should save bank card with encrypted data",
			fields: fields{
				cypherBlock: createCipherBlock("test"),
			},
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data *model.BankingCardData, userID int) {
				s.EXPECT().SaveBankingCardData(ctx, data, userID).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				data: &model.BankingCardData{
					Name:           "test card",
					Number:         "1231",
					ValidUntil:     "10/25",
					CardholderName: "holder",
					CVV:            "124",
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
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.data, tt.args.userID)
			s := &CredentialServiceEncryptedProxy{
				cypherBlock: tt.fields.cypherBlock,
				svc:         svc,
			}
			if err := s.SaveBankingCardData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveBankingCardData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialServiceEncryptedProxy_SaveBinaryData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, reader io.Reader, data *model.BinaryData, userID int)
	type fields struct {
		cypherBlock cipher.Block
	}
	type args struct {
		ctx    context.Context
		data   *model.BinaryData
		reader io.Reader
		userID int
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		//{
		//	name: "should save encrypted binary data",
		//	fields: fields{
		//		cypherBlock: createCipherBlock("test"),
		//	},
		//	mockBehavior: func(ctx context.Context,
		//		s *mock_service.MockDataService,
		//		reader io.Reader,
		//		data *model.BinaryData,
		//		userID int,
		//	) {
		//		s.EXPECT().SaveBinaryData(ctx, reader, data, userID)
		//	},
		//	args: args{
		//		ctx: context.Background(),
		//		data: &model.BinaryData{
		//			Data: []byte("test"),
		//			Name: "test binary",
		//		},
		//		userID: 1,
		//		reader: bytes.NewReader([]byte("test")),
		//	},
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.reader, tt.args.data, tt.args.userID)
			s := &CredentialServiceEncryptedProxy{
				cypherBlock: tt.fields.cypherBlock,
				svc:         svc,
			}
			if err := s.SaveBinaryData(tt.args.ctx, tt.args.reader, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveBinaryData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialServiceEncryptedProxy_SaveCredentialsData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.CredentialsData, userID int)
	type fields struct {
		cypherBlock cipher.Block
	}
	type args struct {
		ctx    context.Context
		data   *model.CredentialsData
		userID int
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should save encrypted credentials",
			fields: fields{
				cypherBlock: createCipherBlock("test"),
			},
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data *model.CredentialsData, userID int) {
				s.EXPECT().SaveCredentialsData(ctx, data, userID).Return(nil)
			},
			args: args{
				ctx: context.Background(),
				data: &model.CredentialsData{
					Name:     "test",
					Login:    "test_login",
					Password: "test_pwd",
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
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.data, tt.args.userID)
			s := &CredentialServiceEncryptedProxy{
				cypherBlock: tt.fields.cypherBlock,
				svc:         svc,
			}
			if err := s.SaveCredentialsData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveCredentialsData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCredentialServiceEncryptedProxy_SaveTextData(t *testing.T) {
	type mockBehavior func(ctx context.Context, s *mock_service.MockDataService, data *model.TextData, userID int)
	type fields struct {
		cypherBlock cipher.Block
	}
	type args struct {
		ctx    context.Context
		data   *model.TextData
		userID int
	}
	tests := []struct {
		name         string
		fields       fields
		mockBehavior mockBehavior
		args         args
		wantErr      bool
	}{
		{
			name: "should save encrypted text data",
			fields: fields{
				cypherBlock: createCipherBlock("test"),
			},
			mockBehavior: func(ctx context.Context, s *mock_service.MockDataService, data *model.TextData, userID int) {
				s.EXPECT().SaveTextData(ctx, data, userID)
			},
			args: args{
				ctx: context.Background(),
				data: &model.TextData{
					Name: "test",
					Data: "text",
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
			svc := mock_service.NewMockDataService(ctrl)
			tt.mockBehavior(tt.args.ctx, svc, tt.args.data, tt.args.userID)
			s := &CredentialServiceEncryptedProxy{
				cypherBlock: tt.fields.cypherBlock,
				svc:         svc,
			}
			if err := s.SaveTextData(tt.args.ctx, tt.args.data, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveTextData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
