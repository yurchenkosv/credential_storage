package clients

import (
	"errors"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"google.golang.org/grpc"
)

type CredentialsStorageGRPCClient struct {
	connect *grpc.ClientConn
	opts    []grpc.CallOption
}

func NewCredentialsStorageGRPCClient(connect *grpc.ClientConn, opts ...grpc.CallOption) *CredentialsStorageGRPCClient {
	return &CredentialsStorageGRPCClient{
		connect: connect,
		opts:    opts,
	}
}

func (c *CredentialsStorageGRPCClient) GetCredentials() (model.Credentials, error) {
	//client := api.NewCredentialServiceClient(c.connect)
	return model.Credentials{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) GetBankCard() (model.BankingCard, error) {
	return model.BankingCard{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) GetBinary() (model.BinaryData, error) {
	return model.BinaryData{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) GetText() (model.TextData, error) {
	return model.TextData{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) SendCredentials(credentials model.Credentials) error {
	return errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) SendBankCard(card model.BankingCard) error {
	return errors.New("not implemented")

}

func (c *CredentialsStorageGRPCClient) SendBinary(binary model.BinaryData) error {
	return errors.New("not implemented")

}

func (c *CredentialsStorageGRPCClient) SendText(text model.TextData) error {

	return errors.New("not implemented")
}
