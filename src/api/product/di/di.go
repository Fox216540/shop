package di

import (
	"shop/src/app/product"
	db "shop/src/infra/db/core"
	r "shop/src/infra/product"
)

func GetProductService() product.UseCase {
	database := db.GetDatabase() // Получаем базу данных из пакета db
	repo := r.NewRepository(database)
	return product.NewService(repo)
}
