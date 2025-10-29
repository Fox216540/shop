package catalog

import "github.com/google/uuid"

type Client interface {
	GetCategories() ([]Category, error)
	GetAllProducts() ([]Product, error)
	GetProductsOfCategoryID(categoryID uuid.UUID) ([]Product, error)
	GetProductByID(ID uuid.UUID) (Product, error)
}
