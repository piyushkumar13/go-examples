package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Control flow statements example ::::")

	var logicCount = 23
	var result string

	if logicCount < 10 {
		result = "Logic count is less than 10"

	} else if logicCount > 10 {
		result = "Logic count is greater than 10"

	} else {
		result = "Logic count is equal to 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than or equal to 10")
	}
}
