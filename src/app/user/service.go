package user

import (
	"errors"
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	orderservice "shop/src/app/order"
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

func (s *service) existsEmailAndPhone(email, phone string) (bool, error) {
	exists, err := s.r.ExistsEmail(email)
	if err != nil {
		return false, err // Return error if unable to check if email exists
	}
	if exists {
		return false, errors.New("email already exists") // Return error if email already exists
	}
	exists, err = s.r.ExistsPhone(phone)
	if err != nil {
		return false, err // Return error if unable to check if phone exists
	}
	if exists {
		return false, errors.New("phone already exists") // Return error if phone already exists
	}
	return true, nil
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

	exists, err := s.existsEmailAndPhone(u.Email, u.Phone) // Check if email or phone already exists(u.Email)
	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to check if email or phone exists
	}
	if !exists {
		return u, jwt.Tokens{}, errors.New("email or phone already exists") // Return error if email or phone already exists
	}

	passwordHashed, err := s.h.Hash(u.Password)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to hash password
	}

	u.Password = passwordHashed

	u, err = s.r.Add(u)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to add user
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	jwtTokens := jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ttl := settings.Config.RefreshTokenTTL
	ttlDuration, err := s.toDuration(ttl)
	if err != nil {
		return u, jwt.Tokens{}, err
	}
	if err = s.ts.Set(refreshJTI, u.ID, ttlDuration); err != nil {
		return user.User{}, jwt.Tokens{}, err
	}

	return u, jwtTokens, nil
}

func (s *service) Login(phoneOrEmail, password string) (user.User, jwt.Tokens, error) {

	u, err := s.r.FindByPhoneOrEmail(phoneOrEmail)
	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to find user
	}

	if err = s.h.Verify(password, u.Password); err != nil {
		return u, jwt.Tokens{}, err // Password is invalid
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	jwtTokens := jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ttl := settings.Config.RefreshTokenTTL
	ttlDuration, err := s.toDuration(ttl)
	if err != nil {
		return u, jwt.Tokens{}, err
	}

	if err = s.ts.Set(refreshJTI, u.ID, ttlDuration); err != nil {
		return user.User{}, jwt.Tokens{}, err
	}

	return u, jwtTokens, nil
}

func (s *service) Logout(token string) error {
	jwtUser, err := s.jwt.DecodeRefreshToken(token)
	// TODO: Вернуть кастомную ошибку
	if err != nil {
		return err
	}

	ok, err := s.ts.Exists(jwtUser.JTI)
	// TODO: Вернуть кастомную ошибку
	if err != nil {
		return err
	}
	// TODO: Вернуть кастомную ошибку
	if !ok {
		return err
	}

	_, err = s.r.GetByID(jwtUser.UserID)
	// TODO: Вернуть кастомную ошибку
	if err != nil {
		return err // Вернуть ошибку, если не удалось найти пользователя
	}
	// TODO: Вернуть кастомную ошибку
	if err = s.ts.Delete(jwtUser.JTI, jwtUser.UserID); err != nil {
		return err
	}

	return nil
}

func (s *service) LogoutAll(token string) error {
	jwtUser, err := s.jwt.DecodeRefreshToken(token)
	if err != nil {
		return err
	}

	_, err = s.r.GetByID(jwtUser.UserID)
	if err != nil {
		return err // Вернуть ошибку, если не удалось найти пользователя
	}

	ok, err := s.ts.Exists(jwtUser.JTI)
	// TODO: Вернуть кастомную ошибку
	if err != nil {
		return err
	}
	// TODO: Вернуть кастомную ошибку
	if !ok {
		return err
	}

	if err = s.ts.DeleteAll(jwtUser.UserID); err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateEmail(userID uuid.UUID, newEmail string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось найти пользователя
	}
	isExists, err := s.r.ExistsEmail(newEmail)
	if isExists {
		//TODO: Вернуть кастомную ошибку
		return user.User{}, errors.New("email already exists") // Вернуть ошибку, если пользователь с таким именем уже существует
	}
	u.Email = newEmail
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdatePassword(userID uuid.UUID, newPassword string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось найти пользователя
	}
	if err = s.h.Verify(newPassword, u.Password); err == nil {
		return user.User{}, err
	}
	hashedPassword, err := s.h.Hash(newPassword)
	if err != nil {
		return user.User{}, err
	}
	u.Password = hashedPassword
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdatePhone(userID uuid.UUID, newPhone string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось найти пользователя
	}
	exists, err := s.r.ExistsPhone(newPhone)
	if exists {
		return user.User{}, err
	}
	u.Phone = newPhone
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdateProfile(userID uuid.UUID, userDTO dto.UpdateProfileRequest) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось найти пользователя
	}
	if userDTO.NewName != "" {
		u.Name = userDTO.NewName
	}
	if userDTO.NewAddress != "" {
		u.Address = userDTO.NewAddress
	}
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) RefreshTokens(token string) (jwt.Tokens, error) {
	jwtUser, err := s.jwt.DecodeRefreshToken(token)
	if err != nil {
		return jwt.Tokens{}, err
	}
	_, err = s.r.GetByID(jwtUser.UserID)
	if err != nil {
		return jwt.Tokens{}, err // Вернуть ошибку, если не удалось найти пользователя
	}

	if err = s.ts.Delete(jwtUser.JTI, jwtUser.UserID); err != nil {
		return jwt.Tokens{}, err
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(jwtUser.UserID)
	if err != nil {
		return jwt.Tokens{}, err // Вернуть ошибку, если не удалось обновить токен
	}

	ttl := settings.Config.RefreshTokenTTL
	ttlDuration, err := s.toDuration(ttl)
	if err != nil {
		return jwt.Tokens{}, err
	}

	if err = s.ts.Set(refreshJTI, jwtUser.UserID, ttlDuration); err != nil {
		return jwt.Tokens{}, err
	}

	accessToken, err := s.jwt.GenerateAccessToken(jwtUser.UserID)
	if err != nil {
		return jwt.Tokens{}, err // Вернуть ошибку, если не удалось обновить токен
	}

	return jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *service) Delete(userID uuid.UUID) (user.User, error) {
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

func (s *service) DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (order.Order, error) {
	o, err := s.o.GetByID(orderID)
	if err != nil {
		return order.Order{}, err // Return error if unable to find order
	}
	_, err = s.o.Cancel(orderID, userID)
	if err != nil {
		return order.Order{}, err // Return error if unable to delete order
	}
	return o, nil
}

func (s *service) CreateOrder(userID uuid.UUID, userDTO dto.CreateOrderRequest) (order.Order, error) {
	productItems := make([]*order.Item, len(userDTO.OrderItems)) // userDTO.OrderItems -> userDTO.OrderItems
	for i, item := range userDTO.OrderItems {
		uuidID, err := uuid.Parse(item.ProductID)
		if err != nil {
			return order.Order{}, err // Return error if unable to parse product ID
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
		return order.Order{}, err // Return error if unable to save order
	}

	return newOrder, nil
}
