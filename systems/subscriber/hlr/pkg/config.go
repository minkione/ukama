package pkg

import (
	"time"

	"github.com/ukama/ukama/systems/common/config"
)

type Config struct {
	config.BaseConfig `mapstructure:",squash"`
	DB                config.Database
	Grpc              config.Grpc
	SimTokenKey       string
	HlrHost           string
	NetworkHost       string
	PCRFHost          string
	FactoryHost       string
	Org               string
	Timeout           time.Duration     `default:"3s"`
	MsgClient         *config.MsgClient `default:"{}"`
}

type SimManager struct {
	Host string
	Name string
}

func NewConfig() *Config {
	return &Config{
		DB: config.Database{
			DbName: ServiceName,
		},

		Grpc: config.Grpc{
			Port: 9090,
		},
		SimTokenKey: "11111111111111111111111111111111",
		HlrHost:     "localhost:8080",
		NetworkHost: "http://localhost:8085",
		PCRFHost:    "http://localhost:8085",
		FactoryHost: "http://localhost:8085",
		Org:         "880f7c63-eb57-461a-b514-248ce91e9b3e",
		MsgClient: &config.MsgClient{
			Timeout:        5 * time.Second,
			ListenerRoutes: []string{"event.cloud.lookup.organization.create"},
		},
	}
}
