`time.Ticker` 结构体自带一个只允许接收的管道。

协程执行期间周期性执行任务。

```go
ticker := time.NewTicker(updateInterval)
defer ticker.Stop()
...
select {
case u:= <-ch1:
    ...
case v:= <-ch2:
    ...
case <-ticker.C:
    // TODO
default: 
    ...
}
```

周期性往通道发送时间时间。

```go
duration := time.Second * 1
ch := time.Tick(duration)
fmt.Println("Starting ticker...")
for now := range ch {
	fmt.Println("Tick at", now)
	// TODO
}
```

`time.Timer` 类 `time.Ticker` 但只发送一次时间。

简单超时，避免阻塞。

```go
timeout := make(chan bool, 1)
go func() {
        time.Sleep(time.Second * 1)
        timeout <- true
}()

select {
    case <-ch:
        // TODO
    case <-timeout:
        // TODO
        break
}
```

```go
// perferable
timer := time.NewTimer(2 * time.Second)
defer timer.Stop()

select {
  case msg1 := <-ch1:
    fmt.Println("Received", msg1)
  case <-timer.C:
    fmt.Println("Timeout")
}

fmt.Println("After select statement")
```

从多个 DB 数据同时读取数据，只拿一个。

```go
func Query(conns []conn, query string) Result {
    ch := make(chan Result, 1)
    for _, conn := range conns {
        go func(c Conn) {
            select {
            case ch <- c.DoQuery(query):  // first return 
            default:                      // follow-up block
            }
        }(conn)
    }
    return <- ch
}
```

