package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker1(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		// 执行计算任务
		result := job * 2
		// 将结果发送到结果通道
		results <- result
	}
}

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	jobs := make(chan int, numCPU)
	results := make(chan int, numCPU)

	var wg sync.WaitGroup
	for i := 1; i <= numCPU; i++ {
		wg.Add(1)
		go func(worker1ID int) {
			defer wg.Done()
			worker1(worker1ID, jobs, results)
		}(i)
	}

	// 发送任务到工作通道
	for i := 1; i <= numCPU*2; i++ {
		jobs <- i
	}
	close(jobs)

	// 等待所有 worker1 goroutine 完成任务
	wg.Wait()

	// 关闭结果通道
	close(results)

	// 输出结果
	for result := range results {
		fmt.Println("Result:", result)
	}
}
