package user

import (
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	"shop/src/domain/jwt"
	"shop/src/domain/order"
	"shop/src/domain/user"
)

type UseCase interface {
	Register(userDto dto.RegisterRequest) (user.User, jwt.Tokens, error)
	Login(usernameOrEmail, password string) (user.User, jwt.Tokens, error)
	Logout(token string) error
	LogoutAll(token string) error
	UpdateEmail(userID uuid.UUID, newEmail string) (user.User, error)
	UpdatePassword(userID uuid.UUID, newPassword string) (user.User, error)
	UpdatePhone(userID uuid.UUID, newPhone string) (user.User, error)
	UpdateProfile(userID uuid.UUID, userDTO dto.UpdateProfileRequest) (user.User, error)
	Delete(userID uuid.UUID) (user.User, error)
	Orders(userID uuid.UUID) ([]order.Order, error)
	DeleteOrder(userID, orderID uuid.UUID) (order.Order, error)
	CreateOrder(userID uuid.UUID, userDTO dto.CreateOrderRequest) (order.Order, error)
}
