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
	categories, err := s.catalogClient.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *service) GetProducts() ([]catalog.Product, error) {
	products, err := s.catalogClient.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetProductsOfCategoryID(categoryID uuid.UUID) ([]catalog.Product, error) {
	products, err := s.catalogClient.GetProductsOfCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetProductByID(ID uuid.UUID) (catalog.Product, error) {
	product, err := s.catalogClient.GetProductByID(ID)
	if err != nil {
		return catalog.Product{}, err
	}
	return product, nil
}
