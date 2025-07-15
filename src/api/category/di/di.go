package categorydi

import (
	service "shop/src/app/category"
	repo "shop/src/infra/category"
	db "shop/src/infra/db/core"
)

func GetCategoryService() service.UseCase {
	database := db.GetDatabase()
	r := repo.NewRepository(database)
	return service.NewService(r)
}
