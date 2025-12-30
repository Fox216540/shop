package api

import (
	"context"
	"errors"
	"github.com/Fox216540/shop/user-service/core/exception"
	"github.com/Fox216540/shop/user-service/core/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorMapper struct {
	errorLog logger.Logger
}

func NewErrorMapper(errorLog logger.Logger) *ErrorMapper {
	return &ErrorMapper{errorLog: errorLog}
}

func (m *ErrorMapper) MapError(ctx context.Context, e error) error {
	var bre *exception.BadRequestError
	var ue *exception.UnauthorizedError
	var nfe *exception.NotFoundError

	// -------------------------------
	// 1) gRPC ошибка → прокинуть вверх
	// -------------------------------
	if _, ok := status.FromError(e); ok {
		return e
	}

	// -------------------------------
	// 2) Ваши ошибки → маппинг
	// -------------------------------
	var code codes.Code
	switch {
	case errors.As(e, &bre):
		code = codes.InvalidArgument

	case errors.As(e, &ue):
		code = codes.Unauthenticated

	case errors.As(e, &nfe):
		code = codes.NotFound

	default:
		m.errorLog.Error(ctx, e)
		code = codes.Internal
	}

	return status.Error(code, e.Error())
}
