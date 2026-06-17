package di

import (
	"github.com/Fox216540/shop/auth-service/app"
	"github.com/Fox216540/shop/auth-service/infra/client"
	"github.com/Fox216540/shop/auth-service/infra/config"
	infraJwt "github.com/Fox216540/shop/auth-service/infra/jwt"
	"github.com/Fox216540/shop/auth-service/infra/redis/core"
	infraTokenStorage "github.com/Fox216540/shop/auth-service/infra/tokenstorage"
	infraUser "github.com/Fox216540/shop/auth-service/infra/user"
)

func GetAuthService(conf *config.Config, rdb *core.Redis, client *client.GRPCClient) app.UseCase {
	jwtRepo := infraJwt.NewService(conf.AccessTokenTTL, conf.RefreshTokenTTL, conf.AccessTokenSecret, conf.RefreshTokenSecret)
	tokenStorageRepo := infraTokenStorage.NewRepository(rdb, conf.RefreshTokenTTL)
	userClient := infraUser.NewGRPCClient(client)
	return app.NewService(jwtRepo, tokenStorageRepo, userClient)
}
