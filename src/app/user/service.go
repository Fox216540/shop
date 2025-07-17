package userservice

import (
	"github.com/google/uuid"
	dto "shop/src/api/user/dto"
	"shop/src/app/order"
	"shop/src/domain/hasher"
	"shop/src/domain/order"
	"shop/src/domain/user"
)

type service struct {
	r user.Repository
	o orderservice.UseCase
	h hasher.Hasher
}

func NewUserService(r user.Repository, o orderservice.UseCase, h hasher.Hasher) UseCase {
	return &service{r: r, o: o, h: h}
}

func (s *service) Register(userDto dto.RegisterRequest) (user.User, error) {
	// TODO: Добавить проверку на уникальность логина и email
	u := user.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: userDto.Password,
		Address:  userDto.Address,
	}

	u.ID = uuid.New() // Assuming GenerateID is a function that generates a new user ID

	passwordHashed, err := s.h.Hash(u.Password)

	if err != nil {
		return u, err // Return error if unable to hash password
	}

	u.Password = passwordHashed

	u, err = s.r.Add(u)

	if err != nil {
		return u, err // Return error if unable to save user
	}
	return u, nil
}

func (s *service) Login(usernameOrEmail, password string) (user.User, error) {

	u, err := s.r.FindByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return u, err // Return error if unable to find user
	}

	isValidPass, err := s.h.Verify(password, u.Password)

	if !isValidPass {
		return u, err // Password is invalid
	}

	if err != nil {
		return u, err // Assuming ErrInvalidCredentials is defined in the user package
	}
	return u, nil
}

func (s *service) Logout(userID uuid.UUID) error {
	// TODO: Добавить вызов удаления токена из хранилища
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	if err := s.Logout(userID); err != nil {
	}
}

func (s *service) LogoutAll(userID uuid.UUID) error {
	// TODO: Добавить вызов удаления токена из хранилища
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	if err := s.LogoutAll(userID); err != nil {
	}
}

func (s *service) Update(userID uuid.UUID, u user.User) (user.User, error) {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	updatedUser, err := s.r.Update(u)
	if err != nil {
		return updatedUser, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return updatedUser, nil
}

func (s *service) Delete(userID uuid.UUID) error {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	if err := s.Delete(userID); err != nil {
		return err // Вернуть ошибку, если не удалось удалить пользователя
	}
	return nil
}

func (s *service) Orders(userID uuid.UUID) ([]order.Order, error) {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	orders, err := s.o.OrdersByUserID(userID)
	if err != nil {
		return nil, err // Return error if unable to find orders
	}
	return orders, nil
}

func (s *service) DeleteOrder(userID, orderID uuid.UUID) (uuid.UUID, error) {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	ID, err := s.o.CancelOrder(orderID, userID)
	if err != nil {
		return uuid.Nil, err // Return error if unable to delete order
	}
	return ID, nil
}

func (s *service) CreateOrder(userID uuid.UUID, o order.Order) (order.Order, error) {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	// TODO: Принимать DTO
	if o.ID == uuid.Nil {
		o.ID = uuid.New() // Assuming GenerateID is a function that generates a new order ID
	}
	newOrder, err := s.o.PlaceOrder(o)
	if err != nil {
		return o, err // Return error if unable to save order
	}
	return newOrder, nil
}
