package migration

import (
	category "github.com/Fox216540/shop/catalog-service/infra/category/models"
	product "github.com/Fox216540/shop/catalog-service/infra/product/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	//if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
	//	return fmt.Errorf("failed to create extension pgcrypto: %w", err)
	//}

	return db.AutoMigrate(
		&category.CategoryORM{},
		&product.ProductORM{})
}
