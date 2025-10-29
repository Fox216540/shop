package auth

import "github.com/google/uuid"

type Client interface {
	LogInUser(phoneOrEmail, password string) (name string, tokens Tokens, message string, err error)
	LogOutUser(token string) (msg string, err error)
	LogOutAllUser(token string) (msg string, err error)
	RefreshTokensOfUser(token string) (tokens Tokens, err error)
	DecodeAccessTokenOfUser(token string) (userID uuid.UUID, err error)
}
