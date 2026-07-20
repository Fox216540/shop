package main

import (
	"fmt"
	"log"

	shopApiGen "github.com/Fox216540/shop/api/gen"
	"github.com/Fox216540/shop/apigateway-service/api"
	"github.com/Fox216540/shop/apigateway-service/infra/config"
	"github.com/Fox216540/shop/apigateway-service/infra/di"
	"github.com/Fox216540/shop/apigateway-service/infra/logger"
	"github.com/Fox216540/shop/apigateway-service/infra/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger.InitLogger()
	logg := logger.NewStdLogger()

	metrics := di.ProvideMetrics()
	handler, clients := di.NewHTTPHandler(conf, api.NewHTTPMapper(logg), metrics)
	defer clients.Close()

	tokenDecoder, authClient := di.NewTokenDecoder(conf)
	defer authClient.Close()

	router := gin.Default()
	router.Use(middleware.NewMetricsMiddleware(metrics).Metrics())

	shopApiGen.RegisterHandlersWithOptions(router, handler, shopApiGen.GinServerOptions{
		Middlewares: []shopApiGen.MiddlewareFunc{
			shopApiGen.MiddlewareFunc(middleware.NewJWTMiddleware(tokenDecoder).Security()),
		},
		ErrorHandler: func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, shopApiGen.MessageResponse{Message: err.Error()})
		},
	})

	addr := fmt.Sprintf(":%d", conf.APIGatewayPort)
	logg.Info("apigateway-service listening on " + addr)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
