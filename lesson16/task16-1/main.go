package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine %v stopped by contex's cancel\n", i)
					return
				default:
					fmt.Printf("goroutine %v is working\n", i)
				}
			}
		}()
	}

	// поскольку printf это системный вызов мы видим залипание на одной из горутин,
	// но увеличивая sleep можно увидеть что все горутины выполняются параллельно
	time.Sleep(time.Millisecond)

	cancel()

	wg.Wait()

}
