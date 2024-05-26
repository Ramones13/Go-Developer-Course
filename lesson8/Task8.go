package main

import "fmt"

func main() {

	// Task 8.1
	animals1 := map[string]struct{}{
		"слон" : {},
		"бегемот" : {},
		"носорог" : {},
		"лев" : {},
	}
	fmt.Printf("Task 8.1. %v\n", animals1)

	// Task 8.2
	animals2 := map[string]int{
		"слон" : 3,
		"бегемот" : 0,
		"носорог" : 5,
		"лев" : 1,
	}
	fmt.Println("Task 8.2")
	animalExists(animals2, "слон")
	animalExists(animals2, "бегемот")
	animalExists(animals2, "осьминог")

	// Task 8.3
	animals3 := map[string]struct{}{
		"слон" : {},
		"бегемот" : {},
		"носорог" : {},
		"лев" : {},
	}
	delete(animals3, "бегемот")
	fmt.Printf("Task 8.3. %v\n", animals3)


	// Task 8.4
	animals4 := map[string]struct{}{
		"слон" : {},
		"бегемот" : {},
		"носорог" : {},
		"лев" : {},
	}
	animals4["выдра"] = struct{}{}
	fmt.Printf("Task 8.4. %v\n", animals4)

	// Task 8.5
	animals5 := map[string]int{
		"слон" : 3,
		"бегемот" : 0,
		"носорог" : 5,
		"лев" : 1,
	}
	animals5["бегемот"] = 2
	fmt.Printf("Task 8.5. %v\n", animals5)

}

func animalExists(animals map[string]int, name string){
	count, ok := animals[name]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", name, count, ok)

}

