package user

import "shop/src/app/globalError"

const domain = "User"

type UserServerError struct {
	*globalError.AppServerError
}

func (e *UserServerError) Error() string {
	return e.AppServerError.Error()
}

func NewCategoryServerError(msg string, err error) *UserServerError {
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

func NewInvalidRegister(err error) error {
	return &InvalidRegister{
		UserServerError: NewCategoryServerError("Invalid Register", err),
	}
}

type InvalidLogin struct {
	*UserServerError
}

func (e *InvalidLogin) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidLogin(err error) error {
	return &InvalidLogin{
		UserServerError: NewCategoryServerError("Invalid Login", err),
	}
}

type InvalidLogout struct {
	*UserServerError
}

func (e *InvalidLogout) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidLogout(err error) error {
	return &InvalidLogout{
		UserServerError: NewCategoryServerError("Invalid Logout", err),
	}
}

type InvalidLogoutAll struct {
	*UserServerError
}

func (e *InvalidLogoutAll) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidLogoutAll(err error) error {
	return &InvalidLogoutAll{
		UserServerError: NewCategoryServerError("Invalid Logout All", err),
	}
}

type InvalidUpdateEmail struct {
	*UserServerError
}

func (e *InvalidUpdateEmail) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidUpdateEmail(err error) error {
	return &InvalidUpdateEmail{
		UserServerError: NewCategoryServerError("Invalid Update Email", err),
	}
}

type InvalidUpdatePassword struct {
	*UserServerError
}

func (e *InvalidUpdatePassword) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidUpdatePassword(err error) error {
	return &InvalidUpdatePassword{
		UserServerError: NewCategoryServerError("Invalid Update Password", err),
	}
}

type InvalidUpdatePhone struct {
	*UserServerError
}

func (e *InvalidUpdatePhone) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidUpdatePhone(err error) error {
	return &InvalidUpdatePhone{
		UserServerError: NewCategoryServerError("Invalid Update Phone", err),
	}
}

type InvalidUpdateProfile struct {
	*UserServerError
}

func (e *InvalidUpdateProfile) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidUpdateProfile(err error) error {
	return &InvalidUpdateProfile{
		UserServerError: NewCategoryServerError("Invalid Update Profile", err),
	}
}

type InvalidRefreshToken struct {
	*UserServerError
}

func (e *InvalidRefreshToken) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidRefreshToken(err error) error {
	return &InvalidRefreshToken{
		UserServerError: NewCategoryServerError("Invalid Refresh Token", err),
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
		UserServerError: NewCategoryServerError("Invalid Delete", err),
	}
}

type InvalidOrders struct {
	*UserServerError
}

func (e *InvalidOrders) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidOrders(err error) error {
	return &InvalidOrders{
		UserServerError: NewCategoryServerError("Invalid Orders", err),
	}
}

type InvalidDeleteOrder struct {
	*UserServerError
}

func (e *InvalidDeleteOrder) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidDeleteOrder(err error) error {
	return &InvalidDeleteOrder{
		UserServerError: NewCategoryServerError("Invalid Delete Order", err),
	}
}

type InvalidCreateOrder struct {
	*UserServerError
}

func (e *InvalidCreateOrder) Error() string {
	return e.UserServerError.Error()
}

func NewInvalidCreateOrder(err error) error {
	return &InvalidCreateOrder{
		UserServerError: NewCategoryServerError("Invalid Create Order", err),
	}
}
