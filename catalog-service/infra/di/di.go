package di

import (
	"github.com/Fox216540/shop/catalog-service/app"
	rCategory "github.com/Fox216540/shop/catalog-service/infra/category"
	"github.com/Fox216540/shop/catalog-service/infra/db/core"
	rProduct "github.com/Fox216540/shop/catalog-service/infra/product"
)

func GetCatalogService(database *core.Database) app.UseCase {
	categoryRepo := rCategory.NewRepository(database)
	productRepo := rProduct.NewRepository(database)
	return app.NewService(categoryRepo, productRepo)
}
