package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log Logger

func InitLogs(rootDir ...string) (Logger, error) {

	//basePath := "."
	if len(rootDir) != 0 {
		//	basePath = rootDir[0]
	}

	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})
	logrusLogger.SetReportCaller(true)
	if os.Getenv("APP_DEBUG") == "true" {
		logrusLogger.SetLevel(logrus.TraceLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	log := NewDefaultLogger(logrusLogger)

	return log, nil

}

type DefaultLogger struct {
	logger *logrus.Logger
}

func NewDefaultLogger(logger *logrus.Logger) *DefaultLogger {
	return &DefaultLogger{logger: logger}
}

type Logger interface {
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}

func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

func (l *DefaultLogger) Info(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

func (l *DefaultLogger) Error(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

func (l *DefaultLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}
