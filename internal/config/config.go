package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ApiPort  string `env:"API_PORT" env-default:"8080"`
	DiagPort string `env:"DIAG_PORT" env-default:"8081"`
	DBUrl    string `env:"DB_URL" env-default:"postgres://user:password@localhost:5432/petstore?sslmode=disable"`
}

func MustLoad() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("none exists .env file")
	}

	return &cfg, nil
}
