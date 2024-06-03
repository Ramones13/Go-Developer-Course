package main

import (
	"fmt"
	"sync"
)

func main() {

	var counter int

	counterWaitGroup := sync.WaitGroup{}
	var counterMutex sync.Mutex
	for i := 1; i <= 5; i++ {
		counterWaitGroup.Add(1)
		go func() {
			defer counterWaitGroup.Done()
			counterMutex.Lock()
			defer counterMutex.Unlock()
			fmt.Printf("goroutine %v has started\n", i)
			counter++
		}()
	}

	counterWaitGroup.Wait()

	counterMutex.Lock()
	defer counterMutex.Unlock()
	fmt.Printf("result: %v", counter)

}
