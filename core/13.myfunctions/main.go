package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Functions example ::::")

	greeting := greeting()
	fmt.Println(greeting)
	addedNum := addr(3, 4)
	fmt.Println("The added num is :::: ", addedNum)
	addedNumForStr, str := addrWithStrMsg(5, 6)
	fmt.Println("The added num with string msg is :::: ", addedNumForStr, str)
}

func greeting() string {
	return "Welcome to the functions"
}

func addr(a int, b int) int {

	return a + b
}

func addrWithStrMsg(a int, b int) (int, string) {

	return a + b, "We found the result"
}
