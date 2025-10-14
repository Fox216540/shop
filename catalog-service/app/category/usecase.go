package category

import "github.com/Fox216540/shop/catalog-service/domain/category"

type UseCase interface {
	GetAllCategories() ([]category.Category, error)
}
