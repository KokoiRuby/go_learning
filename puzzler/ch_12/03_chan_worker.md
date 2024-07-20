一个 worker 处理一项任务，任务可以封装成一个 struct。

- Lock UC
  - 访问共享数据结构中的缓存信息
  - 保存应用程序上下文和状态信息数据
- Channel UC
  - 与异步操作的结果进行交互
  - 分发任务
  - 传递数据所有权

### Lock

任务池共享内存，加锁避免竞争。worker 之前开始结束需要加锁和解锁。**效率低**。

```go
type Task struct {
    //
}

type Pool struct {
    Mu      sync.Mutex
    Tasks   []Task
}
```

```go
func Worker(pool *Pool) {
    for {
        pool.Mu.lock()
        
        task := pool.Task[0]
        pool.Tasks = pool.Task[1:] 
        
        pool.Mu.Unlock()
        process(task)
    }
}
```

### Channel

> **Use an in- and out-channel instead of locking**

双管道：一个通道负责接收任务，一个通道负责处理。

主线程负责发送任务到 pending 通道，同时开启多个协程处理，从 pending 管道中取出处理或放入 done 管道，最终由主线程消费。

```go
func main() {
    pending, done := make(chan *Task), make(chan *Task)
    go sendWork(pending)       
    for i := 0; i < N; i++ {   
        go Worker(pending, done)
    }
    consumeWork(done)          
}

func Worker(in, out chan *Task) {
    for {
        t := <-in
        process(t)
        out <- t
    }
}
```

