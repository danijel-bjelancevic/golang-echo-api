package logger

import (
	"github.com/sirupsen/logrus"
)

// New will return new configured logger instance
func New(debug bool, apiName string, graylogHost string) (logrus.FieldLogger, error) {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	if debug {
		l.SetLevel(logrus.DebugLevel)
	}
	logger := l.WithField("app", apiName)

	return logger, nil
}
