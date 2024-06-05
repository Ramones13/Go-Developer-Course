package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	stop := make(chan struct{}, 2)

	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					fmt.Printf("стоп горутина: %v\n", i)
					return
				default:
					fmt.Printf("сложные вычисления горутины: %v\n", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("ой, всё!")
	close(stop)
	// stop <- struct{}{}

	wg.Wait()

}
