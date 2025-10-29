package auth

import "github.com/google/uuid"

type Client interface {
	LogOut(token string) (msg string, err error)
	LogOutAll(token string) (msg string, err error)
	RefreshTokens(token string) (tokens Tokens, err error)
	DecodeAccessToken(token string) (userID uuid.UUID, err error)
}
