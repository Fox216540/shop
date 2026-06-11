package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pbApi "github.com/Fox216540/shop/proto/user-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/user-service/gen/interservice"
	"github.com/Fox216540/shop/user-service/api"
	"github.com/Fox216540/shop/user-service/infra/client"
	"github.com/Fox216540/shop/user-service/infra/config"
	"github.com/Fox216540/shop/user-service/infra/db"
	dbcore "github.com/Fox216540/shop/user-service/infra/db/core"
	"github.com/Fox216540/shop/user-service/infra/di"
	"github.com/Fox216540/shop/user-service/infra/logger"
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

	authClient := client.NewClient(fmt.Sprintf("localhost:%d", conf.AuthServicePort))
	defer authClient.Close()

	handler := api.NewGRPCHandler(
		di.GetUserService(database, authClient),
		api.NewErrorMapper(logg),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.UserServicePort))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pbApi.RegisterApiServiceServer(server, handler)
	pbInterservice.RegisterInterserviceServiceServer(server, handler)

	log.Printf("user-service listening on %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
