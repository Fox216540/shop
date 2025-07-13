package productservice

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
)

type Service interface {
	ProductsOfCategory(category *string) ([]product.Product, error)
	ProductById(productId uuid.UUID) (product.Product, error)
}
