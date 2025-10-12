package tokenstorage

import (
	"context"
	"fmt"
	"github.com/Fox216540/shop/auth-service/domain/tokenstorage"
	"github.com/Fox216540/shop/auth-service/infra/redis/core"
	"github.com/google/uuid"
	"time"
)

type Repository struct {
	ttl time.Duration
	rdb *core.Redis
}

func NewRepository(rdb *core.Redis, ttl time.Duration) tokenstorage.Repository {
	return &Repository{ttl: ttl, rdb: rdb}
}

func (r *Repository) Set(jti, userID uuid.UUID) error {
	ctx := context.Background()
	userSetKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	if err := r.rdb.AddToSet(ctx, userSetKey, jti.String()); err != nil {
		return NewInvalidSet(err)
	}
	if err := r.rdb.Set(ctx, jti.String(), userID.String(), r.ttl); err != nil {
		return NewInvalidSet(err)
	}
	return nil
}

func (r *Repository) Exists(jti uuid.UUID) error {
	ctx := context.Background()
	exists, err := r.rdb.Exist(ctx, jti.String())
	if err != nil {
		return NewInvalidExists(err)
	}
	if !exists {
		return nil
	}
	return tokenstorage.NewNotFoundTokenOfUserError(nil)
}

func (r *Repository) Delete(jti, userID uuid.UUID) error {
	ctx := context.Background()
	userSetKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	if err := r.rdb.DeleteFromSet(ctx, userSetKey, jti.String()); err != nil {
		return NewInvalidDelete(err)
	}
	if err := r.rdb.DeleteKeys(ctx, jti.String()); err != nil {
		return NewInvalidDelete(err)
	}
	return nil
}

// TODO: Использовать в микросервисе user-service при удалении юзера
func (r *Repository) DeleteAll(userID uuid.UUID) error {
	ctx := context.Background()
	setKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	jtis, err := r.rdb.SMembers(ctx, setKey)
	if err != nil {
		return NewInvalidDeleteAll(err)
	}
	if len(jtis) == 0 {
		return tokenstorage.NewNotFoundTokensOfUserError(nil)
	}

	keysToDelete := make([]string, 0, len(jtis)+1)
	for _, jti := range jtis {
		keysToDelete = append(keysToDelete, jti)
	}
	keysToDelete = append(keysToDelete, setKey)
	if err := r.rdb.DeleteKeys(ctx, keysToDelete...); err != nil {
		return NewInvalidDeleteAll(err)
	}
	return nil

}
