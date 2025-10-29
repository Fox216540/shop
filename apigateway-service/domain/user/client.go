package user

import "github.com/Fox216540/shop/apigateway-service/domain/auth"

type Client interface {
	RegisterUser(user User) (name string, tokens auth.Tokens, message string, err error)
	LogIn(phoneOrEmail, password string) (name string, tokens auth.Tokens, message string, err error)
}
