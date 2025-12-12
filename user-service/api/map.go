package api

import (
	"errors"
	"github.com/Fox216540/shop/user-service/core/exception"
	"github.com/Fox216540/shop/user-service/core/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorMapper struct {
	log      logger.Logger
	errorLog logger.Logger
}

func NewErrorMapper(log, errorLog logger.Logger) *ErrorMapper {
	return &ErrorMapper{log: log, errorLog: errorLog}
}

func (m *ErrorMapper) MapError(e error) error {
	var be *exception.BadRequestError
	var ue *exception.UnauthorizedError
	var ne *exception.NotFoundError

	// -------------------------------
	// 1) gRPC ошибка → прокинуть вверх
	// -------------------------------
	if _, ok := status.FromError(e); ok {
		return e
	}

	// -------------------------------
	// 2) Ваши ошибки → маппинг
	// -------------------------------
	switch {
	case errors.As(e, &be):
		m.log.Error(e.Error())
		return status.Error(codes.InvalidArgument, NewMessages().InvalidArgument)

	case errors.As(e, &ue):
		m.log.Error(e.Error())
		return status.Error(codes.Unauthenticated, NewMessages().Unauthorized)

	case errors.As(e, &ne):
		m.log.Error(e.Error())
		return status.Error(codes.NotFound, NewMessages().NotFound)

	default:
		m.log.Error("Server Error")
		m.errorLog.Error(e.Error())
		return status.Error(codes.Internal, NewMessages().ServerError)
	}
}
