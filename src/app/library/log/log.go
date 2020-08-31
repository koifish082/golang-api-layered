package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

type Config struct {
	ServiceName    string
	ServiceDetails interface{}
	LogLevel       logrus.Level
}

func InitLogger(c *Config) {
	fields := logrus.Fields{
		"service":         c.ServiceName,
		"service_details": c.ServiceDetails,
	}
	Logger = logrus.WithFields(fields)
	Logger.Logger.SetLevel(c.LogLevel)
	Logger.Time = time.Now().UTC()
}

// Debugf creates a string formatted log at the debug level.
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Infof creates a string formatted log at the debug level.
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Printf creates a string formatted log at the debug level.
func Printf(format string, args ...interface{}) {
	Logger.Printf(format, args...)
}

// Warnf creates a string formatted log at the debug level.
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Errorf creates a string formatted log with an error attached.
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Fatalf calls Fatalf on the package level logger.
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

// Panicf calls Panicf on the package level logger.
func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}

// Debug calls debug on the package level logger.
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Info calls Info on the package level logger.
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Print calls Print on the package level logger.
func Print(args ...interface{}) {
	Logger.Print(args...)
}

// Warn calls Warn on the package level logger.
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Error calls Error on the package level logger.
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Fatal calls Fatal on the package level logger.
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Panic calls Panic on the package level logger.
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

// Debugln calls Debugln on the package level logger.
func Debugln(args ...interface{}) {
	Logger.Debugln(args...)
}

// Infoln calls Infoln on the package level logger.
func Infoln(args ...interface{}) {
	Logger.Infoln(args...)
}

// Println calls Println on the package level logger.
func Println(args ...interface{}) {
	Logger.Println(args...)
}

// Warnln calls Warnln on the package level logger.
func Warnln(args ...interface{}) {
	Logger.Warnln(args...)
}

// Errorln calls Errorln on the package level logger.
func Errorln(args ...interface{}) {
	Logger.Errorln(args...)
}

// Fatalln calls Fatalln on the package level logger.
func Fatalln(args ...interface{}) {
	Logger.Fatalln(args...)
}

// Panicln calls Panicln on the package level logger.
func Panicln(args ...interface{}) {
	Logger.Panicln(args...)
}
