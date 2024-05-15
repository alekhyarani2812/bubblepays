package log

import (
	"context"
	bcontext "github.com/bubblepays/context"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

var logger *logrus.Logger

func init() {
	logger = NewLogger()
}

func NewLogger() *logrus.Logger {
	// Create a new instance of the logger
	logger := logrus.New()
	// Set the output to stdout
	logger.Out = os.Stdout
	// Set the log format
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return logger
}

func getLogFields(ctx context.Context) logrus.Fields {
	transactionId := ""
	if ctx != nil {
		transactionId = bcontext.GetTransactionID(ctx)
	}
	functionName := "unknown"
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	} else {
		file = filepath.Base(file) // Get the base file name
		//get function name
		functionName = runtime.FuncForPC(pc).Name()
	}

	return logrus.Fields{
		"fileName":      file,
		"functionName":  functionName,
		"lineNumber":    line,
		"transactionId": transactionId,
	}
}

func Infot(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(ctx)).Info(message)
}

func Infoft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(ctx)).Infof(format, args...)
}

func Warnt(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.WarnLevel)
	logger.WithFields(getLogFields(ctx)).Warn(message)
}

func Warnft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.WarnLevel)
	logger.WithFields(getLogFields(ctx)).Warnf(format, args...)
}

func Errort(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(ctx)).Error(message)
}

func Errorft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(ctx)).Errorf(format, args...)
}

func Debugt(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.DebugLevel)
	logger.WithFields(getLogFields(ctx)).Debug(message)
}

func Debugft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.DebugLevel)
	logger.WithFields(getLogFields(ctx)).Debugf(format, args...)
}

// without context
func Info(message interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(nil)).Info(message)
}

func Infof(format string, args ...interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	//fmt.Sprintf()
	logger.WithFields(getLogFields(nil)).Infof(format, args...)
}

func Fatal(message interface{}) {
	logger.SetLevel(logrus.FatalLevel)
	logger.WithFields(getLogFields(nil)).Fatal(message)
}

func Fatalf(format string, args ...interface{}) {
	logger.SetLevel(logrus.FatalLevel)
	logger.WithFields(getLogFields(nil)).Fatalf(format, args...)
}

func Error(message interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(nil)).Error(message)
}

func Errorf(format string, args ...interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(nil)).Errorf(format, args...)
}
