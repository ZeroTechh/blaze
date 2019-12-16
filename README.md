[![Go Report Card](https://goreportcard.com/badge/github.com/ZeroTechh/blaze)](https://goreportcard.com/report/github.com/ZeroTechh/blaze) [![GoDoc](https://godoc.org/github.com/ZeroTechh/blaze?status.svg)](https://godoc.org/github.com/ZeroTechh/blaze) [![GoCover](https://gocover.io/_badge/github.com/ZeroTechh/blaze)](https://gocover.io/github.com/ZeroTechh/blaze)

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

## Function Status Logging
As of now, this only supports zap logger
```go

log := zap.NewDevelopement(
        zap.AddCallerSkip(1), 
        /* remember to add 1 call skip to your zap logger
        or else the caller would always be blaze */
    )

func someFunction(someArgument string) {
    

    funcLog := blaze.NewFuncLog(
        "someFunction", // function name
        log, // zap logger
        zap.String("someArgument", someArgument), // you can provide as many zap fields
    )

    defer func(){
        if r:=recover();r!=nil {
            funcLog.Panic(r) // logs the panic message PANIC
            // Works better with panic handler of blaze
        }
    }()

    funcLog.Started() // logs that function started DEBUG

    output, err := someTask()
    if err != nil {
        funcLog.Error(err) // logs the error ERROR
    }

    funcLog.Completed(zap.String("output", output)) 
    /* logs that function completed INFO.
       In all of the function of FuncLog 
       you can enter more zap fields like done here
    */
}


```