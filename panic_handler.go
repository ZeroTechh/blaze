package blaze

// NewPanicHandler is used to create a new panic handler
func NewPanicHandler(handler func(interface{}, ...interface{})) *PanicHandler {
	return &PanicHandler{
		handler: handler,
	}
}

// PanicHandler will be used to handle panics
type PanicHandler struct {
	handler func(interface{}, ...interface{})
}

// Check will check for panic and execute handler with data
func (panicHandler PanicHandler) Check(data interface{}) {
	if r := recover(); r != nil {
		panicHandler.handler(r, data)
	}
}
