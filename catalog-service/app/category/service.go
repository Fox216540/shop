package category

import (
	"errors"
	"github.com/Fox216540/shop/catalog-service/core/exception"
	"github.com/Fox216540/shop/catalog-service/domain/category"
)

type service struct {
	r category.Repository
}

func NewService(r category.Repository) UseCase {
	return &service{r: r}
}

func (s *service) GetAllCategories() ([]category.Category, error) {
	categories, err := s.r.FindAll()
	if err != nil {
		var domainError *exception.DomainError
		var serverError *exception.ServerError
		if errors.As(err, &domainError) {
			return nil, err
		}
		if errors.As(err, &serverError) {
			return nil, err
		}
		return nil, NewInvalidGetCategories(err)
	}
	return categories, nil
}
