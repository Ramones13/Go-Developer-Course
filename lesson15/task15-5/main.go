package main

import (
	"fmt"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type meteoCaller struct {
	mu                 sync.RWMutex
	currentTemperature string
}

func (m *meteoCaller) ReadTemp() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.currentTemperature
}

func (m *meteoCaller) ChangeTemp(v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentTemperature = v
}
func main() {

	m := meteoCaller{}

	writeWaitGrout := sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		writeWaitGrout.Add(1)
		go func() {
			defer writeWaitGrout.Done()
			m.ChangeTemp(fmt.Sprintf("%v Celsius", 23.3+float32(i)))
			time.Sleep(100 * time.Millisecond)
		}()
	}

	readWaitGroup := sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		readWaitGroup.Add(1)
		go func() {
			defer readWaitGroup.Done()
			fmt.Printf("temp from %v goroutine - %v\n", i, m.ReadTemp())
			time.Sleep(200 * time.Millisecond)
		}()
	}

	writeWaitGrout.Wait()
	readWaitGroup.Wait()

}
