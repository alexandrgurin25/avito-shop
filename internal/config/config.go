package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     uint16 `env:"POSTGRES_PORT"`
	Username string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`

	MinConns int32 `env:"POSTGRES_MIN_CONN"`
	MaxConns int32 `env:"POSTGRES_MAX_CONN"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
