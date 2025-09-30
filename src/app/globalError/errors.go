package globalError

import "shop/src/core/exception"

const (
	layer = "App"
)

type AppServerError struct {
	*exception.ServerError
}

func (e *AppServerError) Error() string {
	return e.ServerError.Error()
}

func NewAppServerError(msg, domain string, err error) *AppServerError {
	return &AppServerError{
		ServerError: exception.NewServerError(msg, domain, layer, err),
	}
}
