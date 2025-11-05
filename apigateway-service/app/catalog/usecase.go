package catalog

import (
	"github.com/Fox216540/shop/apigateway-service/domain/catalog"
	"github.com/google/uuid"
)

type UseCase interface {
	GetCategories() ([]catalog.Category, error)
	GetProducts() ([]catalog.Product, error)
	GetProductsOfCategoryID(categoryID uuid.UUID) ([]catalog.Product, error)
	GetProductByID(ID uuid.UUID) (catalog.Product, error)
}
