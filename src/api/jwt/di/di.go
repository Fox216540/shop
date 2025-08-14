package di

import (
	"shop/src/domain/jwt"
	infra "shop/src/infra/jwt"
)

func GetJwtService() jwt.Service {
	return infra.NewService()
}
