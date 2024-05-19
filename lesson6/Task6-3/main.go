// Необходимо объявить глобальную структуру contract с полями: ID int, Number string, Date string.
// Далее создать экземпляр структуры со значениями полей: ID=1, Number=«#000A\n101», Date=«2024-01-31».
// При передачи экземпляра структуры в fmt.Println
// в консоли должно отображаться: Договор № #000A\n101 от 2024-01-31

package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}

func (c contract) String() string {
	return fmt.Sprintf("Договор № %v от %v", c.Number, c.Date)
}

func main() {

	contractExample := contract{
		ID:     1,
		Number: `#000A\n101`,
		Date:   "2024-01-31",
	}

	fmt.Println(contractExample)
}
