package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Fox216540/shop/catalog-service/api"
	"github.com/Fox216540/shop/catalog-service/infra/config"
	"github.com/Fox216540/shop/catalog-service/infra/db"
	dbcore "github.com/Fox216540/shop/catalog-service/infra/db/core"
	"github.com/Fox216540/shop/catalog-service/infra/di"
	"github.com/Fox216540/shop/catalog-service/infra/logger"
	pbApi "github.com/Fox216540/shop/proto/catalog-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/catalog-service/gen/interservice"
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

	handler := api.NewGRPCHandler(
		di.GetCatalogService(database),
		api.NewErrorMapper(logg),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.CatalogServicePort))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pbApi.RegisterApiServiceServer(server, handler)
	pbInterservice.RegisterInterserviceServiceServer(server, handler)

	log.Printf("catalog-service listening on %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
