package core

import (
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(client *redis.Client) {
	if client == nil {
		panic("InitRedis: client is nil")
	}
	redisClient = client
}

func GetRedis() *redis.Client {
	return redisClient
}
