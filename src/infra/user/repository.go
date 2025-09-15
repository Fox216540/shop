package user

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/user"
	db "shop/src/infra/db/core"
	"shop/src/infra/user/models"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) user.Repository {
	return &repository{db: db}
}

func (r *repository) Add(u user.User) (user.User, error) {
	newUser := &models.UserORM{
		UserID:   u.ID,
		Email:    u.Email,
		Name:     u.Name,
		Password: u.Password,
		Address:  u.Address,
		Phone:    u.Phone,
	}
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Create(&newUser).Error
	})
	if err != nil {
		return user.User{}, NewInvalidAdd(err) // Возвращаем ошибку, если не удалось сохранить пользователя
	}
	return models.FromORM(*newUser), nil
}

func (r *repository) Delete(ID uuid.UUID) (uuid.UUID, error) {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.Unscoped().Where("user_id = ?", ID).Delete(&models.UserORM{})
		if result.RowsAffected == 0 {
			return user.NewNotFoundUserError(nil) // Возвращаем ошибку, если пользователь не найден
		}
		return NewInvalidDelete(result.Error) // Возвращаем ошибку, если не удалось удалить пользователя
	})

	if err != nil {
		return uuid.Nil, err // Возвращаем ошибку, если не удалось удалить пользователя
	}
	return ID, nil
}

func (r *repository) GetByID(ID uuid.UUID) (user.User, error) {
	var u models.UserORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("user_id = ?", ID).First(&u).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, user.NewNotFoundUserError(err)
		}
		return user.User{}, NewInvalidGetByID(err) // Возвращаем ошибку, если не удалось найти пользователя
	}
	return models.FromORM(u), nil
}

func (r *repository) FindByPhoneOrEmail(phoneOrEmail string) (user.User, error) {
	var u models.UserORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("phone = ? OR email = ?", phoneOrEmail, phoneOrEmail).First(&u).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, user.NewNotFoundUserError(err)
		}
		return user.User{}, NewInvalidFindByPhoneOrEmail(err) // Возвращаем ошибку, если не удалось найти пользователя
	}
	return models.FromORM(u), nil
}

func (r *repository) Update(u user.User) (user.User, error) {
	updateUser := &models.UserORM{
		UserID:   u.ID,
		Email:    u.Email,
		Name:     u.Name,
		Password: u.Password,
		Address:  u.Address,
		Phone:    u.Phone,
	}
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Model(&models.UserORM{}).
			Where("user_id = ?", updateUser.UserID).
			Updates(map[string]interface{}{
				"email":    updateUser.Email,
				"name":     updateUser.Name,
				"password": updateUser.Password,
				"address":  updateUser.Address,
				"phone":    updateUser.Phone,
			}).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user.User{}, user.NewNotFoundUserError(err)
		}
		return user.User{}, NewInvalidUpdate(err) // Возвращаем ошибку, если не удалось обновить пользователя
	}
	return models.FromORM(*updateUser), nil
}

func (r *repository) ExistsPhone(phone string) (bool, error) {
	var u models.UserORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("phone = ?", phone).First(&u).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, NewInvalidExistsPhone(err) // Возвращаем ошибку, если не удалось найти пользователя
	}
	return true, nil
}

func (r *repository) ExistsEmail(email string) (bool, error) {
	var u models.UserORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("email = ?", email).First(&u).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, NewInvalidExistsEmail(err) // Возвращаем ошибку, если не удалось найти пользователя
	}
	return true, nil
}
