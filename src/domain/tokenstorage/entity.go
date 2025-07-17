package tokenstorage

import (
	"github.com/google/uuid"
)

type TokenStorage interface {
	Set(token string, userID uuid.UUID) error
	Exists(token string) (bool, error)
	Delete(token string) error
	DeleteAll(userID uuid.UUID) error
}
