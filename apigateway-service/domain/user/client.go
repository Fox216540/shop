package user

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/google/uuid"
)

type Client interface {
	RegisterUser(user User) (name string, tokens auth.Tokens, message string, err error)
	DeleteUser(id uuid.UUID) (message string, err error)
}
