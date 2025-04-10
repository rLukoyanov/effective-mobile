package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	DB_PATH   string `env:"DB_PATH" envDefault:"./data/persons.db"`
	PORT      string `env:"PORT" envDefault:"8080"`
	LOG_LEVEL string `env:"LOG_LEVEL" envDefault:"debug"`
	MODE      string `env:"MODE" envDefault:"dev"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
