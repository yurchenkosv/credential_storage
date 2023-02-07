package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"reflect"
	"testing"
)

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

func createCipherBlock(key string) cipher.Block {
	initKey := sha256.Sum256([]byte(key))
	block, _ := aes.NewCipher(initKey[:])
	return block
}
