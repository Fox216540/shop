package productmodel

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
	ordermodels "shop/src/infra/order/models"
)

type ProductORM struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Name        string    `gorm:"not null;unique"`             // Unique product name
	Img         string    `gorm:"not null"`                    // Product image URL
	Price       float64   `gorm:"type:decimal(10,2);not null"` // Product price
	Category    string    `gorm:"not null"`                    // Product category
	Description string    `gorm:"not null"`                    // Product description
	Stock       int       `gorm:"not null"`                    // Product stock quantity

	ProductItems []*ordermodels.ProductItemORM `gorm:"foreignKey:ProductID;references:ProductID;constraint:OnDelete:CASCADE"`
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
