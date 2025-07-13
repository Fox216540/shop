package productservice

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
)

type service struct {
	r product.Repository
}

func NewService(r product.Repository) Service {
	return &service{r: r}
}

func (s *service) ProductsOfCategory(c *string) ([]product.Product, error) {
	products, err := s.r.FindProductsByCategory(c)
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) ProductById(id uuid.UUID) (product.Product, error) {
	p, err := s.r.FindProductByID(id)
	if err != nil {
		return p, err // Возвращаем ошибку, если не удалось найти продукт
	}
	return p, nil
}
