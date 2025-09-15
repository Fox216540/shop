package user

import (
	"shop/src/domain/user"
)

const layer = "Infra"

type ServerError struct {
	*user.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: user.NewGlobalError(msg, err, layer),
	}
}

type InvalidAdd struct {
	*ServerError
}

func (e *InvalidAdd) Error() string {
	return e.ServerError.Error()
}

func NewInvalidAdd(err error) error {
	return &InvalidAdd{
		ServerError: NewServerError("Invalid Add Error", err),
	}
}

type InvalidDelete struct {
	*ServerError
}

func (e *InvalidDelete) Error() string {
	return e.ServerError.Error()
}

func NewInvalidDelete(err error) error {
	return &InvalidDelete{
		ServerError: NewServerError("Invalid Delete Error", err),
	}
}

type InvalidGetByID struct {
	*ServerError
}

func (e *InvalidGetByID) Error() string {
	return e.ServerError.Error()
}

func NewInvalidGetByID(err error) error {
	return &InvalidGetByID{
		ServerError: NewServerError("Invalid Get By ID Error", err),
	}
}

type InvalidFindByPhoneOrEmail struct {
	*ServerError
}

func (e *InvalidFindByPhoneOrEmail) Error() string {
	return e.ServerError.Error()
}

func NewInvalidFindByPhoneOrEmail(err error) error {
	return &InvalidFindByPhoneOrEmail{
		ServerError: NewServerError("Invalid Find By Phone Or Email Error", err),
	}
}

type InvalidUpdate struct {
	*ServerError
}

func (e *InvalidUpdate) Error() string {
	return e.ServerError.Error()
}

func NewInvalidUpdate(err error) error {
	return &InvalidUpdate{
		ServerError: NewServerError("Invalid Update Error", err),
	}
}

type InvalidExistsPhone struct {
	*ServerError
}

func (e *InvalidExistsPhone) Error() string {
	return e.ServerError.Error()
}

func NewInvalidExistsPhone(err error) error {
	return &InvalidExistsPhone{
		ServerError: NewServerError("Invalid Exists Phone Error", err),
	}
}

type InvalidExistsEmail struct {
	*ServerError
}

func (e *InvalidExistsEmail) Error() string {
	return e.ServerError.Error()
}

func NewInvalidExistsEmail(err error) error {
	return &InvalidExistsEmail{
		ServerError: NewServerError("Invalid Exists Email Error", err),
	}
}
