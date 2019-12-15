package blaze

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultHandlerOccured bool
	customHandlerOccured  bool
)

type customError struct{}

func (ce customError) Error() string { return "custom error" }

func defaultHandler(err error, data ...interface{}) interface{} {
	defaultHandlerOccured = true
	return data
}

func customHandler(err error, data ...interface{}) interface{} {
	customHandlerOccured = true
	return err
}

func TestErrHandler(t *testing.T) {
	errHandler := NewErrorHandler(defaultHandler)
	assert := assert.New(t)

	errHandler.AddHandler(customHandler, customError{})

	// Testing that nil is returned when err is nil
	assert.Nil(errHandler.Check(nil))

	// Testing that error handler uses defaultHandler for generic error
	err := errors.New("generic error")
	data := errHandler.Check(err, nil)
	assert.True(defaultHandlerOccured)

	// data should be nil as no data was passed when err was checked
	assert.Nil(data.([]interface{})[0])

	// Testing that data is successfully passed on to the handler
	err = errors.New("generic error")
	data = errHandler.Check(err, "test")

	// data should not be nil as "test" was passed when err was checked
	assert.Equal("test", data.([]interface{})[0])

	// Testing that error handler uses customHandler for generic customError
	err = customError{}
	data = errHandler.Check(err, nil)
	assert.True(customHandlerOccured)
	assert.Equal(err, data) // Testing id err was successfully passed on
}
