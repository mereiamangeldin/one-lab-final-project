package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     int    `env:"PORT" envDefault:"8000"`
	Database string `env:"DATABASE" envDefault:"kazakTeam"`
	PgURL    string `env:"PG_URL" envDefault:"user=user password=secret dbname=db sslmode=disable host=localhost port=5432"`
}

func New() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
