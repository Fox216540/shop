package metrics

type Metrics interface {
	IncRegistration()
	IncLoginFailure()
	IncOrder()
	ObserveOrderProcessing(seconds float64)
	IncHTTPRequests(path, method, status string)
	IncHTTPErrors(path, method string, status int)
	ObserveHTTPRequestDuration(path, method, status string, seconds float64)
}
