package configProvider

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type ClientConfig struct {
	ServerAddress string `env:"CRED_SERVER_ADDRESS"`
}

type ClientConfigProvider struct {
	cnf *ClientConfig
}

func NewClientConfigProvider() (*ClientConfigProvider, error) {
	provider := &ClientConfigProvider{}
	err := provider.Parse()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *ClientConfigProvider) Parse() error {
	flag.StringVar(&p.cnf.ServerAddress, "address", "localhost:8090", "server address, default localhost:8090")
	flag.Parse()
	err := env.Parse(&p.cnf)
	return err
}

func (p *ClientConfigProvider) GetConfig() ClientConfig {
	return *p.cnf
}
