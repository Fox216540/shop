package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/user"
)

type UserORM struct {
	gorm.Model
	ID       int       `gorm:"primaryKey;autoIncrement"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Email    string    `gorm:"type:varchar(50);not null;uniqueIndex"`
	Name     string    `gorm:"type:varchar(20);not null"`
	Password string    `gorm:"type:varchar(100);not null"`
	Phone    string    `gorm:"type:varchar(20);not null;uniqueIndex"`
	Address  string    `gorm:"type:varchar(100);not null"`
}

func (UserORM) TableName() string {
	return "users"
}

func FromORM(orm UserORM) user.User {
	return user.User{
		ID:       orm.UserID,
		Email:    orm.Email,
		Name:     orm.Name,
		Password: orm.Password,
		Address:  orm.Address,
		Phone:    orm.Phone,
	}
}
