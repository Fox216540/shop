package orderdi

import (
	orderservice "shop/src/app/order"
	"shop/src/domain/order"
	db "shop/src/infra/db/core"
	orderrepo "shop/src/infra/order"
)

func GetOrderService() order.Service {
	database := db.GetDatabase()
	repo := orderrepo.NewRepository(database)
	return orderservice.NewOrderService(repo)
}
