package blaze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func handler(msg interface{}, data interface{}) {
	asserter := data.(*assert.Assertions)
	asserter.Equal("panic", msg)
}

func TestPanicHandler(t *testing.T) {
	assert := assert.New(t)
	panicHandler := NewPanicHandler(handler)

	defer panicHandler.Check(assert)
	panic("panic")
}
