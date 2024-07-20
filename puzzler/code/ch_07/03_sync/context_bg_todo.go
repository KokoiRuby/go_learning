package main

import (
	"context"
	"fmt"
	"time"
)

func exampleWithBackgroundContext() {
	ctx := context.Background()

	go func() {
		time.Sleep(2 * time.Second)

		if ctx.Err() != nil {
			fmt.Println("Operation cancelled")
			return
		}

		fmt.Println("Operation completed")
	}()

	time.Sleep(1 * time.Second)
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	cancel()
	
	time.Sleep(3 * time.Second)
}

func exampleWithTODOContext() {
	ctx := context.TODO()

	go func() {
		time.Sleep(2 * time.Second)

		if ctx.Err() != nil {
			fmt.Println("Operation cancelled")
			return
		}

		fmt.Println("Operation completed")
	}()

	time.Sleep(1 * time.Second)
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	cancel()

	time.Sleep(3 * time.Second)
}

func main() {
	fmt.Println("Example with Background Context:")
	exampleWithBackgroundContext()

	fmt.Println("\nExample with TODO Context:")
	exampleWithTODOContext()
}
