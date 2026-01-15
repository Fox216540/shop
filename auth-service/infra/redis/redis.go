package redis

import (
	"context"
	"fmt"
	"github.com/Fox216540/shop/auth-service/core/logger"
	"github.com/Fox216540/shop/auth-service/infra/config"
	"github.com/Fox216540/shop/auth-service/infra/redis/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

// InitRedis инициализирует Redis клиента
func InitRedis(
	ctx context.Context,
	log logger.Logger,
	conf *config.Config,
) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.RedisPassword, // может быть пустым
		DB:       0,
	})

	ctxRedis, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctxRedis).Err(); err != nil {
		log.Error(ctx, errors.NewRedisError(err))
		return nil, errors.NewRedisError(err)
	}

	return rdb, nil
}

// CloseRedis закрывает Redis соединение
func CloseRedis(
	ctx context.Context,
	log logger.Logger,
	rdb *redis.Client,
) {
	if rdb == nil {
		return
	}

	if err := rdb.Close(); err != nil {
		log.Error(ctx, errors.NewCloseError(err))
	}
}
