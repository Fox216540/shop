package exception

import "fmt"

type Exception struct {
	Domain  string
	Layer   string
	Message string
	Err     error
}

func (e Exception) Error() string {
	return fmt.Sprintf("Domain: %s\nLayer: %s\nMessage: %s\nCause: %v",
		e.Domain, e.Layer, e.Message, e.Err)
}

func (e Exception) Unwrap() error {
	return e.Err
}

func NewException(msg, domain, layer string, err error) *Exception {
	return &Exception{
		Message: msg,
		Err:     err,
		Layer:   layer,
		Domain:  domain,
	}
}

const (
	domainLayer = "Domain"
)

type NotFoundError struct {
	*Exception
}

func (e NotFoundError) Error() string {
	return e.Exception.Error()
}

func (e NotFoundError) Unwrap() error { return e.Exception }

func NewNotFoundError(msg, domain string, err error) *NotFoundError {
	return &NotFoundError{
		Exception: NewException(msg, domain, domainLayer, err),
	}
}

type BadRequestError struct {
	*Exception
}

func (e BadRequestError) Error() string {
	return e.Exception.Error()
}

func (e BadRequestError) Unwrap() error { return e.Exception }

func NewBadRequestError(msg, domain string, err error) *BadRequestError {
	return &BadRequestError{
		Exception: NewException(msg, domain, domainLayer, err),
	}
}

type UnauthorizedError struct {
	*Exception
}

func (e UnauthorizedError) Error() string {
	return e.Exception.Error()
}

func (e UnauthorizedError) Unwrap() error { return e.Exception }

func NewUnauthorizedError(msg, domain string, err error) *UnauthorizedError {
	return &UnauthorizedError{
		Exception: NewException(msg, domain, domainLayer, err),
	}
}

type ServerError struct {
	*Exception
}

func (e ServerError) Error() string {
	return e.Exception.Error()
}

func (e ServerError) Unwrap() error { return e.Exception }

func NewServerError(msg, domain, layer string, err error) *ServerError {
	return &ServerError{
		Exception: NewException(msg, domain, layer, err),
	}
}
