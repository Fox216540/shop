package di

import (
	"shop/src/api/order/di"
	"shop/src/app/user"
	db "shop/src/infra/db/core"
	hash "shop/src/infra/hasher"
	r "shop/src/infra/user"
)

func GetUserService() user.UseCase {
	database := db.GetDatabase()
	repo := r.NewRepository(database)
	orderService := di.GetOrderService()
	hasher := hash.NewHasher()
	return user.NewUserService(repo, orderService, hasher)
}
