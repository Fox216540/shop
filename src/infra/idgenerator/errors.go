package idgenerator

import (
	"shop/src/infra/globalError"
)

const domain = "IDGenerator"

type IDGeneratorServerError struct {
	*globalError.InfraServerError
}

func (e *IDGeneratorServerError) Error() string {
	return e.InfraServerError.Error()
}

func (e *IDGeneratorServerError) Unwrap() error {
	return e.InfraServerError
}

func NewIDGeneratorServerError(msg string, err error) *IDGeneratorServerError {
	return &IDGeneratorServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidGenerateError struct {
	*IDGeneratorServerError
}

func (e *InvalidGenerateError) Error() string {
	return e.IDGeneratorServerError.Error()
}

func (e *InvalidGenerateError) Unwrap() error {
	return e.IDGeneratorServerError
}

func NewInvalidGenerateError(err error) error {
	return &InvalidGenerateError{
		IDGeneratorServerError: NewIDGeneratorServerError("Invalid Generate New ID Error", err),
	}
}
