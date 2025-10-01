package tokenstorage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"shop/src/domain/tokenstorage"
	"time"
)

type repository struct {
	rdb *redis.Client
}

func NewRepository(rdb *redis.Client) tokenstorage.TokenStorage {
	return &repository{rdb: rdb}
}

func (r *repository) Set(jti, userID uuid.UUID, ttl time.Duration) error {
	ctx := context.Background()
	userSetKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	if err := r.rdb.SAdd(ctx, userSetKey, jti.String()).Err(); err != nil {
		return NewInvalidSet(err)
	}
	if err := r.rdb.Set(ctx, jti.String(), userID.String(), ttl).Err(); err != nil {
		return NewInvalidSet(err)
	}
	return nil
}

func (r *repository) Exists(jti uuid.UUID) error {
	ctx := context.Background()
	exists, err := r.rdb.Exists(ctx, jti.String()).Result()
	if err != nil {
		return NewInvalidExists(err)
	}
	if exists > 0 {
		return nil
	}

	return tokenstorage.NewNotFoundTokenOfUserError(nil)
}

func (r *repository) Delete(jti, userID uuid.UUID) error {
	ctx := context.Background()
	userSetKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	if err := r.rdb.SRem(ctx, userSetKey, jti.String()).Err(); err != nil {
		return NewInvalidDelete(err)
	}
	err := r.rdb.Del(ctx, jti.String()).Err()
	if err != nil {
		return NewInvalidDelete(err)
	}
	return error(nil)
}

func (r *repository) DeleteAll(userID uuid.UUID) error {
	ctx := context.Background()
	setKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	jtis, err := r.rdb.SMembers(ctx, setKey).Result()
	if err != nil {
		return NewInvalidDeleteAll(err)
	}
	if len(jtis) == 0 {
		return tokenstorage.NewNotFoundTokensOfUserError(nil)
	}

	// Формируем список ключей для удаления
	keysToDelete := make([]string, 0, len(jtis)+1)
	for _, jti := range jtis {
		keysToDelete = append(keysToDelete, jti)
	}
	keysToDelete = append(keysToDelete, setKey) // Добавляем сам set
	// Удаляем всё одной командой
	if err := r.rdb.Del(ctx, keysToDelete...).Err(); err != nil {
		return NewInvalidDeleteAll(err)
	}
	return nil

}
