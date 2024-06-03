package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int32

	counterWaitGroup := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		counterWaitGroup.Add(1)
		go func() {
			defer counterWaitGroup.Done()
			fmt.Printf("goroutine %v has started\n", i)
			atomic.AddInt32(&counter, 1)
		}()
	}

	counterWaitGroup.Wait()

	fmt.Printf("result: %v", counter)

}
