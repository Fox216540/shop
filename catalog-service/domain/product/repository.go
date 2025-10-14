package product

import (
	"github.com/google/uuid"
)

type Repository interface {
	FindAllProducts() ([]Product, error)
	FindProductsByCategoryID(ID uuid.UUID) ([]Product, error)
	FindProductByID(ID uuid.UUID) (Product, error)
	FindProductsByIDs(IDs []uuid.UUID) ([]Product, error)
}
