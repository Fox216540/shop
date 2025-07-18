package di

import (
	"shop/src/app/category"
	r "shop/src/infra/category"
	db "shop/src/infra/db/core"
)

func GetCategoryService() category.UseCase {
	database := db.GetDatabase()
	repo := r.NewRepository(database)
	return category.NewService(repo)
}
