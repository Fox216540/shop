package user

import "github.com/Fox216540/shop/auth-service/infra/globalError"

const domain = "User"

type UserServerError struct {
	*globalError.InfraServerError
}

func (e *UserServerError) Error() string {
	return e.InfraServerError.Error()
}

func (e *UserServerError) Unwrap() error {
	return e.InfraServerError.Unwrap()
}

func NewUserServerError(msg string, err error) *UserServerError {
	return &UserServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidVerifyCredentialsOfUserError struct {
	*UserServerError
}

func (e *InvalidVerifyCredentialsOfUserError) Error() string {
	return e.UserServerError.Error()
}

func (e *InvalidVerifyCredentialsOfUserError) Unwrap() error {
	return e.UserServerError.Unwrap()
}

func NewInvalidVerifyCredentialsOfUserError(err error) *InvalidVerifyCredentialsOfUserError {
	return &InvalidVerifyCredentialsOfUserError{
		UserServerError: NewUserServerError("Invalid verify credentials of user", err),
	}
}
