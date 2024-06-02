package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	go func() {
		ch <- "Привет, строковый канал!"
	}()

	fmt.Println(<-ch)

}
