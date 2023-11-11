package logger

import (
	"fmt"
	"log"
)

var logger = log.Default()

func Error(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Error]: %s\n", msg)
}

func Info(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}
