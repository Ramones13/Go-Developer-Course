package main

import (
	"errors"
	"fmt"
)

func main() {

	// Task 9.1
	fmt.Println("Task 9.1")
	count, error := fruitMarket("апельсин")
	if error == nil {
		fmt.Println(count)
	} else {
		fmt.Println(error)
	}

	// Task 9.2
	fmt.Println("Task 9.2")
	mySlice := []int{1, 2, 3}
	for _, v1 := range mySlice {
		fmt.Printf("v1: %v\n", v1)
	CycleBreaker:
		for _, v2 := range mySlice {
			fmt.Printf("\tv2: %v\n", v2)

			for _, v3 := range mySlice {
				fmt.Printf("\t\tv3: %v\n", v3)

				for i4, v4 := range mySlice {
					fmt.Printf("\t\t\tv4: %v\n", v4)
					if i4 == 1 {
						break CycleBreaker
					}
				}
			}
		}
	}

	// Task 9.3
	fmt.Println("Task 9.3")
	checkFood("яблоко")
	checkFood("сова")
	checkFood("помидор")

}

func fruitMarket(fruitName string) (int, error) {
	fruits := map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0,
	}
	fruitCount, ok := fruits[fruitName]
	if !ok {
		return 0, errors.New("Фрукт не найден")
	}
	return fruitCount, nil
}

func checkFood(foodName string) {
	//  "груша", "яблоко", "апельсин", "тыква", "огурец", "помидор"
	switch foodName {
	case "груша", "яблоко", "апельсин":
		println("это фрукт")
	case "тыква", "огурец", "помидор":
		println("это овощ")
	default:
		println("что-то странное…")
	}
}
