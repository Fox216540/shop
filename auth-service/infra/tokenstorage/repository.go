package tokenstorage

import (
	"context"
	"fmt"
	"github.com/Fox216540/shop/auth-service/domain/tokenstorage"
	"github.com/Fox216540/shop/auth-service/infra/redis/core"
	"github.com/google/uuid"
	pkgerrors "github.com/pkg/errors"
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
		return NewInvalidSet(pkgerrors.WithStack(err))
	}
	if err := r.rdb.Set(ctx, jti.String(), userID.String(), r.ttl); err != nil {
		return NewInvalidSet(pkgerrors.WithStack(err))
	}
	return nil
}

func (r *Repository) Delete(jti, userID uuid.UUID) error {
	ctx := context.Background()
	userSetKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	removed, err := r.rdb.DeleteFromSet(ctx, userSetKey, jti.String())
	if err != nil {
		return NewInvalidDelete(pkgerrors.WithStack(err))
	}
	if removed == 0 {
		return tokenstorage.NewNotFoundTokenOfUserError(nil)
	}
	removed, err = r.rdb.DeleteKeys(ctx, jti.String())
	if err != nil {
		return NewInvalidDelete(pkgerrors.WithStack(err))
	}
	if removed == 0 {
		return tokenstorage.NewNotFoundTokenOfUserError(nil)
	}
	return nil
}

func (r *Repository) DeleteAll(userID uuid.UUID) error {
	ctx := context.Background()
	setKey := fmt.Sprintf("user:%s:refresh_tokens", userID.String())
	jtis, err := r.rdb.SMembers(ctx, setKey)
	if err != nil {
		return NewInvalidDeleteAll(pkgerrors.WithStack(err))
	}
	if len(jtis) == 0 {
		return tokenstorage.NewNotFoundTokensOfUserError(nil)
	}

	keysToDelete := make([]string, 0, len(jtis)+1)
	for _, jti := range jtis {
		keysToDelete = append(keysToDelete, jti)
	}
	keysToDelete = append(keysToDelete, setKey)
	removed, err := r.rdb.DeleteKeys(ctx, keysToDelete...)
	if err != nil {
		return NewInvalidDeleteAll(pkgerrors.WithStack(err))
	}
	if removed == 0 {
		return tokenstorage.NewNotFoundTokensOfUserError(nil)
	}
	return nil
}
