package di

import (
	"github.com/Fox216540/shop/basket-service/app"
	"github.com/Fox216540/shop/basket-service/infra/basket"
	"github.com/Fox216540/shop/basket-service/infra/client"
	"github.com/Fox216540/shop/basket-service/infra/db/core"
	infraProductShort "github.com/Fox216540/shop/basket-service/infra/productShort"
)

func GetBasketService(database *core.Database, catalogClient *client.GRPCClient) app.UseCase {
	repo := basket.NewRepository(database)
	productClient := infraProductShort.NewGRPCClient(catalogClient)
	return app.NewService(repo, productClient)
}
