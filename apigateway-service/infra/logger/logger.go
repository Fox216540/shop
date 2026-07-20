package logger

import (
	"context"
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

func (l *ZeroLogger) Info(ctx context.Context, msg string) {
	l.log.Info().Msg(msg)
}

func (l *ZeroLogger) Warn(ctx context.Context, msg string) {
	l.log.Warn().Msg(msg)
}

func (l *ZeroLogger) Error(ctx context.Context, err error) {
	l.log.Error().Caller().Stack().Err(err).Send()
}

func (l *ZeroLogger) Fatal(ctx context.Context, err error) {
	l.log.Fatal().Caller().Stack().Err(err).Send()
}
