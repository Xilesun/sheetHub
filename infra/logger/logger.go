package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
}

// Infof logs a message at the info log level.
func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// Debugf logs a message at the debug log level.
func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// Warnf logs a message at the warn log level.
func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

// Errorf logs a message at the error log level.
func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// Fatalf logs a message at the fatal log level.
func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

// Panicf logs a message at the panic log level.
func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}
