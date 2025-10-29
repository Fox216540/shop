package app

import (
	"github.com/Fox216540/shop/user-service/app/dto"
	"github.com/Fox216540/shop/user-service/domain/auth"
	"github.com/Fox216540/shop/user-service/domain/user"
	"github.com/google/uuid"
)

type UseCase interface {
	Register(user user.User) (user.User, auth.Tokens, error)
	VerifyCredentials(usernameOrEmail, password string) (name string, id uuid.UUID, err error)
	UpdateEmail(userID uuid.UUID, newEmail string) (user.User, error)
	UpdatePassword(userID uuid.UUID, newPassword string) (user.User, error)
	UpdatePhone(userID uuid.UUID, newPhone string) (user.User, error)
	UpdateProfile(userID uuid.UUID, userDTO dto.ProfileUpdate) (user.User, error)
	DeleteUser(userID uuid.UUID) error
}
