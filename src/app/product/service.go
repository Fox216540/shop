package productservice

import (
	"errors"
	"github.com/google/uuid"
	"shop/src/domain/product"
)

type service struct {
	r product.Repository
}

func NewService(r product.Repository) UseCase {
	return &service{r: r}
}

func (s *service) ProductsOfCategoryID(ID *uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByCategoryID(ID)
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) ProductByID(ID uuid.UUID) (product.Product, error) {
	p, err := s.r.FindProductByID(ID)
	if err != nil {
		return p, err // Возвращаем ошибку, если не удалось найти продукт
	}
	return p, nil
}

func (s *service) ValidateProductsByIDs(IDs []uuid.UUID) error {
	products, err := s.r.FindProductsByIDs(IDs)
	if err != nil {
		return err // Возвращаем ошибку, если не удалось найти продукты
	}

	if len(products) != len(IDs) {
		return errors.New("some products not found") // Возвращаем ошибку, если не удалось найти продукты
	}
	return nil
}
