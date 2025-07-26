package user

import (
	"fmt"
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	orderservice "shop/src/app/order"
	"shop/src/domain/hasher"
	"shop/src/domain/jwt"
	"shop/src/domain/order"
	"shop/src/domain/product"
	"shop/src/domain/tokenstorage"
	"shop/src/domain/user"
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

func (s *service) Register(userDto dto.RegisterRequest) (user.User, jwt.Tokens, error) {
	//_, err := s.r.FindByUsernameOrEmail(userDto.Username)
	//if err == nil {
	//	return user.User{}, jwt.Tokens{}, err // Return error if user with the same username already exists
	//}
	//_, err = s.r.FindByUsernameOrEmail(userDto.Email)
	//if err == nil {
	//	return user.User{}, jwt.Tokens{}, err // Return error if user with the same email already exists
	//}

	u := user.User{
		ID:       uuid.New(),
		Username: userDto.Username,
		Email:    userDto.Email,
		Name:     userDto.Name,
		Password: userDto.Password,
		Address:  userDto.Address,
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

	fmt.Println(u.ID)

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(u.ID)

	fmt.Println(refreshToken)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID, u.Username)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	jwtTokens := jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err = s.ts.Set(refreshJTI, u.ID); err != nil {
		return user.User{}, jwt.Tokens{}, err
	}

	return u, jwtTokens, nil
}

func (s *service) Login(usernameOrEmail, password string) (user.User, jwt.Tokens, error) {

	u, err := s.r.FindByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to find user
	}

	isValidPass, err := s.h.Verify(password, u.Password)

	if !isValidPass {
		return u, jwt.Tokens{}, err // Password is invalid
	}

	if err != nil {
		return u, jwt.Tokens{}, err // Assuming ErrInvalidCredentials is defined in the user package
	}

	refreshToken, refreshJTI, err := s.jwt.GenerateRefreshToken(u.ID)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID, u.Username)

	if err != nil {
		return u, jwt.Tokens{}, err // Return error if unable to generate tokens
	}

	jwtTokens := jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err = s.ts.Set(refreshJTI, u.ID); err != nil {
		return user.User{}, jwt.Tokens{}, err
	}

	return u, jwtTokens, nil
}

func (s *service) Logout(userID uuid.UUID, token string) error {
	_, err := s.r.GetByID(userID)
	if err != nil {
		return err // Вернуть ошибку, если не удалось найти пользователя
	}
	//TODO: проверить токен
	if err = s.ts.Delete(uuid.New()); err != nil {
		return err
	}

	return nil
}

func (s *service) LogoutAll(userID uuid.UUID) error {
	_, err := s.r.GetByID(userID)
	if err != nil {
		return err // Вернуть ошибку, если не удалось найти пользователя
	}

	if err = s.ts.DeleteAll(userID); err != nil {
		return err
	}
	return nil
}

func (s *service) Update(userID uuid.UUID, u user.User) (user.User, error) {
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

func (s *service) DeleteOrder(userDTO dto.TestDeleteOrderRequest) (order.Order, error) {
	orderID, err := uuid.Parse(userDTO.ID)
	if err != nil {
		return order.Order{}, err // Return error if unable to parse order ID
	}

	userID, err := uuid.Parse(userDTO.UserID)
	if err != nil {
		return order.Order{}, err // Return error if unable to parse user ID
	}

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

func (s *service) CreateOrder(userDTO dto.TestCreateOrderRequest) (order.Order, error) {
	userID, err := uuid.Parse(userDTO.ID)
	if err != nil {
		return order.Order{}, err // Return error if unable to parse user ID
	}

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
