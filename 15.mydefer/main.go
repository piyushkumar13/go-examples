package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Defer example ::::")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	checkDefer()

}

func checkDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}

	fmt.Println("Inside checkDefer method")
}
