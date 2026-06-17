package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Fox216540/shop/auth-service/api"
	"github.com/Fox216540/shop/auth-service/infra/client"
	"github.com/Fox216540/shop/auth-service/infra/config"
	"github.com/Fox216540/shop/auth-service/infra/di"
	"github.com/Fox216540/shop/auth-service/infra/logger"
	redisinfra "github.com/Fox216540/shop/auth-service/infra/redis"
	rediscore "github.com/Fox216540/shop/auth-service/infra/redis/core"
	pbApi "github.com/Fox216540/shop/proto/auth-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/auth-service/gen/interservice"
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

	rdb, err := redisinfra.InitRedis(ctx, logg, conf)
	if err != nil {
		log.Fatal(err)
	}
	defer redisinfra.CloseRedis(ctx, logg, rdb)

	redisCore := rediscore.NewRedis(rdb)

	userClient := client.NewClient(fmt.Sprintf("localhost:%d", conf.UserServicePort))
	defer userClient.Close()

	handler := api.NewGRPCHandler(
		di.GetAuthService(conf, redisCore, userClient),
		api.NewErrorMapper(logg),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.AuthServicePort))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pbApi.RegisterApiServiceServer(server, handler)
	pbInterservice.RegisterInterserviceServiceServer(server, handler)

	log.Printf("auth-service listening on %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
