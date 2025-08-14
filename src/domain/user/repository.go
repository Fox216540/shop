package user

import "github.com/google/uuid"

type Repository interface {
	Add(u User) (User, error)
	Delete(ID uuid.UUID) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (User, error)
	FindByPhoneOrEmail(phoneOrEmail string) (User, error)
	Update(u User) (User, error)
	ExistsPhone(phone string) (bool, error)
	ExistsEmail(email string) (bool, error)
}
