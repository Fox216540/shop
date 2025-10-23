package user

import "github.com/google/uuid"

type Repository interface {
	Create(u User) (User, error)
	Delete(ID uuid.UUID) error
	GetByID(ID uuid.UUID) (User, error)
	FindByPhoneOrEmail(phone, email string) (User, error)
	Update(u User) (User, error)
	ExistsByPhone(phone string) error
	ExistsByEmail(email string) error
}
