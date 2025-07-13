package productrepository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/product"
	db "shop/src/infra/db/core"
	models "shop/src/infra/product/models"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) product.Repository {
	return &repository{db: db}
}

func (r *repository) FindProductsByCategory(category string) ([]product.Product, error) {
	var productsORM []models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("category = ?", category).Find(&productsORM).Error
	})
	if err != nil {
		return []product.Product{}, err // Возвращаем ошибку, если не удалось найти продукты
	}
	products := make([]product.Product, len(productsORM))
	for i, p := range productsORM {
		products[i] = models.FromORM(p)
	}

	return products, nil
}

func (r *repository) FindProductByID(productID uuid.UUID) (product.Product, error) {
	var p models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", productID).First(&p)
		if result.RowsAffected == 0 {
			// TODO: Поменять на кастомную ошибку
			return fmt.Errorf("product not found")
		}
		return result.Error // Возвращаем ошибку, если не удалось найти продукт
	})

	if err != nil {
		return product.Product{}, err // Возвращаем ошибку, если не удалось найти продукт
	}
	return models.FromORM(p), nil
}
