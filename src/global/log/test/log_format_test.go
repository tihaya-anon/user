package test

import (
	"MVC_DI/global/log"
	"testing"
	"time"
)

func Test_LogFormat(t *testing.T) {
	interval := 5
	frequency := 3
	logger := log.GetLogger(time.Duration(interval) * time.Second)
	logger.Error("This is an error message.")
	for range 3 {
		for range frequency {
			logger.Debug("This is a debug message.")
			logger.Info("This is an info message.")
			logger.Warn("This is a warn message.")
			time.Sleep(time.Duration(interval/frequency) * time.Second)
		}
	}
}
