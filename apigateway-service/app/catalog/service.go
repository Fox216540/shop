package catalog

import (
	"github.com/Fox216540/shop/apigateway-service/domain/catalog"
	"github.com/google/uuid"
)

type service struct {
	catalogClient catalog.Client
}

func NewService(catalogClient catalog.Client) UseCase {
	return &service{
		catalogClient: catalogClient,
	}
}

func (s *service) GetCategories() ([]catalog.Category, error) {
	return s.catalogClient.GetAllCategories()
}

func (s *service) GetProducts() ([]catalog.Product, error) {
	return s.catalogClient.GetAllProducts()
}

func (s *service) GetProductsOfCategoryID(categoryID uuid.UUID) ([]catalog.Product, error) {
	return s.catalogClient.GetProductsOfCategoryID(categoryID)
}

func (s *service) GetProductByID(ID uuid.UUID) (catalog.Product, error) {
	return s.catalogClient.GetProductByID(ID)
}
