package user

import (
	"errors"
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	orderservice "shop/src/app/order"
	"shop/src/core/exception"
	"shop/src/core/settings"
	"shop/src/domain/hasher"
	"shop/src/domain/jwt"
	"shop/src/domain/order"
	"shop/src/domain/product"
	"shop/src/domain/tokenstorage"
	"shop/src/domain/user"
	"strconv"
	"time"
)

type service struct {
	r   user.Repository
	o   orderservice.UseCase
	h   hasher.Hasher
	jwt jwt.Service
	ts  tokenstorage.TokenStorage
}

func NewUserService(
	r user.Repository,
	o orderservice.UseCase,
	h hasher.Hasher,
	jwt jwt.Service,
	ts tokenstorage.TokenStorage,
) UseCase {
	return &service{
		r: r, o: o, jwt: jwt, ts: ts, h: h}
}

func (s *service) mapError(err, appServerError error) error {
	var domainError *exception.DomainError
	var serverError *exception.ServerError
	if errors.As(err, &domainError) {
		return err
	}
	if errors.As(err, &serverError) {
		return err
	}
	return appServerError
}

func (s *service) existsEmailAndPhone(email, phone string) error {
	err := s.r.ExistsEmail(email)
	if err != nil {
		return err // Return error if unable to check if email exists
	}
	err = s.r.ExistsPhone(phone)
	if err != nil {
		return err // Return error if unable to check if phone exists
	}
	return nil
}

func (s *service) toDuration(ttl string) (time.Duration, error) {
	seconds, err := strconv.ParseInt(ttl, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Duration(seconds) * time.Second, nil
}

func (s *service) Register(userDto dto.RegisterRequest) (user.User, jwt.Tokens, error) {
	u := user.User{
		ID:       uuid.New(),
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: userDto.Password,
		Address:  userDto.Address,
		Phone:    userDto.Phone,
	}

	err := s.existsEmailAndPhone(u.Email, u.Phone) // Check if email or phone already exists(u.Email)
	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to check if email or phone exists
	}

	passwordHashed, err := s.h.Hash(u.Password)

	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to hash password
	}

	u.Password = passwordHashed

	u, err = s.r.Add(u)

	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to add user
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to generate tokens
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to generate tokens
	}

	jwtTokens := jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ttl := settings.Config.RefreshTokenTTL
	ttlDuration, err := s.toDuration(ttl)
	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err))
	}
	if err = s.ts.Set(refreshJTI, u.ID, ttlDuration); err != nil {
		return user.User{}, jwt.Tokens{}, s.mapError(err, NewInvalidRegister(err))
	}

	return u, jwtTokens, nil
}

func (s *service) Login(phoneOrEmail, password string) (user.User, jwt.Tokens, error) {

	u, err := s.r.FindByPhoneOrEmail(phoneOrEmail)
	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Return error if unable to find user
	}

	if err = s.h.Verify(password, u.Password); err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Password is invalid
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Return error if unable to generate tokens
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Return error if unable to generate tokens
	}

	jwtTokens := jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ttl := settings.Config.RefreshTokenTTL
	ttlDuration, err := s.toDuration(ttl)
	if err != nil {
		return u, jwt.Tokens{}, s.mapError(err, NewInvalidLogin(err))
	}

	if err = s.ts.Set(refreshJTI, u.ID, ttlDuration); err != nil {
		return user.User{}, jwt.Tokens{}, s.mapError(err, NewInvalidLogin(err))
	}

	return u, jwtTokens, nil
}

func (s *service) Logout(token string) error {
	jwtUser, err := s.jwt.DecodeRefreshToken(token)

	if err != nil {
		return s.mapError(err, NewInvalidLogout(err))
	}

	err = s.ts.Exists(jwtUser.JTI)

	if err != nil {
		return s.mapError(err, NewInvalidLogout(err))
	}

	_, err = s.r.GetByID(jwtUser.UserID)

	if err != nil {
		return s.mapError(err, NewInvalidLogout(err)) // Вернуть ошибку, если не удалось найти пользователя
	}

	if err = s.ts.Delete(jwtUser.JTI, jwtUser.UserID); err != nil {
		return s.mapError(err, NewInvalidLogout(err))
	}

	return nil
}

func (s *service) LogoutAll(token string) error {
	jwtUser, err := s.jwt.DecodeRefreshToken(token)
	if err != nil {
		return s.mapError(err, NewInvalidLogoutAll(err))
	}

	_, err = s.r.GetByID(jwtUser.UserID)
	if err != nil {
		return s.mapError(err, NewInvalidLogoutAll(err)) // Вернуть ошибку, если не удалось найти пользователя
	}

	err = s.ts.Exists(jwtUser.JTI)

	if err != nil {
		return s.mapError(err, NewInvalidLogoutAll(err))
	}

	if err = s.ts.DeleteAll(jwtUser.UserID); err != nil {
		return s.mapError(err, NewInvalidLogoutAll(err))
	}
	return nil
}

