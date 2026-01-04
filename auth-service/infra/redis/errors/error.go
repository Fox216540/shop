package errors

import (
	"github.com/Fox216540/shop/auth-service/infra/globalError"
)

const domain = "Redis"

type DomainServerError struct {
	*globalError.InfraServerError
}

func (e *DomainServerError) Error() string {
	return e.InfraServerError.Error()
}

func (e *DomainServerError) Unwrap() error {
	return e.InfraServerError
}

func NewDomainServerError(msg string, err error) *DomainServerError {
	return &DomainServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type RedisError struct {
	*DomainServerError
}

func (e *RedisError) Error() string {
	return e.DomainServerError.Error()
}

func (e *RedisError) Unwrap() error {
	return e.DomainServerError
}

func NewRedisError(err error) *RedisError {
	return &RedisError{
		DomainServerError: NewDomainServerError("Redis ping failed", err),
	}
}

type CloseError struct {
	*DomainServerError
}

func (e *CloseError) Error() string {
	return e.DomainServerError.Error()
}

func (e *CloseError) Unwrap() error {
	return e.DomainServerError
}

func NewCloseError(err error) *CloseError {
	return &CloseError{
		DomainServerError: NewDomainServerError("Close Error", err),
	}
}
