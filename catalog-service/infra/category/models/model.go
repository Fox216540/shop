package models

import (
	"github.com/Fox216540/shop/catalog-service/domain/category"
	"github.com/google/uuid"
)

type CategoryORM struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	CategoryID uuid.UUID `gorm:"type:uuid;not null;default:gen_random_uuid();uniqueIndex"`
	Name       string    `gorm:"not null;unique"`
}

func (CategoryORM) TableName() string {
	return "category"
}

func FromORM(orm CategoryORM) category.Category {
	return category.Category{
		ID:   orm.CategoryID,
		Name: orm.Name,
	}
}
