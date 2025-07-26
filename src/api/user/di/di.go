package di

import (
	"shop/src/api/order/di"
	"shop/src/app/user"
	db "shop/src/infra/db/core"
	hash "shop/src/infra/hasher"
	"shop/src/infra/jwt"
	"shop/src/infra/tokenstorage"
	r "shop/src/infra/user"
)

func GetUserService() user.UseCase {
	database := db.GetDatabase()
	tokenStorage := db.GetRedis()
	repo := r.NewRepository(database)
	orderService := di.GetOrderService()
	hasher := hash.NewHasher()
	jwtService := jwt.NewService()
	tSRepo := tokenstorage.NewRepository(tokenStorage)

	return user.NewUserService(
		repo,
		orderService,
		hasher,
		jwtService,
		tSRepo,
	)
}
