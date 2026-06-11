package product

import (
	"errors"
	"github.com/Fox216540/shop/catalog-service/domain/product"
	db "github.com/Fox216540/shop/catalog-service/infra/db/core"
	"github.com/Fox216540/shop/catalog-service/infra/product/models"
	"github.com/google/uuid"
	pkgerrors "github.com/pkg/errors"
	"gorm.io/gorm"
)

type repository struct {
	db *db.Database
}

func NewRepository(db *db.Database) product.Repository {
	return &repository{db: db}
}

func (r *repository) FindAllProducts() ([]product.Product, error) {
	var productsORM []models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Preload("Category").
			Find(&productsORM).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []product.Product{}, product.NewNotFoundProductsError(err)
		}
		return []product.Product{}, NewInvalidFindAll(pkgerrors.WithStack(err))
	}
	products := make([]product.Product, 0, len(productsORM))
	for _, p := range productsORM {
		products = append(products, models.FromORM(p))
	}

	return products, nil
}

func (r *repository) FindProductsByCategoryID(ID uuid.UUID) ([]product.Product, error) {
	var productsORM []models.ProductORM
	err := r.db.WithSession(func(tx *gorm.DB) error {
		return tx.
			Joins("Category").
			Where(`"Category"."category_id" = ?`, ID).
			Preload("Category").
			Find(&productsORM).Error
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []product.Product{}, product.NewNotFoundProductsError(err)
		}

		return []product.Product{}, NewInvalidFindProductsByCategoryID(pkgerrors.WithStack(err))
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
			return product.NewNotFoundProductError(result.Error)
		}
		return NewInvalidFindProductByID(pkgerrors.WithStack(result.Error))
	})

	if err != nil {
		return product.Product{}, err
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []product.Product{}, product.NewNotFoundProductsError(err)
		}
		return []product.Product{}, NewInvalidFindProductsByIDs(pkgerrors.WithStack(err))
	}

	products := make([]product.Product, 0, len(productsORM))
	for _, p := range productsORM {
		products = append(products, models.FromORM(p))
	}

	return products, nil
}
