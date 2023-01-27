// Logger package providing support for multiple log levels.  These include
// Verbose, Info, Error, and Warning log levels.
package logging

import (
	"log"
	"os"
)

// Logger flags
const (
	flags = log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
)

// info logger instance
var info *log.Logger

// error logger instance
var error *log.Logger

// warning logger instance
var warning *log.Logger

// verbose lgger instance
var verbose *log.Logger

// Info logger
func Info(v ...any) {
	info.Println(v...)
}

// Error logger
func Error(v ...any) {
	error.Println(v...)
}

// Warning logger
func Warning(v ...any) {
	warning.Println(v...)
}

// Verbose logger
func Verbose(v ...any) {
	verbose.Println(v...)
}

// Package initialization creates the set of loggers
func init() {
	info = log.New(os.Stderr, "INFO: ", flags)
	error = log.New(os.Stderr, "ERROR: ", flags)
	warning = log.New(os.Stderr, "WARNING: ", flags)
	verbose = log.New(os.Stderr, "VERBOSE: ", flags)
}
