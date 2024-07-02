> **不要通过共享内存来通信，要通过通信来共享内存**
>
> Don’t communicate by sharing memory; share memory by communicating.

:confused: **不同协程间如何通信？**

- 全局变量加锁
- **Channel** ≈ pipe-like 以一种安全的方式，实现协程间的同步。



 管道 Channel 属于**引用类型**。同 slice/map，需要通过 make 进行初始化。

- Unbufferred：无缓冲，即没有指定 cap；发送和接收是同步的。
  - 发送会阻塞管道，直到接收。
  - 接收会阻塞管道，直到发送。
  - UC：强制同步协作
- Buffered：缓冲，指定 cap；管道未满不会阻塞。
  - 缓冲满则阻塞发送
  - 缓冲空会阻塞接收
  - UC：（不同速率）生产消费

```go
newChan := make(chan int[, cap])

// len <= cap
len(newChan)
cap(newChan)
```

```go
// send
intChan := make(chan int, 3)
intChan <- 10

// recv
n := <-intChan
```

可以创建一个空接口管道，可以存放任何类型的变量

```go
allChan := make(chan inerface{},10)
```

关闭

```go
intChan := make(chan int, 3)
intChan <- 1
intChan <- 2
close(intChan) // disable push only
a := <-intChan // can still pop
```

遍历：只能使用 for-range 遍历，因为 chan 的长度是变化的。

```go
intChan := make(chan int, 10)
for i := 0; i < 10; i++ {
	intChan <- i
}
close(intChan) // must close before iter

for v := range intChan {
	fmt.Println(v)
}
```

https://blog.devtrovert.com/p/go-channels-explained-more-than-just