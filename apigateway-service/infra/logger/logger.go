package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	log zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("service", "apigateway-service").
		Logger()

	return &ZeroLogger{
		log: logger,
	}
}

func (l *ZeroLogger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *ZeroLogger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *ZeroLogger) Error(msg string) {
	l.log.Error().Caller().Msg(msg)
}
