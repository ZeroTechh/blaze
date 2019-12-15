package blaze

import "reflect"

// NewErrorHandler is used to create a new error handler
func NewErrorHandler(
	defaultHandler func(error, ...interface{}) interface{}) *ErrorHandler {

	handler := &ErrorHandler{
		defaultHandler: defaultHandler,
	}
	handler.Init()
	return handler
}

// ErrorHandler is used for error handling
type ErrorHandler struct {
	handlers       map[reflect.Type]func(error, ...interface{}) interface{}
	defaultHandler func(error, ...interface{}) interface{}
}

// Init initializes
func (errorHandler *ErrorHandler) Init() {
	errorHandler.handlers = map[reflect.Type](func(error, ...interface{}) interface{}){}
}

// AddHandler adds a function handler. for the err type of err provided
func (errorHandler *ErrorHandler) AddHandler(
	handlerFunc func(error, ...interface{}) interface{},
	err error) {

	errType := reflect.TypeOf(err)
	errorHandler.handlers[errType] = handlerFunc
}

/* Check will check if err is nil,
   if not then it will check for handler of that err type and run it
   else it will run default err handler.
   The data will be passed on to the err handler */
func (errorHandler ErrorHandler) Check(err error, data ...interface{}) interface{} {

	if err == nil {
		return nil
	}

	errType := reflect.TypeOf(err)
	handler, valid := errorHandler.handlers[errType]

	if !valid {
		handler = errorHandler.defaultHandler
	}

	return handler(err, data...)
}
