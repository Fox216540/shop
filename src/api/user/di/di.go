package di

import (
	jwtdi "shop/src/api/jwt/di"
	orderdi "shop/src/api/order/di"
	"shop/src/app/user"
	db "shop/src/infra/db/core"
	hash "shop/src/infra/hasher"
	"shop/src/infra/tokenstorage"
	r "shop/src/infra/user"
)

func GetUserService() user.UseCase {
	database := db.GetDatabase()
	tokenStorage := db.GetRedis()
	repo := r.NewRepository(database)
	orderService := orderdi.GetOrderService()
	hasher := hash.NewHasher()
	jwtService := jwtdi.GetJwtService()
	tSRepo := tokenstorage.NewRepository(tokenStorage)

	return user.NewUserService(
		repo,
		orderService,
		hasher,
		jwtService,
		tSRepo,
	)
}
