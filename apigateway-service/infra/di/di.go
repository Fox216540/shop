package di

import (
	"fmt"

	"github.com/Fox216540/shop/apigateway-service/api"
	authApp "github.com/Fox216540/shop/apigateway-service/app/auth"
	basketApp "github.com/Fox216540/shop/apigateway-service/app/basket"
	catalogApp "github.com/Fox216540/shop/apigateway-service/app/catalog"
	orderApp "github.com/Fox216540/shop/apigateway-service/app/order"
	tokenDecoderApp "github.com/Fox216540/shop/apigateway-service/app/tokenDecoder"
	userApp "github.com/Fox216540/shop/apigateway-service/app/user"
	"github.com/Fox216540/shop/apigateway-service/core/metrics"
	authInfra "github.com/Fox216540/shop/apigateway-service/infra/auth"
	basketInfra "github.com/Fox216540/shop/apigateway-service/infra/basket"
	catalogInfra "github.com/Fox216540/shop/apigateway-service/infra/catalog"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	"github.com/Fox216540/shop/apigateway-service/infra/config"
	infrMetrics "github.com/Fox216540/shop/apigateway-service/infra/metrics"
	orderInfra "github.com/Fox216540/shop/apigateway-service/infra/order"
	userInfra "github.com/Fox216540/shop/apigateway-service/infra/user"
)

type Clients struct {
	Auth    *client.GRPCClient
	Basket  *client.GRPCClient
	Catalog *client.GRPCClient
	Order   *client.GRPCClient
	User    *client.GRPCClient
}

func ProvideMetrics() metrics.Metrics {
	return infrMetrics.NewPrometheusMetrics(
		infrMetrics.Registrations,
		infrMetrics.LoginFailures,
		infrMetrics.OrdersTotal,
		infrMetrics.OrderProcessing,
		infrMetrics.HTTPRequests,
		infrMetrics.HTTPErrors,
		infrMetrics.HTTPRequestDuration,
	)
}

func NewClients(conf *config.Config) Clients {
	return Clients{
		Auth:    client.NewClient(fmt.Sprintf("localhost:%d", conf.AuthServicePort)),
		Basket:  client.NewClient(fmt.Sprintf("localhost:%d", conf.BasketServicePort)),
		Catalog: client.NewClient(fmt.Sprintf("localhost:%d", conf.CatalogServicePort)),
		Order:   client.NewClient(fmt.Sprintf("localhost:%d", conf.OrderServicePort)),
		User:    client.NewClient(fmt.Sprintf("localhost:%d", conf.UserServicePort)),
	}
}

func (c Clients) Close() {
	_ = c.Auth.Close()
	_ = c.Basket.Close()
	_ = c.Catalog.Close()
	_ = c.Order.Close()
	_ = c.User.Close()
}

func NewHTTPHandler(conf *config.Config, mapper *api.HTTPMapper, metrics metrics.Metrics) (*api.HTTPHandler, Clients) {
	clients := NewClients(conf)

	authClient := authInfra.NewGRPCClient(clients.Auth)
	basketClient := basketInfra.NewGRPCClient(clients.Basket)
	catalogClient := catalogInfra.NewGRPCClient(clients.Catalog)
	orderClient := orderInfra.NewGRPClient(clients.Order)
	userClient := userInfra.NewGRPCClient(clients.User)

	return api.NewHTTPHandler(
		authApp.NewService(authClient),
		basketApp.NewService(basketClient),
		catalogApp.NewService(catalogClient),
		userApp.NewService(userClient),
		orderApp.NewService(orderClient),
		mapper,
		metrics,
		int(conf.RefreshTokenTTL),
	), clients
}

func NewTokenDecoder(conf *config.Config) (tokenDecoderApp.UseCase, *client.GRPCClient) {
	authConn := client.NewClient(fmt.Sprintf("localhost:%d", conf.AuthServicePort))
	return tokenDecoderApp.NewService(authInfra.NewGRPCClient(authConn)), authConn
}
