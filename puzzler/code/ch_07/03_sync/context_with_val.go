package main

import (
	"context"
	"fmt"
)

func processRequest(ctx context.Context) {
	// Retrieve the value associated with the key from the context
	if value := ctx.Value("requestID"); value != nil {
		requestID, ok := value.(string)
		if ok {
			fmt.Println("Processing request with ID:", requestID)
		}
	}
}

func main() {
	// Create a new context with a key-value pair
	ctx := context.WithValue(context.Background(), "requestID", "12345")

	// Pass the context to the function
	processRequest(ctx)
}
