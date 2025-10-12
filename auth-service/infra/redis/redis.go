package redis

import (
	"context"
	"fmt"
	"github.com/Fox216540/shop/auth-service/infra/config"
	"github.com/redis/go-redis/v9"
	"log"
)

// Init инициализирует Redis клиента
func InitRedis(conf *config.Config) *redis.Client {
	host := conf.RedisHost
	port := conf.RedisPort
	password := conf.RedisPassword // может быть пустым

	addr := fmt.Sprintf("%s:%s", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Проверка соединения
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	return rdb
}

// Close закрывает Redis соединение
func CloseRedis(rdb *redis.Client) {
	if rdb == nil {
		return
	}
	if err := rdb.Close(); err != nil {
		log.Printf("failed to close Redis: %v", err)
	}
}
