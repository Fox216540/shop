package migration

import (
	"gorm.io/gorm"
	category "shop/src/infra/category/models"
	order "shop/src/infra/order/models"
	product "shop/src/infra/product/models"
	user "shop/src/infra/user/models"
)

func Migration(db *gorm.DB) error {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
		return fmt.Errorf("failed to create extension pgcrypto: %w", err)
	}

	return db.AutoMigrate(
		&category.CategoryORM{},
		&product.ProductORM{},
		&order.OrderORM{},
		&user.UserORM{},
		&order.OrderProductORM{})
}
