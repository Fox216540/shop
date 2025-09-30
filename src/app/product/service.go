package product

import (
	"errors"
	"github.com/google/uuid"
	"shop/src/core/exception"
	"shop/src/domain/product"
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

func (s *service) ProductsOfCategoryID(ID *uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByCategoryID(ID)
	if err != nil {
		return nil, s.mapError(err, NewInvalidProductsOfCategoryID(err)) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) ProductByID(ID uuid.UUID) (product.Product, error) {
	p, err := s.r.FindProductByID(ID)
	if err != nil {
		return p, s.mapError(err, NewInvalidProductByID(err)) // Возвращаем ошибку, если не удалось найти продукт
	}
	return p, nil
}

func (s *service) ProductsByIDs(IDs []uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByIDs(IDs)
	if err != nil {
		return nil, s.mapError(err, NewInvalidProductsByIDs(err)) // Возвращаем ошибку, если не удалось найти продукты
	}

	if len(products) != len(IDs) {
		return nil, product.NewNotFoundProductError(nil) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}
