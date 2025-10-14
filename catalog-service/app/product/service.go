package product

import (
	"errors"
	"github.com/Fox216540/shop/catalog-service/core/exception"
	"github.com/Fox216540/shop/catalog-service/domain/product"
	"github.com/google/uuid"
)

type service struct {
	r product.Repository
}

func NewService(r product.Repository) UseCase {
	return &service{r: r}
}

func (s *service) mapError(err, appServerError error) error {
	var domainError *exception.DomainError
	var serverError *exception.ServerError
	if errors.As(err, &domainError) {
		return err
	}
	if errors.As(err, &serverError) {
		return err
	}
	return appServerError
}

func (s *service) GetProducts() ([]product.Product, error) {
	products, err := s.r.FindAllProducts()
	if err != nil {
		return nil, s.mapError(err, NewInvalidGetAll(err)) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) GetProductsOfCategoryID(ID uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByCategoryID(ID)
	if err != nil {
		return nil, s.mapError(err, NewInvalidGetProductsOfCategoryID(err)) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) GetProductByID(ID uuid.UUID) (product.Product, error) {
	p, err := s.r.FindProductByID(ID)
	if err != nil {
		return p, s.mapError(err, NewInvalidGetProductByID(err)) // Возвращаем ошибку, если не удалось найти продукт
	}
	return p, nil
}

func (s *service) GetProductsByIDs(IDs []uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByIDs(IDs)
	if err != nil {
		return nil, s.mapError(err, NewInvalidGetProductsByIDs(err)) // Возвращаем ошибку, если не удалось найти продукты
	}

	if len(products) != len(IDs) {
		return nil, product.NewNotFoundProductError(nil) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}
