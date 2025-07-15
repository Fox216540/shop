package categoriseservice

import "shop/src/domain/category"

type UseCase interface {
	GetCategories() ([]category.Category, error)
}
