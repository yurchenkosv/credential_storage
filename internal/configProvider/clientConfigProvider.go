package configProvider

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type ClientConfig struct {
	ServerAddress string `env:"CRED_SERVER_ADDRESS"`
	Login         string
	Password      string
	Name          string
	RegisterUser  bool
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
	flag.BoolVar(&p.cnf.RegisterUser, "r", false, "flag indicated that user need to be registered")
	flag.StringVar(&p.cnf.Name, "n", "", "name of the user to be registered with")
	flag.Parse()
	err := env.Parse(p.cnf)

	if p.cnf.RegisterUser {
		if p.cnf.Name == "" {
			return errors.New("Name of the user must be specified for registering (-n flag)")
		}
	}
	return err
}

func (p *ClientConfigProvider) GetConfig() ClientConfig {
	return *p.cnf
}
