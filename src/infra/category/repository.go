package category

import (
	"gorm.io/gorm"
	"shop/src/domain/category"
	"shop/src/infra/category/models"
	db "shop/src/infra/db/core"
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
		return nil, err
	}

	categories := make([]category.Category, 0, len(categoriesORM))
	for _, c := range categoriesORM {
		categories = append(
			categories, models.FromORM(c),
		)
	}

	return categories, nil
}
