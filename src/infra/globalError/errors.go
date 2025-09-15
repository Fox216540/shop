package globalError

import "shop/src/core/exception"

const (
	layer = "Infra"
)

type InfraServerError struct {
	*exception.ServerError
}

func (e *InfraServerError) Error() string {
	return e.ServerError.Error()
}

func NewInfraServerError(msg, domain string, err error) *InfraServerError {
	return &InfraServerError{
		ServerError: exception.NewServerError(msg, domain, layer, err),
	}
}
