package app

import (
	categoryDomain "github.com/Fox216540/shop/catalog-service/domain/category"
	productDomain "github.com/Fox216540/shop/catalog-service/domain/product"
	"github.com/google/uuid"
)

type service struct {
	categoryRepo categoryDomain.Repository
	productRepo  productDomain.Repository
}

func NewService(categoryRepo categoryDomain.Repository, productRepo productDomain.Repository) UseCase {
	return &service{
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
	}
}

func (s *service) GetCategories() ([]categoryDomain.Category, error) {
	categories, err := s.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *service) GetAllProducts() ([]productDomain.Product, error) {
	products, err := s.productRepo.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetProductsOfCategoryID(categoryID uuid.UUID) ([]productDomain.Product, error) {
	products, err := s.productRepo.FindProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetProductByID(productID uuid.UUID) (productDomain.Product, error) {
	product, err := s.productRepo.FindProductByID(productID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) GetProductsByIDs(ids []uuid.UUID) ([]productDomain.Product, error) {
	products, err := s.productRepo.FindProductsByIDs(ids)
	if err != nil {
		return nil, err
	}

	if len(products) != len(ids) {
		return nil, productDomain.NewNotFoundProductsError(nil)
	}

	return products, nil
}
