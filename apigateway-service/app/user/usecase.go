package user

import (
	"github.com/Fox216540/shop/apigateway-service/app/dto"
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
)

type UseCase interface {
	RegisterUser(user dto.User) (name string, tokens auth.Tokens, message string, err error)
}
