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

func (c *CredentialsStorageGRPCClient) GetCredentials() (model.CredentialsData, error) {
	//client := api.NewCredentialServiceClient(c.connect)
	return model.CredentialsData{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) GetBankCard() (model.BankingCardData, error) {
	return model.BankingCardData{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) GetBinary() (model.BinaryData, error) {
	return model.BinaryData{}, errors.New("not implemented")
}

func (c *CredentialsStorageGRPCClient) GetText() (model.TextData, error) {
	return model.TextData{}, errors.New("not implemented")
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
