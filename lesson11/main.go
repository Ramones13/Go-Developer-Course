package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	Message string
}

func (e *myFirstError) Error() string {
	return e.Message
}

func main() {

	// Task 11.1
	fmt.Println("Task 11.1")
	//err1 := errors.New("ошибка1")
	err1 := &myFirstError{
		Message: "ошибка1",
	}

	err2 := fmt.Errorf("ошибка2: %w", err1)
	err3 := fmt.Errorf("ошибка2: %w", err2)
	fmt.Println(err3)

	// Task 11.2
	fmt.Println("Task 11.2")
	fmt.Println(errors.Unwrap(err3))

	// Task 11.3
	fmt.Println("Task 11.3")
	if errors.Is(err3, err1) {
		fmt.Println("действия при вхождении нашей ошибки в последней полученной")
	}

	// Task 11.4
	fmt.Println("Task 11.4")
	var tempErr *myFirstError
	fmt.Printf("первая ошибка приводится к myFirstError: %v.\t\tЕе тип %T\n",
		errors.As(err1, &tempErr), err1)
	fmt.Printf("конечная ошибка приводится к myFirstError: %v.\tЕе тип %T\n",
		errors.As(err3, &tempErr), err3)

	// тестовый код не для проверки
	//if errors.As(err3, &tempErr) {
	//	fmt.Printf("u can envoke nested element from wraped error: %v\n", tempErr.Message)
	//} else {
	//	fmt.Println("you can't use nested")
	//}
}
