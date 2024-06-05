package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))

	defer cancel()

	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine %v stopped by %v\n", i, ctx.Err())
					return
				default:
					fmt.Printf("goroutine %v is working\n", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	wg.Wait()

}
