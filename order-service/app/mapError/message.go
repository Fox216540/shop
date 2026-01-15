package mapError

import (
	"errors"
	"github.com/Fox216540/shop/order-service/core/exception"
	pkgerrors "github.com/pkg/errors"
	"google.golang.org/grpc/status"
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

	// Проверка на gRPC ошибку
	if _, ok := status.FromError(err); ok {
		// это grpc-status ошибка
		return err
	}

	return pkgerrors.WithStack(anotherError)
}
