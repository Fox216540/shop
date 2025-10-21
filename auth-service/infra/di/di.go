package di

import (
	"github.com/Fox216540/shop/auth-service/app/jwt"
	"github.com/Fox216540/shop/auth-service/app/tokenstorage"
	infraJwt "github.com/Fox216540/shop/auth-service/infra/jwt"
	"github.com/Fox216540/shop/auth-service/infra/redis/core"
	infraTokenStorage "github.com/Fox216540/shop/auth-service/infra/tokenstorage"
	"time"
)

func GetJwtService(ttlRefresh, ttlAccess time.Duration, secretRefresh, secretAccess string) jwt.UseCase {
	repo := infraJwt.NewService(ttlRefresh, ttlAccess, secretRefresh, secretAccess)
	return jwt.NewService(repo)
}

func GetTokenStorageService(rdb *core.Redis, ttlStorage time.Duration) tokenstorage.UseCase {
	repo := infraTokenStorage.NewRepository(rdb, ttlStorage)
	return tokenstorage.NewService(repo)
}
