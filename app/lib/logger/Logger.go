package logger

import (
	"log"
	"time"
)

type Logger struct {
	Log *log.Logger
}

func NewLogger(l *log.Logger) *Logger {
	if l == nil {
		l = log.Default()
	}

	return &Logger{
		Log: l,
	}
}

func (l *Logger) Info(message string) {
	l.Log.Printf("[%v] INFO: %s\n", curFmtTime("2006-01-02 15:04:05"), message)
}

func (l *Logger) Error(err error, message string) {
	l.Log.Fatalf("[%v] FATAL: %s\n%v", curFmtTime("2006-01-02 15:04:05"), message, err)
}

func curFmtTime(format string) string {
	return time.Now().Format(format)
}
