package app

import (
	"github.com/Fox216540/shop/user-service/app/dto"
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

func NewService(
	r user.Repository,
	auth auth.Client,
	h hasher.UseCase,
) UseCase {
	return &service{
		r:    r,
		auth: auth,
		h:    h,
	}
}

func (s *service) Register(u user.User) (user.User, auth.Tokens, error) {
	if err := s.r.ExistsByEmail(u.Email); err != nil {
		return u, auth.Tokens{}, err
	}
	if err := s.r.ExistsByPhone(u.Phone); err != nil {
		return u, auth.Tokens{}, err
	}
	passwordHashed, err := s.h.HashPass(u.Password)
	if err != nil {
		return u, auth.Tokens{}, err
	}

	u.Password = passwordHashed

	u, err = s.r.Create(u)
	if err != nil {
		return u, auth.Tokens{}, err
	}

	tokens, err := s.auth.GenerateTokens(u.ID)
	if err != nil {
		return u, auth.Tokens{}, err
	}

	return u, tokens, nil
}

func (s *service) VerifyCredentials(phoneOrEmail, password string) (name string, id uuid.UUID, error error) {
	u, err := s.r.FindByPhoneOrEmail(phoneOrEmail, phoneOrEmail)
	if err != nil {
		return "", uuid.Nil, err
	}

	if err = s.h.VerifyPass(password, u.Password); err != nil {
		return "", uuid.Nil, err
	}

	return u.Name, u.ID, nil
}

func (s *service) UpdateEmail(userID uuid.UUID, newEmail string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err
	}
	if err = s.r.ExistsByEmail(newEmail); err != nil {
		return user.User{}, err
	}
	u.Email = newEmail
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (s *service) UpdatePassword(userID uuid.UUID, newPassword string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err
	}
	if err = s.h.VerifyPass(newPassword, u.Password); err == nil {
		return user.User{}, user.NewExistingPasswordError(nil)
	}
	hashedPassword, err := s.h.HashPass(newPassword)
	if err != nil {
		return user.User{}, err
	}
	u.Password = hashedPassword
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (s *service) UpdatePhone(userID uuid.UUID, newPhone string) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err
	}
	if err = s.r.ExistsByPhone(newPhone); err != nil {
		return user.User{}, err
	}
	u.Phone = newPhone
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (s *service) UpdateProfile(userID uuid.UUID, userDTO dto.ProfileUpdate) (user.User, error) {
	u, err := s.r.GetByID(userID)
	if err != nil {
		return user.User{}, err
	}
	if userDTO.Name != nil {
		if *userDTO.Name == "" {
			return user.User{}, user.NewNameCannotBeEmptyError(nil)
		}
		u.Name = *userDTO.Name
	}

	if userDTO.Address != nil {
		if *userDTO.Address == "" {
			return user.User{}, user.NewAddressCannotBeEmptyError(nil)
		}
		u.Address = *userDTO.Address
	}
	u, err = s.r.Update(u)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func (s *service) DeleteUser(userID uuid.UUID) error {
	if _, err := s.r.GetByID(userID); err != nil {
		return err
	}
	if err := s.r.Delete(userID); err != nil {
		return err
	}
	if err := s.auth.DeleteAllRefreshTokens(userID); err != nil {
		return err
	}
	return nil
}
