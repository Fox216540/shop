package logger

import (
	"io"
	"log"
)

type CustomLogger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func InitLogger(out io.Writer) *CustomLogger {
	Log := NewCustomLogger(out)
	return Log
}

func InitErrorLogger(out io.Writer) *CustomLogger {
	ErrorLog := NewCustomLogger(out)
	return ErrorLog
}

func (l *CustomLogger) Error(s string) {
	l.error.Println(s)
}

func (l *CustomLogger) Info(s string) {
	l.info.Println(s)
}

func (l *CustomLogger) Warn(s string) {
	l.warn.Println(s)
}

func NewCustomLogger(out io.Writer) *CustomLogger {
	return &CustomLogger{
		info:  log.New(out, "INFO: ", log.Ldate|log.Ltime),
		warn:  log.New(out, "WARN: ", log.Ldate|log.Ltime),
		error: log.New(out, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
