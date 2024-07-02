`sync.WaitGroup` 监听协程何时执行完毕，执行完毕后通知主线程开始。确保一组并发任务结束后再继续执行。

本身线程安全，可以被多个 Goroutines 共享。

1. **Add(delta int)**：增加或减少等待计数器的值。通常在**启动新的 Goroutine 之前调用**。
   - 不能为负数，否则 panic。
2. **Done()**：减少等待计数器的值，表示一个 Goroutine 完成了工作。通常在 Goroutine **结束时调用**。
3. **Wait()**：阻塞，直到等待计数器的值变为零。通常**在主 Goroutine 中调用**，等待所有其他 Goroutine 完成。

```go
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting...")
	time.Sleep(time.Second * 1)
	fmt.Println("Done.")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// continue until counter reach 0
    wg.Wait()
	fmt.Println("All workers done.")
}
```

