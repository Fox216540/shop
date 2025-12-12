package di

import (
	"github.com/Fox216540/shop/order-service/app"
	c "github.com/Fox216540/shop/order-service/infra/catalog"
	"github.com/Fox216540/shop/order-service/infra/client"
	"github.com/Fox216540/shop/order-service/infra/db/core"
	idgen "github.com/Fox216540/shop/order-service/infra/idgenerator"
	r "github.com/Fox216540/shop/order-service/infra/order"
)

func GetOrderService(db *core.Database, client *client.GRPCClient) order.UseCase {
	repo := r.NewRepository(db)
	idGen := idgen.NewSonyFlakeGenerator()
	catalogClient := c.NewGRPCClient(client)
	return order.NewOrderService(repo, idGen, catalogClient)
}
