package category

import (
	"errors"
	"github.com/Fox216540/shop/catalog-service/domain/category"
	"github.com/Fox216540/shop/catalog-service/infra/category/models"
	db "github.com/Fox216540/shop/catalog-service/infra/db/core"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *db.Database) category.Repository {
	return &repository{db: db.DB}
}

func (r *repository) FindAll() ([]category.Category, error) {
	var categoriesORM []models.CategoryORM
	if err := r.db.Find(&categoriesORM).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, category.NewNotFoundCategoryError(err)
		}
		return nil, NewInvalidFindAllError(err)
	}

	categories := make([]category.Category, 0, len(categoriesORM))
	for _, c := range categoriesORM {
		categories = append(
			categories, models.FromORM(c),
		)
	}

	return categories, nil
}
