package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Context() {
	var wg sync.WaitGroup

	// Context
	fmt.Println("\nContext:")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go taskWithContext(ctx, &wg)

	// Cancel the context after 500 milliseconds
	time.Sleep(time.Millisecond * 500)
	cancel()

	wg.Wait()
}

func taskWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled. Exiting.")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Millisecond * 100)
		}
	}
}
