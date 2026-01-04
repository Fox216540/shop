package logger

import "context"

type Logger interface {
	Info(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Error(ctx context.Context, err error)
	Fatal(ctx context.Context, err error)
}
