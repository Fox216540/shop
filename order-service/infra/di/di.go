package di

import (
	"github.com/Fox216540/shop/order-service/app"
	"github.com/Fox216540/shop/order-service/infra/client"
	"github.com/Fox216540/shop/order-service/infra/db/core"
	idgen "github.com/Fox216540/shop/order-service/infra/idgenerator"
	r "github.com/Fox216540/shop/order-service/infra/order"
	p "github.com/Fox216540/shop/order-service/infra/product"
)

func GetOrderService(db *core.Database, client *client.GRPCClient) order.UseCase {
	repo := r.NewRepository(db)
	idGen := idgen.NewSonyFlakeGenerator()
	productClient := p.NewGRPCClient(client)
	return order.NewOrderService(repo, idGen, productClient)
}
