package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type (
	DB struct {
		DBUrl         string        `envconfig:"DB_URL"`
		DBDriver      string        `envconfig:"DB_DRIVER"`
		DbMaxConn     int           `envconfig:"DB_MAX_CONN"`
		DbMaxIdle     int           `envconfig:"DB_MAX_IDLE"`
		DbMaxConnTime time.Duration `envconfig:"DB_MAX_CONN_TIME"`
		DbMaxIdleTime time.Duration `envconfig:"DB_MAX_IDLE_TIME"`
	}

	App struct {
		Name    string `envconfig:"APP_NAME"`
		Version string `envconfig:"APP_VERSION"`
	}

	Server struct {
		GRPCPort        string        `envconfig:"GRPC_PORT"`
		HealthPort      string        `envconfig:"HEALTH_PORT"`
		HTTPTimeout     time.Duration `envconfig:"HTTP_TIMEOUT"`
		IdleTimeout     time.Duration `envconfig:"IDLE_TIMEOUT"`
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT"`
	}

	Config struct {
		App    App
		Server Server
		DB     DB
	}
)

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	return &cfg, nil
}
