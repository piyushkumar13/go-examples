package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(":::: Switch statement example ::::")

	num := rand.Intn(3) + 1

	fmt.Println("The random number picked is ::: ", num)

	switch num {

	case 1:
		fmt.Println("The num is 1")
	case 2:
		fmt.Println("The num is 2")
	case 3:
		fmt.Println("The num is 3")
	default:
		fmt.Println("The num is default")

	}

	switch num {

	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fmt.Println("The num is 3")
	default:
		fmt.Println("The num is default")

	}
}
