package user

import (
	"errors"
	"github.com/Fox216540/shop/user-service/app/dto"
	"github.com/Fox216540/shop/user-service/core/exception"
	"github.com/Fox216540/shop/user-service/domain/auth"
	"github.com/Fox216540/shop/user-service/domain/hasher"
	"github.com/Fox216540/shop/user-service/domain/user"
	"github.com/google/uuid"
)

type service struct {
	r    user.Repository
	auth auth.Client
	h    hasher.UseCase
}

func NewUserService(
	r user.Repository,
	auth auth.Client,
	h hasher.UseCase,
) UseCase {
	return &service{
		r: r, auth: auth, h: h}
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

func (s *service) Register(u user.User) (user.User, auth.Tokens, error) {
	if _, err := s.r.FindByPhoneOrEmail(u.Email, u.Phone); err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to check if email or phone exists
	}
	passwordHashed, err := s.h.HashPass(u.Password)

	if err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to hash password
	}

	u.Password = passwordHashed

	u, err = s.r.Create(u)

	if err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to add user
	}

	tokens, err := s.auth.GenerateTokens(u.ID)
	if err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidRegister(err)) // Return error if unable to generate tokens
	}

	return u, tokens, nil
}

func (s *service) Login(phoneOrEmail, password string) (user.User, auth.Tokens, error) {
	u, err := s.r.FindByPhoneOrEmail(phoneOrEmail, phoneOrEmail)
	if err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Return error if unable to find user
	}

	if err = s.h.VerifyPass(password, u.Password); err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Password is invalid
	}

	tokens, err := s.auth.GenerateTokens(u.ID)
	if err != nil {
		return u, auth.Tokens{}, s.mapError(err, NewInvalidLogin(err)) // Return error if unable to generate tokens
	}

	return u, tokens, nil
}

func (s *service) UpdateEmail(userID uuid.UUID, newEmail string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateEmail(err)) // Вернуть ошибку, если не удалось найти пользователя
	}
	if err = s.r.ExistsByEmail(newEmail); err != nil {
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
	if err = s.h.VerifyPass(newPassword, u.Password); err == nil {
		return user.User{}, user.NewExistingPasswordError(nil)
	}
	hashedPassword, err := s.h.HashPass(newPassword)
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
	if err = s.r.ExistsByPhone(newPhone); err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePhone(err))
	}
	u.Phone = newPhone
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdatePhone(err)) // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) UpdateProfile(userID uuid.UUID, userDTO dto.ProfileUpdate) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateProfile(err)) // Вернуть ошибку, если не удалось найти пользователя
	}
	if userDTO.Name != nil {
		if *userDTO.Name == "" {
			//TODO:Замапить ошибку
			return user.User{}, errors.New("name cannot be empty")
		}
		u.Name = *userDTO.Name
	}

	if userDTO.Address != nil {
		if *userDTO.Address == "" {
			//TODO:Замапить ошибку
			return user.User{}, errors.New("address cannot be empty")
		}
		u.Address = *userDTO.Address
	}
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, s.mapError(err, NewInvalidUpdateProfile(err)) // Вернуть ошибку, если не удалось обновить пользователя
	}
	return u, nil
}

func (s *service) DeleteUser(userID uuid.UUID) error {
	if _, err := s.r.GetByID(userID); err != nil {
		return s.mapError(err, NewInvalidDelete(err))
	}
	if err := s.r.Delete(userID); err != nil {
		return s.mapError(err, NewInvalidDelete(err)) // Вернуть ошибку, если не удалось удалить пользователя
	}
	if err := s.auth.DeleteAllRefreshTokens(userID); err != nil {
		return s.mapError(err, NewInvalidDelete(err))
	}
	return nil
}
