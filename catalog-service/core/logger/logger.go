package logger

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}
