package productdi

import (
	productservice "shop/src/app/product"
	db "shop/src/infra/db/core"
	productrepository "shop/src/infra/product"
)

func GetProductService() productservice.Service {
	database := db.GetDatabase() // Получаем базу данных из пакета db
	repo := productrepository.NewRepository(database)
	return productservice.NewService(repo)
}
