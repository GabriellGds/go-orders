package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	info    *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(l string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, l, log.Ldate|log.Ltime)

	return &Logger{
		
		info:    log.New(writer, "INFO: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}