package hasher

import (
	"shop/src/infra/globalError"
)

const domain = "Hasher"

type HasherServerError struct {
	*globalError.InfraServerError
}

func (e *HasherServerError) Error() string {
	return e.InfraServerError.Error()
}

func NewHasherServerError(msg string, err error) *HasherServerError {
	return &HasherServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidHashError struct {
	*HasherServerError
}

func (e *InvalidHashError) Error() string {
	return e.HasherServerError.Error()
}

func NewInvalidHashError(err error) error {
	return &InvalidHashError{
		HasherServerError: NewHasherServerError("Invalid Find All Error", err),
	}
}
