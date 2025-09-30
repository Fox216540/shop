package user

import "shop/src/app/globalError"

const domain = "User"

type CategoryServerError struct {
	*globalError.AppServerError
}

func (e *CategoryServerError) Error() string {
	return e.AppServerError.Error()
}

func NewCategoryServerError(msg string, err error) *CategoryServerError {
	return &CategoryServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidRegister struct {
	*CategoryServerError
}

func (e *InvalidRegister) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidRegister(err error) error {
	return &InvalidRegister{
		CategoryServerError: NewCategoryServerError("Invalid Register", err),
	}
}

type InvalidLogin struct {
	*CategoryServerError
}

func (e *InvalidLogin) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidLogin(err error) error {
	return &InvalidLogin{
		CategoryServerError: NewCategoryServerError("Invalid Login", err),
	}
}

type InvalidLogout struct {
	*CategoryServerError
}

func (e *InvalidLogout) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidLogout(err error) error {
	return &InvalidLogout{
		CategoryServerError: NewCategoryServerError("Invalid Logout", err),
	}
}

type InvalidLogoutAll struct {
	*CategoryServerError
}

func (e *InvalidLogoutAll) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidLogoutAll(err error) error {
	return &InvalidLogoutAll{
		CategoryServerError: NewCategoryServerError("Invalid Logout All", err),
	}
}

type InvalidUpdateEmail struct {
	*CategoryServerError
}

func (e *InvalidUpdateEmail) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidUpdateEmail(err error) error {
	return &InvalidUpdateEmail{
		CategoryServerError: NewCategoryServerError("Invalid Update Email", err),
	}
}

type InvalidUpdatePassword struct {
	*CategoryServerError
}

func (e *InvalidUpdatePassword) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidUpdatePassword(err error) error {
	return &InvalidUpdatePassword{
		CategoryServerError: NewCategoryServerError("Invalid Update Password", err),
	}
}

type InvalidUpdatePhone struct {
	*CategoryServerError
}

func (e *InvalidUpdatePhone) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidUpdatePhone(err error) error {
	return &InvalidUpdatePhone{
		CategoryServerError: NewCategoryServerError("Invalid Update Phone", err),
	}
}

type InvalidUpdateProfile struct {
	*CategoryServerError
}

func (e *InvalidUpdateProfile) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidUpdateProfile(err error) error {
	return &InvalidUpdateProfile{
		CategoryServerError: NewCategoryServerError("Invalid Update Profile", err),
	}
}

type InvalidRefreshToken struct {
	*CategoryServerError
}

func (e *InvalidRefreshToken) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidRefreshToken(err error) error {
	return &InvalidRefreshToken{
		CategoryServerError: NewCategoryServerError("Invalid Refresh Token", err),
	}
}

type InvalidDelete struct {
	*CategoryServerError
}

func (e *InvalidDelete) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidDelete(err error) error {
	return &InvalidDelete{
		CategoryServerError: NewCategoryServerError("Invalid Delete", err),
	}
}

type InvalidOrders struct {
	*CategoryServerError
}

func (e *InvalidOrders) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidOrders(err error) error {
	return &InvalidOrders{
		CategoryServerError: NewCategoryServerError("Invalid Orders", err),
	}
}

type InvalidDeleteOrder struct {
	*CategoryServerError
}

func (e *InvalidDeleteOrder) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidDeleteOrder(err error) error {
	return &InvalidDeleteOrder{
		CategoryServerError: NewCategoryServerError("Invalid Delete Order", err),
	}
}

type InvalidCreateOrder struct {
	*CategoryServerError
}

func (e *InvalidCreateOrder) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidCreateOrder(err error) error {
	return &InvalidCreateOrder{
		CategoryServerError: NewCategoryServerError("Invalid Create Order", err),
	}
}
