package productdi

import (
	productservice "shop/src/app/product"
	"shop/src/domain/product"
	db "shop/src/infra/db/core"
	productrepository "shop/src/infra/product"
)

func GetProductService() product.Service {
	database := db.GetDatabase() // Получаем базу данных из пакета db
	repo := productrepository.NewRepository(database)
	return productservice.NewService(repo)
}
