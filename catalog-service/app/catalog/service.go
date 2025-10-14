package catalog

import (
	"github.com/Fox216540/shop/catalog-service/app/category"
	"github.com/Fox216540/shop/catalog-service/app/product"
	categoryDomain "github.com/Fox216540/shop/catalog-service/domain/category"
	productDomain "github.com/Fox216540/shop/catalog-service/domain/product"
	"github.com/google/uuid"
)

type service struct {
	categoryUC category.UseCase
	productUC  product.UseCase
}

func (s service) GetCategories() ([]categoryDomain.Category, error) {
	return s.categoryUC.GetAllCategories()
}

func (s service) GetAllProducts() ([]productDomain.Product, error) {
	return s.productUC.GetProducts()
}

func (s service) GetProductsOfCategoryID(category uuid.UUID) ([]productDomain.Product, error) {
	products, err := s.productUC.GetProductsOfCategoryID(category)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s service) GetProductByID(productId uuid.UUID) (productDomain.Product, error) {
	prod, err := s.productUC.GetProductByID(productId)
	if err != nil {
		return productDomain.Product{}, err
	}
	return prod, nil
}

func (s service) GetProductsByIDs(IDs []uuid.UUID) ([]productDomain.Product, error) {
	products, err := s.productUC.GetProductsByIDs(IDs)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewService(categoryUC category.UseCase, productUC product.UseCase) UseCase {
	return &service{
		categoryUC: categoryUC,
		productUC:  productUC,
	}
}
