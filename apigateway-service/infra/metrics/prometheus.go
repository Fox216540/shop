package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
	registrations       prometheus.Counter
	loginFailures       prometheus.Counter
	ordersTotal         prometheus.Counter
	orderProcessing     prometheus.Histogram
	httpRequests        *prometheus.CounterVec
	httpRequestDuration *prometheus.HistogramVec
	httpErrors          *prometheus.CounterVec
}

func NewPrometheusMetrics(
	registrations, loginFailures, ordersTotal prometheus.Counter,
	orderProcessing prometheus.Histogram,
	httpRequests, httpErrors *prometheus.CounterVec,
	httpRequestDuration *prometheus.HistogramVec,
) *PrometheusMetrics {
	m := &PrometheusMetrics{
		registrations:       registrations,
		loginFailures:       loginFailures,
		ordersTotal:         ordersTotal,
		orderProcessing:     orderProcessing,
		httpRequests:        httpRequests,
		httpRequestDuration: httpRequestDuration,
		httpErrors:          httpErrors,
	}
	prometheus.MustRegister(
		m.registrations,
		m.loginFailures,
		m.ordersTotal,
		m.orderProcessing,
		m.httpRequests,
		m.httpRequestDuration,
		m.httpErrors,
	)
	return m
}

func (p *PrometheusMetrics) IncRegistration() {
	p.registrations.Inc()
}

func (p *PrometheusMetrics) IncLoginFailure() {
	p.loginFailures.Inc()
}

func (p *PrometheusMetrics) IncOrder() {
	p.ordersTotal.Inc()
}

func (p *PrometheusMetrics) ObserveOrderProcessing(seconds float64) {
	p.orderProcessing.Observe(seconds)
}

func (p *PrometheusMetrics) IncHTTPRequests(path, method, status string) {
	p.httpRequests.WithLabelValues(path, method, status).Inc()
}

func (p *PrometheusMetrics) IncHTTPErrors(path, method string, status int) {
	p.httpErrors.WithLabelValues(path, method, fmt.Sprintf("%d", status)).Inc()
}

func (p *PrometheusMetrics) ObserveHTTPRequestDuration(path, method, status string, seconds float64) {
	p.httpRequestDuration.WithLabelValues(path, method, status).Observe(seconds)
}
