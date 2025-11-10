package user

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/google/uuid"
)

type Client interface {
	RegisterUser(user User) (name string, tokens auth.Tokens, message string, err error)
	DeleteUser(id uuid.UUID) (message string, err error)
	UpdateEmail(id uuid.UUID, email string) (message string, err error)
	UpdatePassword(id uuid.UUID, password string) (message string, err error)
	UpdatePhone(id uuid.UUID, phone string) (message string, err error)
	UpdateProfile(id uuid.UUID, name *string, address *string) (message string, err error)
}
