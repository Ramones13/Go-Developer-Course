package main

import (
	"fmt"
	"sync"
)

var startOnce sync.Once

func main() {

	myWaitGroup := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		myWaitGroup.Add(1)
		go func() {
			defer myWaitGroup.Done()
			startOnce.Do(start)
			fmt.Printf("goroutine %v has started\n", i)
		}()
	}

	myWaitGroup.Wait()

}

func start() {
	fmt.Println("Task 15-4 has started")
}
