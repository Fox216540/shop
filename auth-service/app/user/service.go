package user

import (
	"github.com/Fox216540/shop/auth-service/domain/user"
	"github.com/google/uuid"
)

type Service struct {
	u user.Client
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) VerifyUser(phoneOrEmail, password string) (name string, id uuid.UUID, err error) {
	return s.u.VerifyCredentialsOfUser(phoneOrEmail, password)
}
