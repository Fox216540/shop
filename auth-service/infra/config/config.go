package config

import (
	"github.com/caarlos0/env/v10"
	"time"
)

type Config struct {
	RedisHost          string        `env:"REDIS_HOST,required"`
	RedisPort          uint          `env:"REDIS_PORT,required"`
	RedisPassword      string        `env:"REDIS_PASSWORD,required"`
	AccessTokenSecret  string        `env:"ACCESS_TOKEN_SECRET,required"`
	AccessTokenTTL     time.Duration `env:"ACCESS_TOKEN_TTL,required"`
	RefreshTokenSecret string        `env:"REFRESH_TOKEN_SECRET,required"`
	RefreshTokenTTL    time.Duration `env:"REFRESH_TOKEN_TTL,required"`
	AuthServicePort    uint          `env:"AUTH_SERVICE_PORT,required"`
	UserServicePort    uint          `env:"USER_SERVICE_PORT,required"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
