package configProvider

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type ClientConfig struct {
	ServerAddress string `env:"CRED_SERVER_ADDRESS"`
	Login         string
	Password      string
}

type ClientConfigProvider struct {
	cnf *ClientConfig
}

func NewClientConfigProvider() (*ClientConfigProvider, error) {
	provider := &ClientConfigProvider{cnf: &ClientConfig{}}
	err := provider.Parse()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *ClientConfigProvider) Parse() error {
	flag.StringVar(&p.cnf.ServerAddress,
		"address",
		"localhost:8090",
		"server address, default localhost:8090",
	)
	flag.StringVar(&p.cnf.Login,
		"u",
		"",
		"username to authenticate to server",
	)
	flag.StringVar(&p.cnf.Password,
		"p",
		"",
		"password to authenticate to server",
	)
	flag.Parse()
	err := env.Parse(p.cnf)
	return err
}

func (p *ClientConfigProvider) GetConfig() ClientConfig {
	return *p.cnf
}
