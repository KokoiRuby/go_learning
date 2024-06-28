[time](https://pkg.go.dev/time)

时间常量

```go
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

time.Sleep(time.Second * 0.1)      // nok
time.Sleep(time.Millisecond * 100) // ok
```

unix/unixnano timestamp (the number of s elapsed since Jan. 1, 1970 UTC) → cost.