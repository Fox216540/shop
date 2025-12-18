package tokenDecoder

import "github.com/google/uuid"

type UseCase interface {
	DecodeAccessToken(token string) (userID uuid.UUID, err error)
}
