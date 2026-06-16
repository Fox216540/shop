package config

import "github.com/caarlos0/env/v10"

type Config struct {
	PostgresHost        string `env:"POSTGRES_HOST,required"`
	PostgresPort        uint   `env:"POSTGRES_PORT,required"`
	PostgresUser        string `env:"POSTGRES_USER,required"`
	PostgresPassword    string `env:"POSTGRES_PASSWORD,required"`
	PostgresDatabase    string `env:"POSTGRES_DATABASE,required"`
	PaymentServicePort  uint   `env:"PAYMENT_SERVICE_PORT,required"`
	PaymentProviderMode string `env:"PAYMENT_PROVIDER_MODE" envDefault:"mock"`
	YoKassaAccountID    string `env:"YOOKASSA_ACCOUNT_ID"`
	YoKassaSecretKey    string `env:"YOOKASSA_SECRET_KEY"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
