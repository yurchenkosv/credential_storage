package configProvider

type ConfigProvider interface {
	GetConfig() ServerConfig
	Parse() error
}
