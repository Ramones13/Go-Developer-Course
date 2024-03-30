package main

import "fmt"

func main() {

	const (
		powerOfTwo0 = 1 << iota
		powerOfTwo1
		powerOfTwo2
		powerOfTwo3
		powerOfTwo4
	)

	fmt.Println(powerOfTwo0)
	fmt.Println(powerOfTwo1)
	fmt.Println(powerOfTwo2)
	fmt.Println(powerOfTwo3)
	fmt.Println(powerOfTwo4)

}
