package productmodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/product"
)

type ProductORM struct {
	gorm.Model
	ID          int       `gorm:"primaryKey;autoIncrement"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;default:gen_random_uuid();uniqueIndex"`
	Name        string    `gorm:"not null;unique"`             // Unique product name
	Img         string    `gorm:"not null"`                    // Product image URL
	Price       float64   `gorm:"type:decimal(10,2);not null"` // Product price
	Category    string    `gorm:"not null"`                    // Product category
	Description string    `gorm:"not null"`                    // Product description
	Stock       int       `gorm:"not null"`                    // Product stock quantity
}

func FromORM(orm ProductORM) product.Product {
	return product.Product{
		ID:          orm.ProductID,
		Name:        orm.Name,
		Img:         orm.Img,
		Price:       orm.Price,
		Category:    orm.Category,
		Description: orm.Description,
		Stock:       orm.Stock,
	}
}
