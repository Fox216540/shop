package di

import (
	"github.com/Fox216540/shop/catalog-service/app/category"
	"github.com/Fox216540/shop/catalog-service/app/product"
	rCategory "github.com/Fox216540/shop/catalog-service/infra/category"
	"github.com/Fox216540/shop/catalog-service/infra/db/core"
	rProduct "github.com/Fox216540/shop/catalog-service/infra/product"
)

func GetProductService(database *core.Database) product.UseCase {
	repo := rProduct.NewRepository(database)
	return product.NewService(repo)
}

func GetCategoryService(database *core.Database) category.UseCase {
	repo := rCategory.NewRepository(database)
	return category.NewService(repo)
}
