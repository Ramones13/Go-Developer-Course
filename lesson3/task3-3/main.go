package main

import "fmt"

const usefulConst = "This is an useful constant"

func main() {

	const usefulConst = "Local useful constant"

	fmt.Printf("%v", usefulConst)
}
