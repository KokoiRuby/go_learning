package main

import (
	"context"
	"fmt"
	"time"
)

func performTask(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("task deadline:", deadline)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("cancel task")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("processing...")
		}
	}
}

func main() {
	parentCtx := context.Background()
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	defer cancel()

	go performTask(ctx)

	// 等待一段时间后取消任务
	time.Sleep(3 * time.Second)
	cancel()

	// 等待任务完成
	time.Sleep(2 * time.Second)
}
