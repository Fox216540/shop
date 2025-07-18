package di

import (
	"shop/src/api/order/di"
	"shop/src/app/user"
	db "shop/src/infra/db/core"
	r "shop/src/infra/user"
)

func GetUserService() user.UseCase {
	database := db.GetDatabase()
	repo := r.NewRepository(database)
	orderService := di.GetOrderService()
	return user.NewUserService(repo, orderService)
}
