package blaze

import (
	"fmt"

	"go.uber.org/zap"
)

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

/* Panic will log an Panic of
the panic message and that the function paniced */
func (funcLog FuncLog) Panic(message interface{}, fields ...zap.Field) {
	fieldsToAdd := append(
		funcLog.getFields(fields...), zap.Any("Panic Message", message))
	funcLog.log.Panic(
		fmt.Sprintf("Function %s Ran Into A Panic", funcLog.name),
		fieldsToAdd...,
	)
}

// Error will log the error and that the function ran into error
func (funcLog FuncLog) Error(err error, fields ...zap.Field) {
	fieldsToAdd := append(funcLog.getFields(fields...), zap.Error(err))
	funcLog.log.Error(
		fmt.Sprintf("Function %s Ran Into An Error", funcLog.name),
		fieldsToAdd...,
	)
}

// Completed will log an info that the function ran successfully
func (funcLog FuncLog) Completed(fields ...zap.Field) {
	funcLog.log.Info(
		fmt.Sprintf("Function %s Executed Successfully", funcLog.name),
		funcLog.getFields(fields...)...,
	)
}

// Started will log an debug that the function started successfully
func (funcLog FuncLog) Started(fields ...zap.Field) {
	funcLog.log.Debug(
		fmt.Sprintf("Function %s Started", funcLog.name),
		funcLog.getFields(fields...)...,
	)
}
