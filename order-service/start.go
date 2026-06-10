package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Fox216540/shop/order-service/api"
	"github.com/Fox216540/shop/order-service/infra/client"
	"github.com/Fox216540/shop/order-service/infra/config"
	"github.com/Fox216540/shop/order-service/infra/db"
	dbcore "github.com/Fox216540/shop/order-service/infra/db/core"
	"github.com/Fox216540/shop/order-service/infra/di"
	"github.com/Fox216540/shop/order-service/infra/logger"
	pbApi "github.com/Fox216540/shop/proto/order-service/gen/api"
	"google.golang.org/grpc"
)

func main() {
	logger.InitLogger()

	logg := logger.NewZeroLogger()
	ctx := context.Background()

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.InitPostgres(ctx, logg, conf)
	if err != nil {
		log.Fatal(err)
	}

	database := dbcore.NewDatabase(dbConn)
	defer db.ClosePostgres(ctx, logg, database)

	catalogClient := client.NewClient(fmt.Sprintf("localhost:%d", conf.ProductServicePort))
	defer catalogClient.Close()

	handler := api.NewGRPCHandler(
		di.GetOrderService(database, catalogClient),
		api.NewErrorMapper(logg),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.OrderServicePort))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pbApi.RegisterApiServiceServer(server, handler)

	log.Printf("order-service listening on %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
