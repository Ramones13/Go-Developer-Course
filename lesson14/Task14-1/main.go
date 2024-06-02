package main

import (
	"fmt"
	"time"
)

func main() {

	go fmt.Println("Привет из дочерней горутины!")

	time.Sleep(time.Millisecond)

}
