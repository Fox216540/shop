package config

import "github.com/caarlos0/env/v10"

type Config struct {
	APIGatewayPort     uint `env:"APIGATEWAY_PORT" envDefault:"8080"`
	AuthServicePort    uint `env:"AUTH_SERVICE_PORT,required"`
	BasketServicePort  uint `env:"BASKET_SERVICE_PORT,required"`
	CatalogServicePort uint `env:"CATALOG_SERVICE_PORT,required"`
	OrderServicePort   uint `env:"ORDER_SERVICE_PORT,required"`
	UserServicePort    uint `env:"USER_SERVICE_PORT,required"`
	RefreshTokenTTL    uint `env:"REFRESH_TOKEN_TTL,required"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
