goroutine panic recover 而不影响其他协程。

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)   
    }
}

func safelyDo(work *Work) {
    defer func() {
        // logging if necessary
        if err := recover(); err != nil {
            log.Printf("Work failed with %s in %v", err, work)
        }
    }()
    do(work)
}
```

