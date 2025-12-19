package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"time"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339Nano

	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return fmt.Sprintf("%+v", err)
	}
}
