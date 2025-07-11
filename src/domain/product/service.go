package product

import "github.com/google/uuid"

type Service interface {
	ProductsOfCategory(category string) ([]Product, error)
	ProductById(productId uuid.UUID) (Product, error)
}
