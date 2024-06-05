package main

import "fmt"

func main() {

	ch := make(chan string, 4)
	ch <- "Привет"
	ch <- "буферизованный канал!"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
