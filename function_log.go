package blaze

import (
	"fmt"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// NewFuncLog will create a new func log
func NewFuncLog(name string, log *zap.Logger, fields ...zap.Field) *FuncLog {
	return &FuncLog{
		name,
		log,
		fields,
	}
}

// FuncLog is used to easily log golang function statuses
type FuncLog struct {
	name   string
	log    *zap.Logger
	fields []zap.Field
}

// getFields This will concat fields with fields provided when struct was initialized
func (funcLog FuncLog) getFields(fields ...zap.Field) []zap.Field {
	newFields := funcLog.fields
	if fields != nil {
		newFields = append(newFields, fields...)
	}
	return newFields
}

// Panic will log the panic at PANIC level
func (funcLog FuncLog) Panic(message interface{}, fields ...zap.Field) {
	fieldsToAdd := append(
		funcLog.getFields(fields...), zap.Any("Panic Message", message))
	funcLog.log.Panic(
		fmt.Sprintf("Function %s Ran Into A Panic", funcLog.name),
		fieldsToAdd...,
	)
}

// ErrorWrap will log the error at ERROR level and wrap it with information
func (funcLog FuncLog) Error(err error, fields ...zap.Field) error {
	msg := fmt.Sprintf("Function %s Ran Into An Error", funcLog.name)
	fieldsToAdd := append(funcLog.getFields(fields...), zap.Error(err))
	funcLog.log.Error(
		msg,
		fieldsToAdd...,
	)
	return errors.Wrap(err, msg)
}

// Completed will that function completed at INFO level
func (funcLog FuncLog) Completed(fields ...zap.Field) {
	funcLog.log.Info(
		fmt.Sprintf("Function %s Executed Successfully", funcLog.name),
		funcLog.getFields(fields...)...,
	)
}

// Started will log that function started at DEBUG level
func (funcLog FuncLog) Started(fields ...zap.Field) {
	funcLog.log.Debug(
		fmt.Sprintf("Function %s Started", funcLog.name),
		funcLog.getFields(fields...)...,
	)
}
