package pkg

import (
	"github.com/ukama/ukama/systems/common/config"
)

type Config struct {
	config.BaseConfig `mapstructure:",squash"`
	DB                config.Database
	Grpc              config.Grpc
	SimTokenKey       string
	HlrHost           string
	NetworkHost       string
	PCRF              string
	Factory           string
}

type SimManager struct {
	Host string
	Name string
}

func NewConfig() *Config {
	return &Config{
		DB: config.Database{
			Host:       "localhost",
			Password:   "Pass2020!",
			DbName:     ServiceName,
			Username:   "postgres",
			Port:       5432,
			SslEnabled: false,
		},

		Grpc: config.Grpc{
			Port: 9090,
		},
		SimTokenKey: "11111111111111111111111111111111",
		HlrHost:     "localhost:9090",
	}
}
