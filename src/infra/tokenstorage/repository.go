package tokenstorage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"shop/src/domain/tokenstorage"
)

type repository struct {
	rdb *redis.Client
}

func NewRepository(rdb *redis.Client) tokenstorage.TokenStorage {
	return &repository{rdb: rdb}
}

func (r *repository) Set(jti, userID uuid.UUID) error {
	ctx := context.Background()
	userSetKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	if err := r.rdb.SAdd(ctx, userSetKey, jti.String()).Err(); err != nil {
		return err
	}
	if err := r.rdb.Set(ctx, jti.String(), userID.String(), 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *repository) Exists(jti uuid.UUID) (bool, error) {
	ctx := context.Background()
	exists, err := r.rdb.Exists(ctx, jti.String()).Result()
	if err != nil {
		return false, err
	}
	if exists > 0 {
		return true, nil
	}

	return false, error(nil)
}

func (r *repository) Delete(jti uuid.UUID) error {
	ctx := context.Background()
	err := r.rdb.Del(ctx, jti.String()).Err()
	if err != nil {
		return err
	}
	return error(nil)
}

func (r *repository) DeleteAll(userID uuid.UUID) error {
	ctx := context.Background()
	err := r.rdb.Del(ctx, userID.String()).Err()
	if err != nil {
		return err
	}

	return error(nil)
}
