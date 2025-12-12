package product

import (
	"github.com/Fox216540/shop/catalog-service/app/mapError"
	"github.com/Fox216540/shop/catalog-service/domain/product"
	"github.com/google/uuid"
)

type service struct {
	r product.Repository
}

func NewService(r product.Repository) UseCase {
	return &service{r: r}
}

func (s *service) GetProducts() ([]product.Product, error) {
	products, err := s.r.FindAllProducts()
	if err != nil {
		return nil, mapError.MapError(err, NewInvalidGetAll(err)) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) GetProductsOfCategoryID(ID uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByCategoryID(ID)
	if err != nil {
		return nil, mapError.MapError(err, NewInvalidGetProductsOfCategoryID(err)) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}

func (s *service) GetProductByID(ID uuid.UUID) (product.Product, error) {
	p, err := s.r.FindProductByID(ID)
	if err != nil {
		return p, mapError.MapError(err, NewInvalidGetProductByID(err)) // Возвращаем ошибку, если не удалось найти продукт
	}
	return p, nil
}

func (s *service) GetProductsByIDs(IDs []uuid.UUID) ([]product.Product, error) {
	products, err := s.r.FindProductsByIDs(IDs)
	if err != nil {
		return nil, mapError.MapError(err, NewInvalidGetProductsByIDs(err)) // Возвращаем ошибку, если не удалось найти продукты
	}

	if len(products) != len(IDs) {
		return nil, product.NewNotFoundProductError(nil) // Возвращаем ошибку, если не удалось найти продукты
	}
	return products, nil
}
