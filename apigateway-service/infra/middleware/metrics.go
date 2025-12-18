package middleware

import (
	"github.com/Fox216540/shop/apigateway-service/core/metrics"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type MetricsMiddleware struct {
	metrics metrics.Metrics
}

func NewMetricsMiddleware(metrics metrics.Metrics) *MetricsMiddleware {
	return &MetricsMiddleware{
		metrics: metrics,
	}
}

func (m *MetricsMiddleware) Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.FullPath()

		// fallback, если роут не найден
		if path == "" {
			path = "unknown"
		}

		m.metrics.ObserveHTTPRequestDuration(
			path,
			method,
			strconv.Itoa(status),
			duration,
		)

		m.metrics.IncHTTPRequests(
			path,
			method,
			strconv.Itoa(status),
		)

		if status >= http.StatusBadRequest {
			m.metrics.IncHTTPErrors(
				path,
				method,
				status,
			)
		}
	}
}
