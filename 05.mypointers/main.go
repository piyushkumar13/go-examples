package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Pointers example ::::")

	var ptr *int

	fmt.Println("Pointer value is ::: ", ptr) // it will print nil

	var myInt = 2
	var ptr2 *int = &myInt
	fmt.Println("Pointer2 address is ::: ", ptr2)
	fmt.Println("Pointer2 value is :::: ", *ptr2)

	var myStr = "Pointer is interesting"
	ptr3 := &myStr
	fmt.Println("Pointer2 address is ::: ", ptr3)
	fmt.Println("Pointer2 value is :::: ", *ptr3)

}
