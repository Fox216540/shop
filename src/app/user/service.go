package user

import (
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	orderservice "shop/src/app/order"
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
	_, err := s.r.FindByUsernameOrEmail(userDto.Username)
	if err == nil {
		return user.User{}, err // Return error if user with the same username already exists
	}
	_, err = s.r.FindByUsernameOrEmail(userDto.Email)
	if err == nil {
		return user.User{}, err // Return error if user with the same email already exists
	}

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
	_, err := s.r.GetByID(userID)
	if err != nil {
		return err // Вернуть ошибку, если не удалось найти пользователя
	}

	if err := s.Logout(userID); err != nil {
	}
	return nil
}

func (s *service) LogoutAll(userID uuid.UUID) error {
	// TODO: Добавить вызов удаления токена из хранилища
	_, err := s.r.GetByID(userID)
	if err != nil {
		return err // Вернуть ошибку, если не удалось найти пользователя
	}

	if err := s.LogoutAll(userID); err != nil {
	}
	return nil
}

func (s *service) Update(userID uuid.UUID, u user.User) (user.User, error) {
	// TODO: Добавить вызов обновления пользователя в хранилище
	_, err := s.r.GetByID(userID)
	if err != nil {
		return u, err // Вернуть ошибку, если не удалось найти пользователя
	}
	updatedUser, err := s.r.Update(u)
	if err != nil {
		return updatedUser, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return updatedUser, nil
}

func (s *service) Delete(userID uuid.UUID) (user.User, error) {
	// TODO: Добавить вызов удаления аккаунта из хранилища
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось найти пользователя
	}
	_, err = s.r.Delete(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось удалить пользователя
	}
	return u, nil
}

func (s *service) Orders(userID uuid.UUID) ([]order.Order, error) {
	_, err := s.r.GetByID(userID)
	if err != nil {
		return nil, err // Return error if unable to find user
	}

	orders, err := s.o.OrdersByUserID(userID)
	if err != nil {
		return nil, err // Return error if unable to find orders
	}
	return orders, nil
}

func (s *service) DeleteOrder(userID, orderID uuid.UUID) (uuid.UUID, error) {
	_, err := s.r.GetByID(userID)
	if err != nil {
		return uuid.Nil, err // Return error if unable to find user
	}
	ID, err := s.o.CancelOrder(orderID, userID)
	if err != nil {
		return uuid.Nil, err // Return error if unable to delete order
	}
	return ID, nil
}

func (s *service) CreateOrder(userID uuid.UUID, productItems []*order.Item) (order.Order, error) {
	// TODO: Принимать DTO
	newOrder, err := s.o.PlaceOrder(userID, productItems)
	if err != nil {
		return order.Order{}, err // Return error if unable to save order
	}
	return newOrder, nil
}
