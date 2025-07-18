package product

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/product"
	db "shop/src/infra/db/core"
	"shop/src/infra/product/models"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) product.Repository {
	return &repository{db: db}
}

func (r *repository) FindProductsByCategoryID(ID *uuid.UUID) ([]product.Product, error) {
	var productsORM []models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		if ID == nil {
			// Если категория не указана, возвращаем все продукты
			return tx.
				Preload("Category").
				Find(&productsORM).Error
		} else {
			// Иначе ищем продукты по указанной категории
			// Используем Where для фильтрации по категории
			return tx.
				Joins("Category").
				Where(`"Category"."category_id" = ?`, *ID).
				Preload("Category").
				Find(&productsORM).Error
		}
	})
	if err != nil {
		return []product.Product{}, err // Возвращаем ошибку, если не удалось найти продукты
	}
	products := make([]product.Product, 0, len(productsORM))
	for _, p := range productsORM {
		products = append(products, models.FromORM(p))
	}

	return products, nil
}

func (r *repository) FindProductByID(ID uuid.UUID) (product.Product, error) {
	var p models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		result := tx.
			Preload("Category").
			Where("product_id = ?", ID).
			First(&p)
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

func (r *repository) FindProductsByIDs(IDs []uuid.UUID) ([]product.Product, error) {
	var productsORM []models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Where("product_id IN (?)", IDs).
			Find(&productsORM).Error
	})
	if err != nil {
		return []product.Product{}, err // Возвращаем ошибку, если не удалось найти продукты
	}

	products := make([]product.Product, 0, len(productsORM))
	for _, p := range productsORM {
		products = append(products, models.FromORM(p))
	}

	return products, nil
}
