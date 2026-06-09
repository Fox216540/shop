package logger

import (
	"errors"
	"fmt"
	pkgerrors "github.com/pkg/errors"
	"github.com/rs/zerolog"
)

func InitLogger() {
	zerolog.CallerSkipFrameCount = 3
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		for e := err; e != nil; e = errors.Unwrap(e) {
			if st, ok := e.(interface{ StackTrace() pkgerrors.StackTrace }); ok {
				trace := st.StackTrace()
				if len(trace) > 0 {
					return fmt.Sprintf("%+v", trace[0]) // repo.go:42
				}
			}
		}
		return nil
	}
}
