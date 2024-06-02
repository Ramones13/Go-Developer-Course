package main

import "fmt"

func main() {
	a := 1
	do(a)
}

func do(v any) {
	a := 10

	buf, ok := v.(int)
	if !ok {
		panic("Не возможно привести к int")
	}
	a = a + buf

	fmt.Println(a)
}
