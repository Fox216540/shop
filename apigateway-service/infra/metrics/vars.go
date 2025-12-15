package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	Registrations = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "shop_registrations_total",
		Help: "Total number of successful registrations",
	})

	LoginFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "shop_login_failures_total",
		Help: "Total number of failed login attempts",
	})

	OrdersTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "shop_orders_total",
		Help: "Total number of orders placed",
	})

	OrderProcessing = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "shop_order_processing_seconds",
		Help:    "Order processing duration",
		Buckets: prometheus.ExponentialBuckets(0.1, 2, 10),
	})

	HTTPRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "shop_http_requests_total",
		Help: "Total HTTP requests",
	}, []string{"path", "method", "status"})

	HTTPRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "shop_http_request_duration_seconds",
		Help:    "HTTP request duration",
		Buckets: prometheus.ExponentialBuckets(0.01, 2, 15),
	}, []string{"path", "method"})

	HTTPErrors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "shop_http_errors_total",
		Help: "Total number of HTTP error responses",
	}, []string{"path", "method", "status"})
)
