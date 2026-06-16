package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Fox216540/shop/payment-service/api"
	"github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/Fox216540/shop/payment-service/infra/config"
	"github.com/Fox216540/shop/payment-service/infra/db"
	dbcore "github.com/Fox216540/shop/payment-service/infra/db/core"
	"github.com/Fox216540/shop/payment-service/infra/di"
	"github.com/Fox216540/shop/payment-service/infra/logger"
	"github.com/Fox216540/shop/payment-service/infra/payment/mock"
	"github.com/Fox216540/shop/payment-service/infra/payment/yokassa"
	pbInterservice "github.com/Fox216540/shop/proto/payment-service/gen/interservice"
	yo "github.com/rvinnie/yookassa-sdk-go/yookassa"
	"google.golang.org/grpc"
)

func main() {
	logger.InitLogger()

	ctx := context.Background()
	logg := logger.NewZeroLogger()

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

	provider := selectProvider(conf)

	handler := api.NewGRPCHandler(
		di.NewPaymentModule(database, provider),
		api.NewErrorMapper(logg),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.PaymentServicePort))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pbInterservice.RegisterInterserviceServiceServer(server, handler)

	log.Printf("payment-service listening on %s", lis.Addr().String())
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func selectProvider(conf *config.Config) payment.Provider {
	if conf.PaymentProviderMode == "yookassa" {
		if conf.YoKassaAccountID == "" || conf.YoKassaSecretKey == "" {
			log.Fatal("YOOKASSA_ACCOUNT_ID and YOOKASSA_SECRET_KEY are required for payment provider mode yookassa")
		}
		client := yo.NewClient(conf.YoKassaAccountID, conf.YoKassaSecretKey)
		return yokassa.NewProvider(yo.NewPaymentHandler(client))
	}

	return mock.NewProvider()
}
