package hasher

import (
	"fmt"
	"shop/src/core/exception"
)

const domain = "Hasher"

type GlobalError struct {
	*exception.Exception
	Domain  string
	Layer   string
	Message string
	Err     error
}

func (e *GlobalError) Error() string {
	return fmt.Sprintf("Domain: %s\nLayer: %s\nMessage: %s\nCause: %v",
		e.Domain, e.Layer, e.Message, e.Err)
}

func (e *GlobalError) Unwrap() error {
	return e.Err
}

func NewGlobalError(msg string, err error, layer string) *GlobalError {
	return &GlobalError{
		Exception: &exception.Exception{},
		Domain:    domain,
		Layer:     layer,
		Message:   msg,
		Err:       err,
	}
}
