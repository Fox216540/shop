package product

import (
	"github.com/google/uuid"
)

type Repository interface {
	FindProductsByCategory(category *string) ([]Product, error)
	FindProductByID(productId uuid.UUID) (Product, error)
}
