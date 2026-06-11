package app

import "github.com/Fox216540/shop/user-service/app/globalError"

const domain = "User"

type UserServerError struct {
	*globalError.AppServerError
}

func (e *UserServerError) Error() string {
	return e.AppServerError.Error()
}

func (e *UserServerError) Unwrap() error {
	return e.AppServerError
}

func NewUserServerError(msg string, err error) *UserServerError {
	return &UserServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidRegister struct {
	*UserServerError
}

func (e *InvalidRegister) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidRegister) Unwrap() error {
	return e.UserServerError
}

func NewInvalidRegister(err error) error {
	return &InvalidRegister{
		UserServerError: NewUserServerError("Invalid Register", err),
	}
}

type InvalidLogin struct {
	*UserServerError
}

func (e *InvalidLogin) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidLogin) Unwrap() error {
	return e.UserServerError
}

func NewInvalidLogin(err error) error {
	return &InvalidLogin{
		UserServerError: NewUserServerError("Invalid Login", err),
	}
}

type InvalidUpdateEmail struct {
	*UserServerError
}

func (e *InvalidUpdateEmail) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidUpdateEmail) Unwrap() error {
	return e.UserServerError
}

func NewInvalidUpdateEmail(err error) error {
	return &InvalidUpdateEmail{
		UserServerError: NewUserServerError("Invalid Update Email", err),
	}
}

type InvalidUpdatePassword struct {
	*UserServerError
}

func (e *InvalidUpdatePassword) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidUpdatePassword) Unwrap() error {
	return e.UserServerError
}

func NewInvalidUpdatePassword(err error) error {
	return &InvalidUpdatePassword{
		UserServerError: NewUserServerError("Invalid Update Password", err),
	}
}

type InvalidUpdatePhone struct {
	*UserServerError
}

func (e *InvalidUpdatePhone) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidUpdatePhone) Unwrap() error {
	return e.UserServerError
}

func NewInvalidUpdatePhone(err error) error {
	return &InvalidUpdatePhone{
		UserServerError: NewUserServerError("Invalid Update Phone", err),
	}
}

type InvalidUpdateProfile struct {
	*UserServerError
}

func (e *InvalidUpdateProfile) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidUpdateProfile) Unwrap() error {
	return e.UserServerError
}

func NewInvalidUpdateProfile(err error) error {
	return &InvalidUpdateProfile{
		UserServerError: NewUserServerError("Invalid Update Profile", err),
	}
}
