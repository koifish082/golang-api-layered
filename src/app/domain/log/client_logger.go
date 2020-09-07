package log

import "github.com/koifish082/golang-api-layered/src/app/library/log"

// Logger is used for API client logger
type Logger struct {
}

// ClientLogger returns Logger struct
func ClientLogger() *Logger {
	return &Logger{}
}

// Debugf creates a string formatted log at the debug level.
func (l Logger) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Warnf creates a string formatted log at the debug level.
func (l Logger) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Errorf creates a string formatted log with an error attached.
func (l Logger) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}
