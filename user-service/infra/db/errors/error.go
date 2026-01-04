package errors

import (
	"github.com/Fox216540/shop/user-service/infra/globalError"
)

const domain = "DB"

type DomainServerError struct {
	*globalError.InfraServerError
}

func (e *DomainServerError) Error() string {
	return e.InfraServerError.Error()
}

func (e *DomainServerError) Unwrap() error {
	return e.InfraServerError
}

func NewDomainServerError(msg string, err error) *DomainServerError {
	return &DomainServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type ConnectionError struct {
	*DomainServerError
}

func (e *ConnectionError) Error() string {
	return e.DomainServerError.Error()
}

func (e *ConnectionError) Unwrap() error {
	return e.DomainServerError
}

func NewConnectionError(err error) *ConnectionError {
	return &ConnectionError{
		DomainServerError: NewDomainServerError("Connection Error", err),
	}
}

type MigrationError struct {
	*DomainServerError
}

func (e *MigrationError) Error() string {
	return e.DomainServerError.Error()
}

func (e *MigrationError) Unwrap() error {
	return e.DomainServerError
}

func NewMigrationError(err error) *MigrationError {
	return &MigrationError{
		DomainServerError: NewDomainServerError("Migration Error", err),
	}
}

type CloseError struct {
	*DomainServerError
}

func (e *CloseError) Error() string {
	return e.DomainServerError.Error()
}

func (e *CloseError) Unwrap() error {
	return e.DomainServerError
}

func NewCloseError(err error) *CloseError {
	return &CloseError{
		DomainServerError: NewDomainServerError("Close Error", err),
	}
}

type GetRawDBError struct {
	*DomainServerError
}

func (e *GetRawDBError) Error() string {
	return e.DomainServerError.Error()
}

func (e *GetRawDBError) Unwrap() error {
	return e.DomainServerError
}

func NewGetRawDBError(err error) *GetRawDBError {
	return &GetRawDBError{
		DomainServerError: NewDomainServerError("Get Raw DB Error", err),
	}
}
