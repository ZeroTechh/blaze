package blaze

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

func TestFuncLog(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	log, _ := zap.NewDevelopment(
		zap.AddCallerSkip(1),
	)
	funcLog := NewFuncLog(
		"test",
		log,
		zap.String("test", "test"),
	)

	funcLog.Started()
	funcLog.Error(errors.New("Error"))
	funcLog.Completed(zap.String("Completed", "true"))
	funcLog.Panic("paniced")
}
