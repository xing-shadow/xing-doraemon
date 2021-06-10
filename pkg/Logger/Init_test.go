package Logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestInitA(t *testing.T) {
	if err := Init(SetLevel("info"), SetDevelopment(true), SetLogFileDir("./")); err != nil {
		t.Fatal(err)
	}
	logger.log.Debug("test", zap.String("A", "a"))
	logger.log.Info("test1", zap.String("A", "a"))
	logger.log.Warn("test2", zap.String("A", "a"))
	logger.log.Error("test3", zap.String("A", "a"))
}

func TestB(t *testing.T) {
	if err := Init(SetLevel("info"), SetDevelopment(true), SetLogFileDir("./")); err != nil {
		t.Fatal(err)
	}
	logger.log.Debug("test", zap.String("A", "a"))
	logger.log.Info("test1", zap.String("A", "a"))
	logger.log.Warn("test2", zap.String("A", "a"))
	logger.log.Error("test3", zap.String("A", "a"))
}

func TestMaxBackupsAndMaxSize(t *testing.T) {
	if err := Init(SetLevel("info"), SetDevelopment(false), SetLogFileDir("./"), SetMaxSize(1), SetMaxBackups(2)); err != nil {
		t.Fatal(err)
	}
	for {
		logger.log.Info("testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest")
	}
}
