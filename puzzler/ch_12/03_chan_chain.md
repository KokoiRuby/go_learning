使用 goroutine 链式调用（Goroutine Chaining）来组合多个并发操作。

通过将多个 goroutine 连接在一起，可以实现并发执行的任务流水线或任务链。

每一个阶段函数，都接收一个通道，返回一个通道。

```go
input := make(chan int)
for i := 1; i <= 5; i++ {
	input <- i
	time.Sleep(500 * time.Millisecond)
}

output := stage2(stage1(input))
for result := range output {
	fmt.Println("Output:", result)
}
```

```go
func stage1(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			// Stage 1 processing
			result := num * 2
			out <- result
		}
	}()
	return out
}

func stage2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			// Stage 2 processing
			result := num + 1
			out <- result
		}
	}()
	return out
}
```