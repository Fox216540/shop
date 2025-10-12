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

var Log *CustomLogger
var ErrorLog *CustomLogger

func InitLogger(out io.Writer) {
	Log = NewCustomLogger(out)
}

func InitErrorLogger(out io.Writer) {
	ErrorLog = NewCustomLogger(out)
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
