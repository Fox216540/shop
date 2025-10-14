package product

import (
	"github.com/Fox216540/shop/catalog-service/domain/product"
	"github.com/google/uuid"
)

type UseCase interface {
	GetProducts() ([]product.Product, error)
	GetProductsOfCategoryID(category uuid.UUID) ([]product.Product, error)
	GetProductByID(productId uuid.UUID) (product.Product, error)
	GetProductsByIDs(IDs []uuid.UUID) ([]product.Product, error)
}
