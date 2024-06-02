package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	c := struct {
		Number   int
		Landlord string
		Tenat    string
	}{
		Number:   2,
		Landlord: "Остап",
		Tenat:    "Шура",
	}
	result, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", string(result))

}
