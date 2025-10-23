package di

import (
	"github.com/Fox216540/shop/user-service/app/user"
	"github.com/Fox216540/shop/user-service/infra/auth"
	"github.com/Fox216540/shop/user-service/infra/client"
	"github.com/Fox216540/shop/user-service/infra/db/core"
	h "github.com/Fox216540/shop/user-service/infra/hasher"
	r "github.com/Fox216540/shop/user-service/infra/user"
)

func GetUserService(db *core.Database, client *client.GRPCClient) user.UseCase {
	repo := r.NewRepository(db)
	a := auth.NewGRPCClient(client)
	hRepo := h.NewHasher()
	return user.NewUserService(repo, a, hRepo)
}
