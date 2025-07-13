package usermodels

import (
	"github.com/google/uuid"
	"shop/src/domain/user"
)

type UserORM struct {
	ID       int       `gorm:"primaryKey;autoIncrement"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Username string    `gorm:"type:varchar(20);not null;uniqueIndex"`
	Email    string    `gorm:"type:varchar(50);not null;uniqueIndex"`
	Name     string    `gorm:"type:varchar(20);not null"`
	Password string    `gorm:"type:varchar(100);not null"`
}

func FromORM(orm UserORM) user.User {
	return user.User{
		ID:       orm.UserID,
		Username: orm.Username,
		Email:    orm.Email,
		Name:     orm.Name,
		Password: orm.Password,
	}
}
