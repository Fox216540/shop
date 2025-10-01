package mapError

import (
	"errors"
	"net/http"
	"shop/src/core/exception"
	"shop/src/core/logger"
)

func MapError(e error) (int, string) {
	var be *exception.BadRequestError
	var ue *exception.UnauthorizedError
	var ne *exception.NotFoundError
	switch {
	case errors.As(e, &be):
		logger.Log.Error(e.Error())
		return http.StatusBadRequest, "Bad request"
	case errors.As(e, &ue):
		logger.Log.Error(e.Error())
		return http.StatusUnauthorized, "Unauthorized"
	case errors.As(e, &ne):
		logger.Log.Error(e.Error())
		return http.StatusNotFound, "Not found"
	default:
		logger.Log.Error("Server Error")
		logger.ErrorLog.Error(e.Error())
		return http.StatusInternalServerError, "Server Error"
	}
}
