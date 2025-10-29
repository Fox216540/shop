package user

import "github.com/google/uuid"

type Client interface {
	VerifyCredentialsOfUser(phoneOrEmail, password string) (name string, id uuid.UUID, err error)
}
