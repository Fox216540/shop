package di

import (
	"github.com/Fox216540/shop/apigateway-service/core/metrics"
	infrMetrics "github.com/Fox216540/shop/apigateway-service/infra/metrics"
)

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
