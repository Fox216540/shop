package orderdi

import (
	productdi "shop/src/api/product/di"
	service "shop/src/app/order"
	db "shop/src/infra/db/core"
	idgen "shop/src/infra/idgenerator"
	repo "shop/src/infra/order"
)

func GetOrderService() service.UseCase {
	database := db.GetDatabase()
	r := repo.NewRepository(database)
	ps := productdi.GetProductService()
	idGen := idgen.NewSonyFlakeGenerator()
	return service.NewOrderService(r, ps, idGen)
}
