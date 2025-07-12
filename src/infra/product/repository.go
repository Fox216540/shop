package productrepository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/product"
	"shop/src/infra/db"
)

type repository struct {
	db *db.Database
}

func _(db *db.Database) product.Repository {
	return &repository{db: db}
}

func (r *repository) FindProductsByCategory(category string) ([]product.Product, error) {
	var products []product.Product
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.Where("category = ?", category).Find(&products).Error
	})
	if err != nil {
		return []product.Product{}, err // Возвращаем ошибку, если не удалось сохранить заказ
	}
	return products, nil
}

func (r *repository) FindProductByID(productID uuid.UUID) (product.Product, error) {
	var p product.Product
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", productID).First(&p)
		if result.RowsAffected == 0 {
			// TODO: Поменять на кастомную ошибку
			return fmt.Errorf("product not found")
		}
		if result.Error != nil {
			return result.Error // Возвращаем ошибку, если не удалось удалить заказ
		}
		return nil
	})

	if err != nil {
		return product.Product{}, err // Возвращаем ошибку, если не удалось удалить заказ
	}
	return p, nil
}
