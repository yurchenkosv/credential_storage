package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"github.com/yurchenkosv/credential_storage/internal/binaryRepository"
	"github.com/yurchenkosv/credential_storage/internal/model"
	"github.com/yurchenkosv/credential_storage/internal/repository"
	"io"
)

type CredentialServiceEncryptedProxy struct {
	cypherBlock cipher.Block
	svc         *CredentialsService
}

func NewProxyEncryptedCredentialService(
	repo repository.Repository,
	binaryRepo binaryRepository.BinaryRepository,
	encryptionKey string,
) (*CredentialServiceEncryptedProxy, error) {
	block, err := initCypher(encryptionKey)
	if err != nil {
		return nil, err
	}

	return &CredentialServiceEncryptedProxy{
		svc:         NewCredentialsService(repo, binaryRepo),
		cypherBlock: block,
	}, nil
}

func (s *CredentialServiceEncryptedProxy) SaveCredentialsData(ctx context.Context,
	data *model.CredentialsData,
	userID int) error {
	encryptedData := data
	encryptedData.Name = string(encryptData(s.cypherBlock, []byte(data.Name)))
	encryptedData.Login = string(encryptData(s.cypherBlock, []byte(data.Login)))
	encryptedData.Password = string(encryptData(s.cypherBlock, []byte(data.Password)))
	encryptedData.Metadata = encryptMetadata(encryptedData.Metadata, s.cypherBlock)
	return s.svc.SaveCredentialsData(ctx, encryptedData, userID)
}

func (s *CredentialServiceEncryptedProxy) SaveBankingCardData(ctx context.Context,
	data *model.BankingCardData,
	userID int) error {
	data.Name = string(encryptData(s.cypherBlock, []byte(data.Name)))
	data.CardholderName = string(encryptData(s.cypherBlock, []byte(data.CardholderName)))
	data.Number = string(encryptData(s.cypherBlock, []byte(data.Number)))
	data.ValidUntil = string(encryptData(s.cypherBlock, []byte(data.ValidUntil)))
	data.CVV = string(encryptData(s.cypherBlock, []byte(data.CVV)))
	data.Metadata = encryptMetadata(data.Metadata, s.cypherBlock)
	return s.svc.SaveBankingCardData(ctx, data, userID)
}

func (s *CredentialServiceEncryptedProxy) SaveTextData(ctx context.Context,
	data *model.TextData,
	userID int) error {
	data.Name = string(encryptData(s.cypherBlock, []byte(data.Name)))
	data.Data = string(encryptData(s.cypherBlock, []byte(data.Data)))
	data.Metadata = encryptMetadata(data.Metadata, s.cypherBlock)
	return s.svc.SaveTextData(ctx, data, userID)
}

func (s *CredentialServiceEncryptedProxy) SaveBinaryData(ctx context.Context,
	data *model.BinaryData,
	userID int) error {
	data.Data = encryptData(s.cypherBlock, data.Data)
	data.Name = string(encryptData(s.cypherBlock, []byte(data.Name)))
	data.Metadata = encryptMetadata(data.Metadata, s.cypherBlock)
	return s.svc.SaveBinaryData(ctx, data, userID)
}

func (s *CredentialServiceEncryptedProxy) GetCredentialsByName(ctx context.Context,
	credName string,
	userID int) ([]model.CredentialsData, error) {
	return s.svc.GetCredentialsByName(ctx, credName, userID)
}

func (s *CredentialServiceEncryptedProxy) GetAllUserCredentials(ctx context.Context,
	userID int) ([]model.Credentials, error) {
	data, err := s.svc.GetAllUserCredentials(ctx, userID)
	if err != nil {
		return nil, err
	}
	for idx, cred := range data {
		if cred.CredentialsData != nil {
			cred.CredentialsData.Login = string(decryptData([]byte(cred.CredentialsData.Login), s.cypherBlock))
			cred.CredentialsData.Password = string(decryptData([]byte(cred.CredentialsData.Password), s.cypherBlock))
		}
		if cred.TextData != nil {
			cred.TextData.Data = string(decryptData([]byte(cred.TextData.Data), s.cypherBlock))
		}
		if cred.BankingCardData != nil {
			cred.BankingCardData.CardholderName = string(decryptData([]byte(cred.BankingCardData.CardholderName), s.cypherBlock))
			cred.BankingCardData.Number = string(decryptData([]byte(cred.BankingCardData.Number), s.cypherBlock))
			cred.BankingCardData.ValidUntil = string(decryptData([]byte(cred.BankingCardData.CVV), s.cypherBlock))
			cred.BankingCardData.CVV = string(decryptData([]byte(cred.BankingCardData.CVV), s.cypherBlock))
		}
		if cred.BinaryData != nil {
			cred.BinaryData.Data = decryptData(cred.BinaryData.Data, s.cypherBlock)
			cred.BinaryData.Link = string(decryptData([]byte(cred.BinaryData.Link), s.cypherBlock))
		}
		cred.Name = string(decryptData([]byte(cred.Name), s.cypherBlock))
		cred.Metadata = decryptMetadata(cred.Metadata, s.cypherBlock)
		data[idx] = cred
	}
	return data, nil
}

func initCypher(key string) (cipher.Block, error) {
	initKey := sha256.Sum256([]byte(key))
	return aes.NewCipher(initKey[:])
}

func encryptData(cyperBlock cipher.Block, data []byte) []byte {
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(cyperBlock, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	encrypted := base64.URLEncoding.EncodeToString(ciphertext)
	return []byte(encrypted)
}

func decryptData(data []byte, block cipher.Block) []byte {
	ciphertext, _ := base64.URLEncoding.DecodeString(string(data))
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

func encryptMetadata(metadata []model.Metadata, block cipher.Block) []model.Metadata {
	for idx := range metadata {
		metadata[idx].Value = string(encryptData(block, []byte(metadata[idx].Value)))
	}
	return metadata
}
func decryptMetadata(metadata []model.Metadata, block cipher.Block) []model.Metadata {
	for idx := range metadata {
		metadata[idx].Value = string(decryptData([]byte(metadata[idx].Value), block))
	}
	return metadata
}
