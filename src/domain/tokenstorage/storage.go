package tokenstorage

import (
	"github.com/google/uuid"
)

type TokenStorage interface {
	Set(jti, userID uuid.UUID) error
	Exists(jti uuid.UUID) (bool, error)
	Delete(jti uuid.UUID) error
	DeleteAll(userID uuid.UUID) error
}
