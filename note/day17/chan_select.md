select 可以从**多管道取数据** multi-plexing，每一个分支可以处理一个管道，随机选择一个 ready 管道的分支并执行。

:warning: 如果包含了 default 分支，不会等待管道分支，而是直接快速进入 default 分支。

```go
select {
  case msg1 := <-ch1:
    fmt.Println("Received", msg1)
  case msg2 := <-ch2:
    fmt.Println("Received", msg2)
  // default:
  //   fmt.Println("No activity")
}
```

实际开发中，不好确定什么时候关闭该管道，不能显式使用 close 进行关闭。使用 select 可以不需要关闭。

```go
select {
  case msg1, ok := <-ch1:
    if ok {
      fmt.Println("Received", msg1)
    } else {
      fmt.Println("Channel closed!")
    }
  case msg2 := <-ch2:
    fmt.Println("Received", msg2)
}
```

超时处理：避免无限阻塞。

```go
// ch1 is on-wait, timer starts
// if ch1 becomes ready right then before timeout, noted timer still running in background, and when it expires & send to channel even we're not waiting on it.
// timer chan will stay in mem until its timeout.
select {
  case msg1 := <-ch1:
    fmt.Println("Received", msg1)
  case <-time.After(2 * time.Second):
    fmt.Println("Timeout")
}
```

```go
// perferable
timer := time.NewTimer(2 * time.Second)

select {
  case msg1 := <-ch1:
    fmt.Println("Received", msg1)
	// free timer immediately once ch1 is ready
    timer.Stop()
  case <-timer.C:
    fmt.Println("Timeout")
}

fmt.Println("After select statement")
```

++ for：循环检查多管道；可能存在的问题，如果其中一个管道关闭了，接收会默认 0，会一直进入该分支

++ label：to break/exit the loop instead of exiting select case.

```go
OuterLoop:
  for {
    select {
      case msg1, ok := <-ch1:
        if !ok {
          fmt.Println("ch1 is closed. Exiting loop.")
          break OuterLoop // This exits the loop named OuterLoop
        }
        fmt.Println("Received from ch1:", msg1)
      case msg2, ok := <-ch2:
        if !ok {
          fmt.Println("ch2 is closed. Exiting loop.")
          break OuterLoop // Same here
        }
        fmt.Println("Received from ch2:", msg2)
      case <-time.After(1 * time.Second):
        fmt.Println("Waiting...")
      }
  }
```

**++ nil: Make chan nil to avoid enter case.**

```go
OuterLoop:
  for {
    select {
      case msg1, ok := <-ch1:
        if !ok {
          fmt.Println("ch1 is closed. Switching off.")
          ch1 = nil
        }
        fmt.Println("Received from ch1:", msg1)
      case msg2, ok := <-ch2:
        if !ok {
          fmt.Println("ch2 is closed. Exiting loop.")
          ch2 = nil
        }
        fmt.Println("Received from ch2:", msg2)
      // exit loop if all chan are closed
      case <-time.After(1 * time.Second):
        fmt.Println("Waiting...")
        
        if (ch1 == nil) && (ch2 == nil) {
          fmt.Println("Both channels are closed. Exiting loop.")
          break OuterLoop
        }
    }
  }
```

非阻塞发送

```go
select {
  case ch1 <- 1:
    fmt.Println("Sent 1 to ch1")
  case ch2 <- 2:
    fmt.Println("Sent 2 to ch2")
  default: // <---- 
    fmt.Println("Neither channel was ready for sending")
}
```

