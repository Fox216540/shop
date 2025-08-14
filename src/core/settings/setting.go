package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Setting struct {
	RedisHost          string
	RedisPort          string
	RedisPassword      string
	PostgresHost       string
	PostgresPort       string
	PostgresUser       string
	PostgresPassword   string
	PostgresDatabase   string
	AccessTokenSecret  string
	AccessTokenTTL     string
	RefreshTokenSecret string
	RefreshTokenTTL    string
	BufferSeconds      string
}

func NewSetting() *Setting {
	err := godotenv.Load() // по умолчанию ищет файл .env в текущей папке
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return &Setting{
		RedisHost:          os.Getenv("REDIS_HOST"),
		RedisPort:          os.Getenv("REDIS_PORT"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		PostgresHost:       os.Getenv("POSTGRES_HOST"),
		PostgresPort:       os.Getenv("POSTGRES_PORT"),
		PostgresUser:       os.Getenv("POSTGRES_USER"),
		PostgresPassword:   os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase:   os.Getenv("POSTGRES_DB"),
		AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		AccessTokenTTL:     os.Getenv("ACCESS_TOKEN_TTL"),
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		RefreshTokenTTL:    os.Getenv("REFRESH_TOKEN_TTL"),
		BufferSeconds:      os.Getenv("BUFFER_SECONDS"),
	}
}

var Config = NewSetting()
