package categoriseservice

import "shop/src/domain/category"

type service struct {
	r category.Repository
}

func NewService(r category.Repository) UseCase {
	return &service{r: r}
}

func (s *service) GetCategories() ([]category.Category, error) {
	categories, err := s.r.FindAll()
	if err != nil {
		return []category.Category{}, err // Возвращаем ошибку, если не удалось найти категории
	}
	return categories, nil
}
