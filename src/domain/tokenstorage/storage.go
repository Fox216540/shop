package tokenstorage

import (
	"github.com/google/uuid"
	"time"
)

type TokenStorage interface {
	Set(jti, userID uuid.UUID, ttl time.Duration) error
	Exists(jti uuid.UUID) (bool, error)
	Delete(jti, userID uuid.UUID) error
	DeleteAll(userID uuid.UUID) error
}
