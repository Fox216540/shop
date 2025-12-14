package tokenstorage

import (
	"github.com/Fox216540/shop/auth-service/app/mapError"
	"github.com/Fox216540/shop/auth-service/domain/tokenstorage"
	"github.com/google/uuid"
)

type service struct {
	ts tokenstorage.Repository
}

func NewService(ts tokenstorage.Repository) UseCase {
	return &service{ts: ts}
}

func (s *service) Add(jti, userID uuid.UUID) error {
	if err := s.ts.Set(jti, userID); err != nil {
		return mapError.MapError(err, NewInvalidAdd(err))
	}
	return nil
}

func (s *service) Delete(jti, userID uuid.UUID) error {
	if err := s.ts.Delete(jti, userID); err != nil {
		return mapError.MapError(err, NewInvalidDelete(err))
	}
	return nil
}

func (s *service) DeleteAll(userID uuid.UUID) error {
	if err := s.ts.DeleteAll(userID); err != nil {
		return mapError.MapError(err, NewInvalidDeleteAll(err))
	}
	return nil
}
