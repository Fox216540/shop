package mapError

import (
	"errors"
	"github.com/Fox216540/shop/auth-service/core/exception"
	"github.com/Fox216540/shop/auth-service/core/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapError(e error) error {
	var be *exception.BadRequestError
	var ue *exception.UnauthorizedError
	var ne *exception.NotFoundError
	switch {
	case errors.As(e, &be):
		logger.Log.Error(e.Error())
		return status.Error(codes.InvalidArgument, NewMessages().InvalidArgument)
	case errors.As(e, &ue):
		logger.Log.Error(e.Error())
		return status.Error(codes.Unauthenticated, NewMessages().Unauthorized)
	case errors.As(e, &ne):
		logger.Log.Error(e.Error())
		return status.Error(codes.NotFound, NewMessages().NotFound)
	default:
		logger.Log.Error("Server Error")
		logger.ErrorLog.Error(e.Error())
		return status.Error(codes.Internal, NewMessages().ServerError)
	}
}
