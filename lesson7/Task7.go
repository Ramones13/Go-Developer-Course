package main

import "fmt"

func main() {

	// Task 7.1
	var emptyStrArr [5]string
	fmt.Printf("Task 7.1. %v\n", emptyStrArr)

	// Task 7.2
	strArr2 := [4]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Printf("Task 7.2. %v\n", strArr2)

	// Task 7.3
	strArr3 := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	strArr3[2] = "персик"
	fmt.Printf("Task 7.3. %v\n", strArr3)

	// Task 7.4
	intSlice4 := []int{5, 2, 8, 3, 1, 9}
	fmt.Printf("Task 7.4. %v\n", intSlice4)

	// Task 7.5
	intSlice5 := make([]int, 0, 10)
	fmt.Printf("Task 7.5. %#v емкость: %v \n", intSlice5, cap(intSlice5))

	// Task 7.6
	intSlice6 := make([]int, 0, 10)
	intSlice6 = append(intSlice6, 4, 1, 8, 9)
	fmt.Printf("Task 7.6. %v\n", intSlice6)

	// Task 7.7
	part1 := []int{1, 2, 3}
	part2 := []int{4, 5, 6}
	intSlice7 := append(part1, part2...)
	fmt.Printf("Task 7.7. %v\n", intSlice7)

	// Task 7.8
	intSlice8 := []int{1, 2, 3, 4, 5, 6}
	intSlice8 = append(intSlice8[:3], intSlice8[4:]...)
	fmt.Printf("Task 7.8. %v\n", intSlice8)

}
