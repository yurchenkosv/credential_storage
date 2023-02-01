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
func (c *CredentialsStorageGRPCClient) GetData() (model.Credentials, error) {
	creds := model.Credentials{}
	return creds, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) SendCredentials(credentials model.CredentialsData) error {
	return errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) SendBankCard(card model.BankingCardData) error {
	return errors.New("not implemented")

}

func (c *CredentialsStorageGRPCClient) SendBinary(binary model.BinaryData) error {
	return errors.New("not implemented")

}

func (c *CredentialsStorageGRPCClient) SendText(text model.TextData) error {
	return errors.New("not implemented")
}
