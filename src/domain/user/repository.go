package user

import "github.com/google/uuid"

type Repository interface {
	Add(u User) (User, error)
	Delete(ID uuid.UUID) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (User, error)
	FindByUsernameOrEmail(usernameOrEmail string) (User, error)
	Update(u User) (User, error)
	ExistsPhone(phone string) (bool, error)
	ExistsUsernameOrEmail(usernameOrEmail string) (bool, error)
}
