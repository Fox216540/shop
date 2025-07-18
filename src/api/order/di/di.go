package di

import (
	"shop/src/api/product/di"
	"shop/src/app/order"
	db "shop/src/infra/db/core"
	idgen "shop/src/infra/idgenerator"
	r "shop/src/infra/order"
)

func GetOrderService() order.UseCase {
	database := db.GetDatabase()
	repo := r.NewRepository(database)
	ps := di.GetProductService()
	idGen := idgen.NewSonyFlakeGenerator()
	return order.NewOrderService(repo, ps, idGen)
}
