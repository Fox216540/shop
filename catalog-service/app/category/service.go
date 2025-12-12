package category

import (
	"github.com/Fox216540/shop/catalog-service/app/mapError"
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
		return nil, mapError.MapError(err, NewInvalidGetCategories(err))
	}
	return categories, nil
}
