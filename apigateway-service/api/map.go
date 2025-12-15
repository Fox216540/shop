package api

import (
	"errors"
	"github.com/Fox216540/shop/apigateway-service/core/exception"
	"github.com/Fox216540/shop/apigateway-service/core/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type HTTPMapper struct {
	log logger.Logger
}

func (m *HTTPMapper) MapError(e error) (statusCode int, message string) {
	var (
		ue *exception.UnauthorizedError
		se *exception.ServerError
	)

	// --------------------------------
	// 1) Ваши доменные ошибки из HTTP
	// --------------------------------

	if errors.As(e, &ue) {
		m.log.Info(e.Error())
		return http.StatusUnauthorized, NewMessages().Unauthorized
	}

	if errors.As(e, &se) {
		m.log.Info(e.Error())
		m.log.Error(e.Error())
		return http.StatusInternalServerError, NewMessages().ServerError
	}

	// --------------------------------
	// 2) gRPC ошибка
	// --------------------------------

	if st, ok := status.FromError(e); ok {
		m.log.Info("gRPC error: " + st.Message())

		switch st.Code() {
		case codes.InvalidArgument:
			return http.StatusBadRequest, st.Message()

		case codes.NotFound:
			return http.StatusNotFound, st.Message()

		case codes.Unauthenticated:
			return http.StatusUnauthorized, st.Message()

		case codes.PermissionDenied:
			return http.StatusForbidden, st.Message()

		case codes.Unavailable:
			return http.StatusServiceUnavailable, "service unavailable"

		default:
			return http.StatusInternalServerError, st.Message()
		}
	}

	// --------------------------------
	// 3) Неизвестная ошибка
	// --------------------------------

	m.log.Info("Unknown Server Error")
	m.log.Error(e.Error())
	return http.StatusInternalServerError, NewMessages().ServerError
}
