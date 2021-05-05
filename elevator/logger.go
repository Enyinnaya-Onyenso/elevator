package elevator

import (
	"bytes"
	"fmt"
	"log"
)

type Logger struct {
	*log.Logger
	buf bytes.Buffer
}

func (l *Logger) init() error {
	l.Logger = log.New(&l.buf, "Log: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile) // Logger component initialization
	l.Println("Logger Online")
	fmt.Println(&l.buf)
	l.buf.Reset()
	return nil
}

func newLogger() *Logger {
	// Logger to be used in system.go
	return &Logger{}
}
