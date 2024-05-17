package log

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	bcontext "github.com/bubblepays/context"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = NewLogger()
}

//NewLogger to Create a new instance of the logger
func NewLogger() *logrus.Logger {
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
	transactionID := ""
	if ctx != nil {
		transactionID = bcontext.GetTransactionID(ctx)
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
		"transactionID": transactionID,
	}
}
//Infot for info with transactionID
func Infot(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(ctx)).Info(message)
}
//Infoft for info with arguments and transactionID
func Infoft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(ctx)).Infof(format, args...)
}
//Warnt for Warning with transactionID
func Warnt(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.WarnLevel)
	logger.WithFields(getLogFields(ctx)).Warn(message)
}
//Warnft for Warning with arguments and transactionID
func Warnft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.WarnLevel)
	logger.WithFields(getLogFields(ctx)).Warnf(format, args...)
}
//Errort for Error with transactionID
func Errort(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(ctx)).Error(message)
}
//Errorft for Error with arguments and transactionID
func Errorft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(ctx)).Errorf(format, args...)
}
//Debugt for Debug with transactionID
func Debugt(ctx context.Context, message interface{}) {
	logger.SetLevel(logrus.DebugLevel)
	logger.WithFields(getLogFields(ctx)).Debug(message)
}
//Debugft for Debug with arguments and transactionID
func Debugft(ctx context.Context, format string, args ...interface{}) {
	logger.SetLevel(logrus.DebugLevel)
	logger.WithFields(getLogFields(ctx)).Debugf(format, args...)
}

//Info for info message
func Info(message interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(nil)).Info(message)
}
//Infof for info with arguments
func Infof(format string, args ...interface{}) {
	logger.SetLevel(logrus.InfoLevel)
	logger.WithFields(getLogFields(nil)).Infof(format, args...)
}
//Fatal it logs the message and then calls os.Exit
func Fatal(message interface{}) {
	logger.SetLevel(logrus.FatalLevel)
	logger.WithFields(getLogFields(nil)).Fatal(message)
}
//Fatalf it logs the message and then calls os.Exit
func Fatalf(format string, args ...interface{}) {
	logger.SetLevel(logrus.FatalLevel)
	logger.WithFields(getLogFields(nil)).Fatalf(format, args...)
}
//Error message which returns a string describing the error
func Error(message interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(nil)).Error(message)
}
//Errorf for Error with arguments
func Errorf(format string, args ...interface{}) {
	logger.SetLevel(logrus.ErrorLevel)
	logger.WithFields(getLogFields(nil)).Errorf(format, args...)
}
