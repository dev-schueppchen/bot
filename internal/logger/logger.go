package logger

import (
	"github.com/op/go-logging"
)

const mainLoggerName = "main"

var mainLogger = logging.MustGetLogger(mainLoggerName)

// Set the log level of the logger
func SetLogLevel(lvl int) {
	logging.SetLevel(logging.Level(lvl), mainLoggerName)
}

// Setup sets configuration for the
// loggers format and level
func Setup(format string, lvl int) {
	formatter := logging.MustStringFormatter(format)
	logging.SetFormatter(formatter)
	SetLogLevel(lvl)
}

// GetLogger returns the main logger
// instance
func GetLogger() *logging.Logger {
	return mainLogger
}
