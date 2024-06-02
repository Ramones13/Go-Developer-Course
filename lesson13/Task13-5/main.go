package main

import (
	"fmt"
	Aibolit_1_0 "v1_0"
	Aibolit_1_1 "v1_1"
	Aibolit_2_0 "v2_0"
	Aibolit_2_1 "v2_1"
)

func main() {

	// Вызов версии 1.0
	err := Aibolit_1_0.Do("lesson13/Task13-5/source_json", "./lesson13/Task13-5")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Успешно записан файл v1.0")
	}

	// Вызов версии 1.1
	err = Aibolit_1_1.Do("lesson13/Task13-5/source_json", "./lesson13/Task13-5")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Успешно записан файл v1.1")
	}

	// Вызов версии 2.0
	err = Aibolit_2_0.Do("lesson13/Task13-5/source_json", "./lesson13/Task13-5/result_2_0.xml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Успешно записан файл v2.0")
	}

	// Вызов версии 2.1
	err = Aibolit_2_1.Do("lesson13/Task13-5/source_json", "./lesson13/Task13-5/result_2_1.xml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Успешно записан файл v2.0")
	}

	// Вызов версии тест
	//err = do("lesson13/Task13-5/source_json", "./lesson13/Task13-5/result_2_0.xml")
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Успешно записан файл v2.0")
	//}
}
