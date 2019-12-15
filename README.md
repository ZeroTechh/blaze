# blaze
Making golang easier

## Error Handling
```go
errHandler := blaze.NewErrorHandler()

errHandler.AddHandler(func(err error, data ...interface{}) interface{} {
    fmt.Println("Generic Error Occured")
    return SomeData
})

errHandler.AddHandler(func(err error, data ...interface{}) interface{} {
    fmt.Println("Error Of Type CustomErrorType Occured")
    return SomeData
}, CustomErrorType{})

SomeData := errHandler.Check(err, SomeRandomData1, SomeRandomData2)
// Data passed after the err will be passed to handlers

```

## Panic Handling
```go
func handler(msg interface{}, data ...interface{}) {
    fmt.Println("Panic Occured")
    // msg is the panic message
    // data is data passed when panic was checked
}


panicHandler := blaze.NewPanicHandler(handler)

defer panicHandler.Check("some data")
// "some data" will be passed on to the handler function

panic("some reason")
```