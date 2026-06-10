package config

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	PostgresHost       string `env:"POSTGRES_HOST,required"`
	PostgresPort       uint   `env:"POSTGRES_PORT,required"`
	PostgresUser       string `env:"POSTGRES_USER,required"`
	PostgresPassword   string `env:"POSTGRES_PASSWORD,required"`
	PostgresDatabase   string `env:"POSTGRES_DATABASE,required"`
	ProductServicePort uint   `env:"PRODUCT_SERVICE_PORT,required"`
	OrderServicePort   uint   `env:"ORDER_SERVICE_PORT,required"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
