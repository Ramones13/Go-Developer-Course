package main

import "fmt"

const fibonacciCount23 int = 23

func main() {

	// Task 4.1
	hello41()

	// Task 4.2
	func() {
		fmt.Println("Hello, Go! From task 4-2")
	}()

	// Task 4.3
	hello43(func() {
		fmt.Println("Hello, Go! From task 4-3")
	})

	// Task 4.4
	hello44()()

	// Task 4.5
	hello45()

	// Task 4.6
	fmt.Printf("Fibonacci for %v = %v", fibonacciCount23, fibonacci(fibonacciCount23))
}

// Task 4.1
func hello41() {
	fmt.Println("Hello, Go! From task 4-1")
}

// Task 4.3
func hello43(fn func()) {
	fn()
}

// Task 4.4
func hello44() func() {
	return func() {
		fmt.Println("Hello, Go! From task 4-4")
	}
}

// Task 4.5
func hello45() {
	defer fmt.Println("завершение вызова функции по задаче 4-5")
	fmt.Println("Hello, Go! From task 4-5")
}

// Task 4.6
func fibonacci(deep int) int {
	if deep == 0 || deep == 1 {
		return deep
	}
	return fibonacci(deep-1) + fibonacci(deep-2)
}
