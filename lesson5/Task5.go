package main

import "fmt"

type square int
type square58 int

func main() {

	// Task 5.1
	var stringPointerExamle *string
	fmt.Printf("Task 5.1. pointer on string variable %v\n", stringPointerExamle)

	// Task 5.2
	variableExample := 666
	fmt.Printf("Task 5.2. Variable = %v. Pointer = %v\n", variableExample, &variableExample)

	// Task 5.3
	myVariable := 665
	myPointer := &myVariable
	*myPointer = 666
	fmt.Printf("Task 5.3. myVariable = %v\n", myVariable)

	// Task 5.4
	var intExample1 int = 123
	var byteExample byte = 255
	var boolExample1, boolExample2, boolExample3 bool
	var stringExample1 string = "some text"
	var stringExample2 string = "some text"
	var intExample2 int
	var boolExample4 bool

	fmt.Println("Task 5.4.")
	fmt.Printf("int1:\t%v int needs 8\n", &intExample1)
	fmt.Printf("byte:\t%v byte pack with bool\n", &byteExample)
	fmt.Printf("bool1:\t%v bool pack with bool\n", &boolExample1)
	fmt.Printf("bool2:\t%v bool pack with bool\n", &boolExample2)
	fmt.Printf("bool3:\t%v bool pack with bool\n", &boolExample3)
	fmt.Printf("string:\t%v string is placed in other part of memory\n", &stringExample1)
	fmt.Printf("string:\t%v string needs 10 (maybe 8 + size)\n", &stringExample2)
	fmt.Printf("int2:\t%v gap size after last bool = 5\n", &intExample2)
	fmt.Printf("bool4:\t%v int again took 8\n", &boolExample4)

	// Task 5.5
	var localIntExample int = 665
	сhange(&localIntExample)
	fmt.Printf("Task 5.5. local variable after \"сhange\" = %v\n", localIntExample)

	// Task 5.6
	var s square = 25
	fmt.Printf("Task 5.6. %v\n", s)

	// Task 5.7
	var s2 square = 30
	s2 += 15
	fmt.Printf("Task 5.7. %v\n", s2)

	// Task 5.8
	var s3 square58 = 34
	s3 += 10
	fmt.Printf("Task 5.8. %v\n", s3)
}

func сhange(i *int) {
	*i++
}

func (s square58) String() string {
	return fmt.Sprintf("%d м²", s)
}
