package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/metadata"
)

type ZeroLogger struct {
	log zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("service", "payment-service").
		Logger()

	return &ZeroLogger{
		log: logger,
	}
}

const requestIDKey = "x-request-id"

func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	values := md.Get(requestIDKey)
	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func (l *ZeroLogger) Info(ctx context.Context, msg string) {
	requestID := RequestIDFromContext(ctx)

	e := l.log.Info()
	if requestID != "" {
		e.Str("request_id", requestID)
	}

	e.Msg(msg)
}

func (l *ZeroLogger) Warn(ctx context.Context, msg string) {
	requestID := RequestIDFromContext(ctx)

	e := l.log.Warn()
	if requestID != "" {
		e.Str("request_id", requestID)
	}

	e.Msg(msg)
}

func (l *ZeroLogger) Error(ctx context.Context, err error) {
	requestID := RequestIDFromContext(ctx)
	e := l.log.Error().
		Caller().
		Stack().
		Err(err)

	if requestID != "" {
		e.Str("request_id", requestID)
	}

	e.Send()
}

func (l *ZeroLogger) Fatal(ctx context.Context, err error) {
	e := l.log.Fatal().
		Caller().
		Stack().
		Err(err)

	if rid := RequestIDFromContext(ctx); rid != "" {
		e.Str("request_id", rid)
	}

	e.Send()
}
