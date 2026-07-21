package tokenDecoder

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/google/uuid"
)

type service struct {
	authClient auth.Client
}

func NewService(authClient auth.Client) UseCase {
	return &service{
		authClient: authClient,
	}
}
func (s *service) DecodeAccessToken(token string) (uuid.UUID, error) {
	return s.authClient.DecodeAccessTokenOfUser(token)
}
