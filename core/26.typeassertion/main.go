package main

import "fmt"

/* Type assertion is nothing, but casting the type to another. Its type casting basically. */

func main() {

	var name interface{} = "Piyush"

	fmt.Printf("Type is %T, Value is %v\n", name, name) // type will be printed as string since it is actually string.

	nameStr := name.(string)
	fmt.Println("The name is :::", nameStr)
	fmt.Println("The name is :::", nameStr)

	name = 10
	//nameStr = name.(string) // This will throw exception since trying to cast integer to string.

	/* Go provides a syntax while type asserting which returns a boolean when value cannot be casted instead of panicing. */

	nameInt, isString := name.(string)

	if isString {
		fmt.Println("The nameInt is :::", nameInt)
	} else {
		fmt.Println("Not a valid casting")
	}

	if n, s := name.(string); s { // systax where if we can have assignment as well as condition check.

		fmt.Println("The name is :::", n)

	}
}
