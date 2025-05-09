package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	POSTGRES_DSN string `env:"POSTGRES_DSN" envDefault:"postgresql://root:123@localhost:5432/base?sslmode=disable"`
	PORT         string `env:"PORT" envDefault:"8080"`
	LOG_LEVEL    string `env:"LOG_LEVEL" envDefault:"debug"`
	MODE         string `env:"MODE" envDefault:"dev"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
