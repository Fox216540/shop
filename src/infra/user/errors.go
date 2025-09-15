package user

import (
	"shop/src/infra/globalError"
)

const domain = "User"

type UserServerError struct {
	*globalError.InfraServerError
}

func (e *UserServerError) Error() string {
	return e.InfraServerError.Error()
}

func NewUserServerError(msg string, err error) *UserServerError {
	return &UserServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidAdd struct {
	*UserServerError
}

func (e *InvalidAdd) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidAdd(err error) error {
	return &InvalidAdd{
		UserServerError: NewUserServerError("Invalid Add Error", err),
	}
}

type InvalidDelete struct {
	*UserServerError
}

func (e *InvalidDelete) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidDelete(err error) error {
	return &InvalidDelete{
		UserServerError: NewUserServerError("Invalid Delete Error", err),
	}
}

type InvalidGetByID struct {
	*UserServerError
}

func (e *InvalidGetByID) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidGetByID(err error) error {
	return &InvalidGetByID{
		UserServerError: NewUserServerError("Invalid Get By ID Error", err),
	}
}

type InvalidFindByPhoneOrEmail struct {
	*UserServerError
}

func (e *InvalidFindByPhoneOrEmail) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidFindByPhoneOrEmail(err error) error {
	return &InvalidFindByPhoneOrEmail{
		UserServerError: NewUserServerError("Invalid Find By Phone Or Email Error", err),
	}
}

type InvalidUpdate struct {
	*UserServerError
}

func (e *InvalidUpdate) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidUpdate(err error) error {
	return &InvalidUpdate{
		UserServerError: NewUserServerError("Invalid Update Error", err),
	}
}

type InvalidExistsPhone struct {
	*UserServerError
}

func (e *InvalidExistsPhone) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidExistsPhone(err error) error {
	return &InvalidExistsPhone{
		UserServerError: NewUserServerError("Invalid Exists Phone Error", err),
	}
}

type InvalidExistsEmail struct {
	*UserServerError
}

func (e *InvalidExistsEmail) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidExistsEmail(err error) error {
	return &InvalidExistsEmail{
		UserServerError: NewUserServerError("Invalid Exists Email Error", err),
	}
}
