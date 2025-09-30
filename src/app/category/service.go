package category

import (
	"errors"
	"shop/src/core/exception"
	"shop/src/domain/category"
)

type service struct {
	r category.Repository
}

func NewService(r category.Repository) UseCase {
	return &service{r: r}
}

func (s *service) GetCategories() ([]category.Category, error) {
	categories, err := s.r.FindAll()
	if err != nil {
		var domainError *exception.DomainError
		var serverError *exception.ServerError
		if errors.As(err, &domainError) {
			return nil, domainError
		}
		if errors.As(err, &serverError) {
			return nil, serverError
		}
		return nil, NewInvalidGetCategories(err)
	}
	return categories, nil
}
