利用**多核运算**是一个典型的**信号量模式**。

指定核心数，即信号量，并创建工作和结果管道。

```go
numCPU := runtime.NumCPU()
runtime.GOMAXPROCS(numCPU)

jobs := make(chan int, numCPU)
results := make(chan int, numCPU)
```

通过 sync.WaitGroup 控制主进程等待所有协程结束。每次启动协程++，结束--。

启动对应核心数的协程并执行 worker 函数。如果 job 为空，那么 worker 会阻塞。

```go
var wg sync.WaitGroup
for i := 1; i <= numCPU; i++ {
	wg.Add(1)
	go func(worker1ID int) {
		defer wg.Done()
		worker(workerID, jobs, results)
	}(i)
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		result := job * 2
		results <- result
	}
}
```

发送任务到工作通道。

```go
for i := 1; i <= numCPU*2; i++ {
	jobs <- i
}
close(jobs)
```

等待，关闭结果通道，并输出。

```go
wg.Wait()

close(results)

for result := range results {
	fmt.Println("Result:", result)
}
```

流水线模式

```go
func ParallelProcessData (in <- chan *Data, out <- chan *Data) {

    // make channels:
    preOut := make(chan *Data, 100)
    stepAOut := make(chan *Data, 100)
    stepBOut := make(chan *Data, 100)
    stepCOut := make(chan *Data, 100)

    // start parallel computations:
    go PreprocessData(in, preOut)
    go ProcessStepA(preOut, stepAOut)
    go ProcessStepB(stepAOut, stepBOut)
    go ProcessStepC(stepBOut, stepCOut)
    go PostProcessData(stepCOut, out)

}
```



