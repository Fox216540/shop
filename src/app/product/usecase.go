package product

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
)

type UseCase interface {
	ProductsOfCategoryID(category *uuid.UUID) ([]product.Product, error)
	ProductByID(productId uuid.UUID) (product.Product, error)
	ProductsByIDs(IDs []uuid.UUID) ([]product.Product, error)
}
