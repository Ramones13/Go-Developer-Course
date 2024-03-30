package main

import "fmt"

func main() {
	var zeroValuedRune rune
	var zeroValuedByte byte
	fmt.Println("Ожидаем значение 0 для обоих кейсов, потому что они имеют типы int32 и uint8 соответственно. " +
		"Потому что для целых чисел значение по умолчанию = 0 ")
	fmt.Printf("Значение по умолчанию для rune: %d. "+
		"Значение по умолчанию для byte %d \n", zeroValuedRune, zeroValuedByte)
	fmt.Printf("Тип rune: %T. Тип byte %T", zeroValuedRune, zeroValuedByte)
}
