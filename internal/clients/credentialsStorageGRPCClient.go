package clients

import (
	"context"
	"github.com/yurchenkosv/credential_storage/internal/api"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strconv"
)

type CredentialsStorageGRPCClient struct {
	opts              []grpc.CallOption
	credServiceClient api.CredentialServiceClient
	authClient        api.AuthServiceClient
}

func NewCredentialsStorageGRPCClient(connect *grpc.ClientConn, opts ...grpc.CallOption) *CredentialsStorageGRPCClient {
	grpcClient := api.NewCredentialServiceClient(connect)
	authClient := api.NewAuthServiceClient(connect)
	return &CredentialsStorageGRPCClient{
		credServiceClient: grpcClient,
		authClient:        authClient,
		opts:              opts,
	}
}

func (c *CredentialsStorageGRPCClient) AuthenticateUser(ctx context.Context, login string, password string) (string, error) {
	var header metadata.MD
	userAuth := &api.UserAuthentication{
		Login:    login,
		Password: password,
	}
	_, err := c.authClient.AuthenticateUser(ctx, userAuth, grpc.Header(&header))
	if err != nil {
		return "", err
	}
	jwtToken := header.Get("jwt")[0]
	ctx = metadata.NewOutgoingContext(ctx, header)
	return jwtToken, nil
}

func (c *CredentialsStorageGRPCClient) RegisterUser(ctx context.Context, user model.User) (string, error) {
	registration := &api.UserRegistration{
		Login:    user.Username,
		Password: user.Password,
		Name:     user.Name,
	}
	_, err := c.authClient.RegisterUser(ctx, registration, c.opts...)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (c *CredentialsStorageGRPCClient) GetData(ctx context.Context) ([]model.Credentials, error) {
	creds := []model.Credentials{}
	apiCreds, err := c.credServiceClient.GetData(ctx, &api.AllDataRequest{}, c.opts...)
	if err != nil {
		return []model.Credentials{}, err
	}
	for _, secret := range apiCreds.Secrets {
		credentials := model.Credentials{
			ID:         int(secret.Id),
			Name:       secret.Name,
			BinaryData: nil,
		}
		if secret.CredentialsData != nil {
			credentialsData := model.CredentialsData{
				ID:       int(secret.CredentialsData.Id),
				Name:     secret.CredentialsData.Name,
				Login:    secret.CredentialsData.Login,
				Password: secret.CredentialsData.Password,
				Metadata: nil,
			}
			credentials.CredentialsData = &credentialsData
		}
		if secret.BankingData != nil {
			bankingData := model.BankingCardData{
				ID:             int(secret.BankingData.Id),
				Name:           secret.BankingData.Name,
				Number:         strconv.Itoa(int(secret.BankingData.Number)),
				ValidUntil:     secret.BankingData.ValidTill,
				CardholderName: secret.BankingData.CardholderName,
				CVV:            strconv.Itoa(int(secret.BankingData.Cvv)),
				Metadata:       nil,
			}
			credentials.BankingCardData = &bankingData
		}
		if secret.BinaryData != nil {
			binaryData := model.BinaryData{
				ID:       int(secret.BinaryData.Id),
				Data:     secret.BinaryData.Data,
				Metadata: nil,
			}
			credentials.BinaryData = &binaryData
		}
		if secret.TextData != nil {
			textData := model.TextData{
				ID:       int(secret.TextData.Id),
				Name:     secret.TextData.Name,
				Data:     secret.TextData.Data,
				Metadata: nil,
			}
			credentials.TextData = &textData
		}
		for _, meta := range secret.Metadata {
			data := model.Metadata{Value: meta}
			credentials.Metadata = append(credentials.Metadata, data)
		}
		creds = append(creds, credentials)
	}
	return creds, nil
}

func (c *CredentialsStorageGRPCClient) SendCredentials(ctx context.Context, data model.CredentialsData) error {
	var meta []string
	for _, m := range data.Metadata {
		meta = append(meta, m.Value)
	}
	apiData := &api.CredentialsData{
		Login:    data.Login,
		Password: data.Password,
		Name:     data.Name,
		Metadata: meta,
	}
	_, err := c.credServiceClient.SaveCredentialsData(ctx, apiData, c.opts...)
	if err != nil {
		return err
	}
	return nil
}

func (c *CredentialsStorageGRPCClient) SendBankCard(ctx context.Context, data model.BankingCardData) error {
	var meta []string
	for _, m := range data.Metadata {
		meta = append(meta, m.Value)
	}
	num, _ := strconv.ParseInt(data.Number, 10, 64)
	cvv, _ := strconv.ParseInt(data.CVV, 10, 64)
	apiData := &api.BankingCardData{
		Number:         int32(num),
		ValidTill:      data.ValidUntil,
		CardholderName: data.CardholderName,
		Cvv:            int32(cvv),
		Name:           data.Name,
		Metadata:       meta,
	}
	_, err := c.credServiceClient.SaveBankingData(ctx, apiData, c.opts...)
	if err != nil {
		return err
	}
	return nil
}

func (c *CredentialsStorageGRPCClient) SendBinary(ctx context.Context, data model.BinaryData) error {
	var meta []string
	for _, m := range data.Metadata {
		meta = append(meta, m.Value)
	}
	apiData := &api.BinaryData{
		Data:     data.Data,
		Name:     data.Name,
		Metadata: meta,
	}
	_, err := c.credServiceClient.SaveBinaryData(ctx, apiData, c.opts...)
	return err
}

func (c *CredentialsStorageGRPCClient) SendText(ctx context.Context, data model.TextData) error {
	var meta []string
	for _, m := range data.Metadata {
		meta = append(meta, m.Value)
	}
	apiData := &api.TextData{
		Data:     data.Data,
		Name:     data.Name,
		Metadata: meta,
	}
	_, err := c.credServiceClient.SaveTextData(ctx, apiData, c.opts...)
	if err != nil {
		return err
	}
	return nil
}

func (c *CredentialsStorageGRPCClient) DeleteData(ctx context.Context, data model.Credentials) error {
	credData := api.SecretsData{}
	if data.BankingCardData != nil {
		credData.BankingData = &api.BankingCardData{Id: int32(data.ID)}
	}
	if data.BinaryData != nil {
		credData.BinaryData = &api.BinaryData{Id: int32(data.BinaryData.ID)}
	}
	if data.TextData != nil {
		credData.TextData = &api.TextData{Id: int32(data.BinaryData.ID)}
	}
	if data.CredentialsData != nil {
		credData.CredentialsData = &api.CredentialsData{Id: int32(data.CredentialsData.ID)}
	}
	for _, meta := range data.Metadata {
		credData.Metadata = append(credData.Metadata, meta.Value)
	}
	_, err := c.credServiceClient.DeleteData(ctx, &credData, c.opts...)
	if err != nil {
		return err
	}
	return nil
}
