package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"shop/src/core/settings"
	"shop/src/infra/db/core"
)

// Init инициализирует Redis клиента
func InitRedis() {
	config := settings.Config
	host := config.RedisHost
	port := config.RedisPort
	password := config.RedisPassword // может быть пустым

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

	core.InitRedis(rdb)
}

// Close закрывает Redis соединение
func CloseRedis() {
	rdb := core.GetRedis()
	if rdb == nil {
		return
	}
	if err := rdb.Close(); err != nil {
		log.Printf("failed to close Redis: %v", err)
	}
}
