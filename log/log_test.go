package log

import (
	"github.com/sirupsen/logrus"
	"testing"
)

var (
	logger   = NewPackageLogger(logrus.InfoLevel)
	logEntry = NewPackageLoggerEntry(logger, "log_test")
)

func TestLog(t *testing.T) {
	logEntry.ContextLogger().Info("Hello World")
}
