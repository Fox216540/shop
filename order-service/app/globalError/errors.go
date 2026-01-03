package globalError

import "github.com/Fox216540/shop/order-service/core/exception"

const (
	layer = "App"
)

type AppServerError struct {
	*exception.ServerError
}

func (e *AppServerError) Error() string {
	return e.ServerError.Error()
}

func (e *AppServerError) Unwrap() error {
	return e.ServerError
}

func NewAppServerError(msg, domain string, err error) *AppServerError {
	return &AppServerError{
		ServerError: exception.NewServerError(msg, domain, layer, err),
	}
}
