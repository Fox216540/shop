package logger

import (
	"log"
	"os"
)

type StdLogger struct {
	log *log.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{log: log.New(os.Stdout, "apigateway-service ", log.LstdFlags)}
}

func (l *StdLogger) Info(msg string) {
	l.log.Println("INFO " + msg)
}

func (l *StdLogger) Warn(msg string) {
	l.log.Println("WARN " + msg)
}

func (l *StdLogger) Error(msg string) {
	l.log.Println("ERROR " + msg)
}