func (s *service) UpdateEmail(userID uuid.UUID, newEmail string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateEmail(err)) // Вернуть ошибку, если не удалось найти пользователя
	}
	err = s.r.ExistsEmail(newEmail)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateEmail(err)) // Вернуть ошибку, если пользователь с таким именем уже существует
	}
	u.Email = newEmail
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateEmail(err)) // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdatePassword(userID uuid.UUID, newPassword string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePassword(err)) // Вернуть ошибку, если не удалось найти пользователя
	}
	if err = s.h.Verify(newPassword, u.Password); err == nil {
		return user.User{}, user.NewExistingPasswordError(nil)
	}
	hashedPassword, err := s.h.Hash(newPassword)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePassword(err))
	}
	u.Password = hashedPassword
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePassword(err)) // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdatePhone(userID uuid.UUID, newPhone string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось найти пользователя
	}
	err = s.r.ExistsPhone(newPhone)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePhone(err))
	}
	u.Phone = newPhone
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePhone(err)) // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdateProfile(userID uuid.UUID, userDTO dto.UpdateProfileRequest) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateProfile(err)) // Вернуть ошибку, если не удалось найти пользователя
	}
	if userDTO.NewName != "" {
		u.Name = userDTO.NewName
	}
	if userDTO.NewAddress != "" {
		u.Address = userDTO.NewAddress
	}
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateProfile(err)) // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) RefreshTokens(token string) (jwt.Tokens, error) {
	jwtUser, err := s.jwt.DecodeRefreshToken(token)
	if err != nil {
		return jwt.Tokens{}, s.mapError(err, NewInvalidRefreshToken(err))
	}
	_, err = s.r.GetByID(jwtUser.UserID)
	if err != nil {
		return jwt.Tokens{}, s.mapError(err, NewInvalidRefreshToken(err)) // Вернуть ошибку, если не удалось найти пользователя
	}

	if err = s.ts.Delete(jwtUser.JTI, jwtUser.UserID); err != nil {
		return jwt.Tokens{}, s.mapError(err, NewInvalidRefreshToken(err))
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(jwtUser.UserID)
	if err != nil {
		return jwt.Tokens{}, s.mapError(err, NewInvalidRefreshToken(err)) // Вернуть ошибку, если не удалось обновить токен
	}

	ttl := settings.Config.RefreshTokenTTL
	ttlDuration, err := s.toDuration(ttl)
	if err != nil {
		return jwt.Tokens{}, s.mapError(err, NewInvalidRefreshToken(err))
	}

	if err = s.ts.Set(refreshJTI, jwtUser.UserID, ttlDuration); err != nil {
		return jwt.Tokens{}, err
	}

	accessToken, err := s.jwt.GenerateAccessToken(jwtUser.UserID)
	if err != nil {
		return jwt.Tokens{}, s.mapError(err, NewInvalidRefreshToken(err)) // Вернуть ошибку, если не удалось обновить токен
	}

	return jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *service) Delete(userID uuid.UUID) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidDelete(err)) // Вернуть ошибку, если не удалось найти пользователя
	}
	_, err = s.r.Delete(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidDelete(err)) // Вернуть ошибку, если не удалось удалить пользователя
	}
	return u, nil
}

func (s *service) Orders(userID uuid.UUID) ([]order.Order, error) {
	_, err := s.r.GetByID(userID)
	if err != nil {
		return nil, s.mapError(err, NewInvalidOrders(err)) // Return error if unable to find user
	}

	orders, err := s.o.GetOrdersByUserID(userID)
	if err != nil {
		return nil, s.mapError(err, NewInvalidOrders(err)) // Return error if unable to find orders
	}
	return orders, nil
}

func (s *service) DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (order.Order, error) {
	o, err := s.o.GetByID(orderID)
	if err != nil {
		return order.Order{}, s.mapError(err, NewInvalidDeleteOrder(err)) // Return error if unable to find order
	}
	_, err = s.o.Cancel(orderID, userID)
	if err != nil {
		return order.Order{}, s.mapError(err, NewInvalidDeleteOrder(err)) // Return error if unable to delete order
	}
	return o, nil
}

func (s *service) CreateOrder(userID uuid.UUID, userDTO dto.CreateOrderRequest) (order.Order, error) {
	productItems := make([]*order.Item, len(userDTO.OrderItems)) // userDTO.OrderItems -> userDTO.OrderItems
	for i, item := range userDTO.OrderItems {
		uuidID, err := uuid.Parse(item.ProductID)
		if err != nil {
			return order.Order{}, s.mapError(err, NewInvalidCreateOrder(err)) // Return error if unable to parse product ID
		}
		productItems[i] = &order.Item{
			Product: product.Product{
				ID: uuidID,
			},
			Quantity: item.Quantity,
		}
	}

	newOrder, err := s.o.Place(userID, productItems)
	if err != nil {
		return order.Order{}, s.mapError(err, NewInvalidCreateOrder(err)) // Return error if unable to save order
	}

	return newOrder, nil
}
