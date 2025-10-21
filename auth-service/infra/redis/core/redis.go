package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{client: client}
}

func (r *Redis) Client() *redis.Client {
	return r.client
}

func (r *Redis) AddToSet(ctx context.Context, key string, member string) error {
	return r.client.SAdd(ctx, key, member).Err()
}

func (r *Redis) DeleteKeys(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Del(ctx, keys...).Result()
}

func (r *Redis) DeleteFromSet(ctx context.Context, key string, member string) (int64, error) {
	return r.client.SRem(ctx, key, member).Result()
}

func (r *Redis) SMembers(ctx context.Context, key string) ([]string, error) {
	return r.client.SMembers(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
