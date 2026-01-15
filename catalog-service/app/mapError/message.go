package mapError

import (
	"errors"
	"github.com/Fox216540/shop/catalog-service/core/exception"
	pkgerrors "github.com/pkg/errors"
)

func MapError(err, anotherError error) error {
	var domainError *exception.DomainError
	var serverError *exception.ServerError
	if errors.As(err, &domainError) {
		return err
	}
	if errors.As(err, &serverError) {
		return err
	}
	return pkgerrors.WithStack(anotherError)
}
