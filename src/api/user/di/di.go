package userdi

import (
	orderdi "shop/src/api/order/di"
	userservice "shop/src/app/user"
	"shop/src/domain/user"
	db "shop/src/infra/db/core"
	userrepo "shop/src/infra/user"
)

func GetUserService() user.Service {
	database := db.GetDatabase()
	r := userrepo.NewRepository(database)
	orderService := orderdi.GetOrderService()
	return userservice.NewUserService(r, orderService)
}
