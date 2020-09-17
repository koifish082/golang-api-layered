package log

import (
	"log"
	"time"

	"github.com/sirupsen/logrus"
)

var logrusLogger *logrus.Entry

// Config is struct for log setting values
type Config struct {
	ServiceName    string
	ServiceDetails interface{}
	LogLevel       logrus.Level
}

type logWriter struct {
	*logrus.Entry
}

func init() {
	logrusLogger = logrus.NewEntry(logrus.New())
	log.SetOutput(&logWriter{logrusLogger})
	log.SetFlags(0)
}

func (l *logWriter) Write(p []byte) (n int, err error) {
	if len(p) > 0 {
		l.Infof("%s", p[:len(p)-1])
	}
	return len(p), nil
}

// InitLogger init log by Config
func InitLogger(c *Config) {
	fields := logrus.Fields{
		"service":         c.ServiceName,
		"service_details": c.ServiceDetails,
	}
	logrusLogger = logrus.WithFields(fields)
	logrusLogger.Logger.SetLevel(c.LogLevel)
	logrusLogger.Time = time.Now().UTC()
}

// Debugf creates a string formatted log at the debug level.
func Debugf(format string, args ...interface{}) {
	logrusLogger.Debugf(format, args...)
}

// Infof creates a string formatted log at the debug level.
func Infof(format string, args ...interface{}) {
	logrusLogger.Infof(format, args...)
}

// Printf creates a string formatted log at the debug level.
func Printf(format string, args ...interface{}) {
	logrusLogger.Printf(format, args...)
}

// Warnf creates a string formatted log at the debug level.
func Warnf(format string, args ...interface{}) {
	logrusLogger.Warnf(format, args...)
}

// Errorf creates a string formatted log with an error attached.
func Errorf(format string, args ...interface{}) {
	logrusLogger.Errorf(format, args...)
}

// Fatalf calls Fatalf on the package level logrusLogger.
func Fatalf(format string, args ...interface{}) {
	logrusLogger.Fatalf(format, args...)
}

// Panicf calls Panicf on the package level logrusLogger.
func Panicf(format string, args ...interface{}) {
	logrusLogger.Panicf(format, args...)
}

// Debug calls debug on the package level logrusLogger.
func Debug(args ...interface{}) {
	logrusLogger.Debug(args...)
}

// Info calls Info on the package level logrusLogger.
func Info(args ...interface{}) {
	logrusLogger.Info(args...)
}

// Print calls Print on the package level logrusLogger.
func Print(args ...interface{}) {
	logrusLogger.Print(args...)
}

// Warn calls Warn on the package level logrusLogger.
func Warn(args ...interface{}) {
	logrusLogger.Warn(args...)
}

// Error calls Error on the package level logrusLogger.
func Error(args ...interface{}) {
	logrusLogger.Error(args...)
}

// Fatal calls Fatal on the package level logrusLogger.
func Fatal(args ...interface{}) {
	logrusLogger.Fatal(args...)
}

// Panic calls Panic on the package level logrusLogger.
func Panic(args ...interface{}) {
	logrusLogger.Panic(args...)
}

// Debugln calls Debugln on the package level logrusLogger.
func Debugln(args ...interface{}) {
	logrusLogger.Debugln(args...)
}

// Infoln calls Infoln on the package level logrusLogger.
func Infoln(args ...interface{}) {
	logrusLogger.Infoln(args...)
}

// Println calls Println on the package level logrusLogger.
func Println(args ...interface{}) {
	logrusLogger.Println(args...)
}

// Warnln calls Warnln on the package level logrusLogger.
func Warnln(args ...interface{}) {
	logrusLogger.Warnln(args...)
}

// Errorln calls Errorln on the package level logrusLogger.
func Errorln(args ...interface{}) {
	logrusLogger.Errorln(args...)
}

// Fatalln calls Fatalln on the package level logrusLogger.
func Fatalln(args ...interface{}) {
	logrusLogger.Fatalln(args...)
}

// Panicln calls Panicln on the package level logrusLogger.
func Panicln(args ...interface{}) {
	logrusLogger.Panicln(args...)
}
