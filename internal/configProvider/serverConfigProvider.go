package configProvider

import (
	"errors"
	"github.com/caarlos0/env/v6"
)

type ServerConfig struct {
	DatabaseDSN      string `env:"CRED_SERVER_DATABASE_DSN" yaml:"database_dsn"`
	Listen           string `env:"CRED_SERVER_LISTEN" envDefault:"localhost:8080" yaml:"listen"`
	JWTSecret        string `env:"CRED_SERVER_JWT_SECRET" yaml:"jwt_secret"`
	EncryptionSecret string `env:"CRED_SERVER_ENCRYPTION_SECRET" yaml:"encryption_secret"`
	ListenGRPC       string `env:"CRED_SERVER_GRPC_LISTEN" envDefault:"localhost:8090" yaml:"listen_grpc"`
}

type ServerConfigProvider struct {
	cnf *ServerConfig
}

func NewServerConfigProvider() (*ServerConfigProvider, error) {
	provider := &ServerConfigProvider{}
	err := provider.Parse()
	return provider, err
}

func (p *ServerConfigProvider) Parse() error {
	cnf := &ServerConfig{}
	err := env.Parse(cnf)
	if err != nil {
		return err
	}
	if cnf.JWTSecret == "" {
		return errors.New("JWT secret must be set")
	}
	if cnf.EncryptionSecret == "" {
		return errors.New("encryption secret must be set")
	}
	p.cnf = cnf
	return nil
}

func (p *ServerConfigProvider) GetConfig() ServerConfig {
	return *p.cnf
}
