### Generator

生成器是指当**被调用时返回一个序列中下一个值**的函数。也称惰性求值，而不是返回整个序列。

创建一个无缓冲的通道，启动一个协程不断往通道发送，∵ 无缓冲所以后续会被阻塞。每次调用时从通道取出。

```go
func integers() chan int {
    yield := make(chan int)
    count := 0
    go func() {
        for {
            yield <- count
            count++
        }
    }()
    return yield
}

func generateInteger() int {
    return <-resume
}

func main() {
    resume = integers()
    fmt.Println(generateInteger())
    fmt.Println(generateInteger())
    fmt.Println(generateInteger())
}
```

同时**函数闭包**也可以实现简单的惰性求值。

```go
func lazyAdd(a, b int) func() int {
	return func() int {
		return a + b
	}
}

addFunc := lazyAdd(1, 2)
res := addFunc()  // cal only when called
```

### Future

在使用某一个值之前需要先对其进行计算，通常在另一个处理器上求值，使用时就已计算完毕。类生成器，但需要返回一个值。

求 a/b 逆矩阵的时候可以并行执行，那么就可以分别放入 goroutine 中。在求值完之前通道是阻塞的。

```go
func InverseFuture(a Matrix) {
    future := make(chan Matrix)
    go func() {
        future <- Inverse(a)
    }()
    return future
}

func InverseProduct(a Matrix, b Matrix) {
    a_inv_future := InverseFuture(a)
    b_inv_future := InverseFuture(b)
    a_inv := <-a_inv_future
    b_inv := <-b_inv_future
    return Product(a_inv, b_inv)
}
```

