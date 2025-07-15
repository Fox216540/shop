package productmodel

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
	categorymodels "shop/src/infra/category/models"
)

type ProductORM struct {
	ID          int                        `gorm:"primaryKey;autoIncrement"`
	ProductID   uuid.UUID                  `gorm:"type:uuid;not null;default:gen_random_uuid();uniqueIndex"`
	Name        string                     `gorm:"not null;unique"`             // Unique product name
	Img         string                     `gorm:"not null"`                    // Product image URL
	Price       float64                    `gorm:"type:decimal(10,2);not null"` // Product price
	CategoryID  uuid.UUID                  `gorm:"type:uuid;not null;index"`    // Product category
	Category    categorymodels.CategoryORM `gorm:"references:CategoryID"`       // Product category
	Description string                     `gorm:"not null"`                    // Product description
	Stock       int                        `gorm:"not null"`                    // Product stock quantity
}

func (ProductORM) TableName() string {
	return "products"
}

func FromORM(orm ProductORM) product.Product {
	return product.Product{
		ID:          orm.ProductID,
		Name:        orm.Name,
		Img:         orm.Img,
		Price:       orm.Price,
		CategoryID:  orm.Category.CategoryID,
		Description: orm.Description,
		Stock:       orm.Stock,
	}
}
