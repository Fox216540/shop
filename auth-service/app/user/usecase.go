package user

import (
	"github.com/google/uuid"
)

type UseCase interface {
	VerifyUser(phoneOrEmail, password string) (name string, id uuid.UUID, err error)
}
