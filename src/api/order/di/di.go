package orderdi

import (
	productdi "shop/src/api/product/di"
	orderservice "shop/src/app/order"
	db "shop/src/infra/db/core"
	idgenerator "shop/src/infra/idgenerator"
	orderrepo "shop/src/infra/order"
)

func GetOrderService() orderservice.Service {
	database := db.GetDatabase()
	repo := orderrepo.NewRepository(database)
	ps := productdi.GetProductService()
	idGen := idgenerator.NewSonyFlakeGenerator()
	return orderservice.NewOrderService(repo, ps, idGen)
}
