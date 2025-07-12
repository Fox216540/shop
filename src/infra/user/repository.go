package userrepository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/user"
	"shop/src/infra/db"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) user.Repository {
	return &repository{db: db}
}

func (r *repository) Add(u user.User) (user.User, error) {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Create(&u).Error
	})
	if err != nil {
		return user.User{}, err // Возвращаем ошибку, если не удалось сохранить пользователя
	}
	return u, nil
}

func (r *repository) Delete(ID uuid.UUID) (uuid.UUID, error) {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", ID).Delete(&user.User{})
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound // Возвращаем ошибку, если пользователь не найден
		}
		return result.Error // Возвращаем ошибку, если не удалось удалить пользователя
	})

	if err != nil {
		return uuid.Nil, err // Возвращаем ошибку, если не удалось удалить пользователя
	}
	return ID, nil
}

func (r *repository) GetByID(ID uuid.UUID) (user.User, error) {
	var u user.User
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("id = ?", ID).First(&u).Error
	})
	if err != nil {
		return user.User{}, err // Возвращаем ошибку, если не удалось найти пользователя
	}
	return u, nil
}

func (r *repository) FindByUsernameOrEmail(usernameOrEmail string) (user.User, error) {
	var u user.User
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&u).Error
	})
	if err != nil {
		return user.User{}, err // Возвращаем ошибку, если не удалось найти пользователя
	}
	return u, nil
}

func (r *repository) Update(u user.User) (user.User, error) {
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Save(&u).Error
	})
	if err != nil {
		return user.User{}, err // Возвращаем ошибку, если не удалось обновить пользователя
	}
	return u, nil
}
