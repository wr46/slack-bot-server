package logger

import (
	"io"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	IsDebug bool
)

// Setup used for standardize definition of log messages
func Setup(debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Debug = log.New(debugHandle, "DEBUG: ", log.Ldate|log.Ltime)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime)
	IsDebug = false
}

// SetDebug set Debug mode
func SetDebug(debug bool) {
	IsDebug = debug
}

// Log only log debug messages when in Debug mode
func Log(logger *log.Logger, message string) {
	if logger == Debug && !IsDebug {
		return
	}
	logger.Println(message)

	if logger == Error {
		os.Exit(500)
	}
}
