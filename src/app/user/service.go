package userservice

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
	"shop/src/domain/user"
)

type service struct {
	r user.Repository
	o order.Service
}

//TODO: Добавить хешер

func NewUserService(r user.Repository, o order.Service) user.Service {
	return &service{r: r, o: o}
}

func (s *service) Register(u user.User) (user.User, error) {
	// TODO: Добавить проверку на уникальность логина и email
	// TODO: Добавить хеширование пароля
	// TODO: Принимать данные и из них делать пользователя
	if u.ID == uuid.Nil {
		u.ID = uuid.New() // Assuming GenerateID is a function that generates a new user ID
	}
	u, err := s.r.Add(u)
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
	if u.Password != password {
		return u, err // Assuming ErrInvalidCredentials is defined in the user package
	}
	return u, nil
}

func (s *service) Logout(userID uuid.UUID) error {
	// TODO: Добавить вызов удаления токена из хранилища
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	if err := s.r.Logout(userID); err != nil {
	}
}

func (s *service) LogoutAll(userID uuid.UUID) error {
	// TODO: Добавить вызов удаления токена из хранилища
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	if err := s.r.Logout(userID); err != nil {
	}
}

func (s *service) Update(u user.User) (user.User, error) {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	updatedUser, err := s.r.Update(u)
	if err != nil {
		return updatedUser, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return updatedUser, nil
}

func (s *service) Delete(userID uuid.UUID) error {
	// TODO: Добавить проверку на есть ли пользователь с таким ID
	if err := s.r.Delete(userID); err != nil {
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
	ID, err := s.o.CancelOrder(orderID, userID)
	if err != nil {
		return uuid.Nil, err // Return error if unable to delete order
	}
	return ID, nil
}

func (s *service) CreateOrder(o order.Order) (order.Order, error) {
	// TODO: Решить с одной переменной o
	if o.ID == uuid.Nil {
		o.ID = uuid.New() // Assuming GenerateID is a function that generates a new order ID
	}
	o, err := s.o.PlaceOrder(o)
	if err != nil {
		return o, err // Return error if unable to save order
	}
	return o, nil
}
