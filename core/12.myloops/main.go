package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Looping statement example ::::")

	days := []string{"Sunday", "Monday", "Tuesday"}

	fmt.Println("::::::: Looping through using index ::::::: ")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("::::::: Looping through using range with index ::::::: ")
	for index := range days {

		fmt.Println(index)
		fmt.Println(days[index])
	}

	fmt.Println("::::::: Looping through using range with index and value ::::::: ")
	for index, day := range days {

		fmt.Println(index)
		fmt.Println(day)
	}

	fmt.Println("::::::: Looping through using range with only value ::::::: ")
	for _, day := range days {
		fmt.Println(day)
	}

	fmt.Println("::::::: Looping through like while loop ::::::: ")
	i := 0
	for i < 4 {
		fmt.Println(i)
		i++
	}

	fmt.Println("::::::: Looping through like while loop ::::::: ")
	p := 0
	for p < 10 {

		if p == 5 {
			break
		}
		if p == 1 {
			p++
			continue
		}
		fmt.Println(p)
		p++
	}

}
