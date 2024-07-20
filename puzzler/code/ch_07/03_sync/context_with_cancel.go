package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Task running")
		}
	}
}

func main() {
	cancelCtx, cancel := context.WithCancel(context.Background())

	go longRunningTask(cancelCtx)

	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
}
