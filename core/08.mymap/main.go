package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Map example ::::")

	/* Syntax 1 */
	myDetails1 := make(map[string]string)

	myDetails1["firstName"] = "Piyush"
	myDetails1["lastName"] = "Kumar"
	myDetails1["city"] = "Bengaluru"
	myDetails1["state"] = "Karnataka"

	fmt.Println("My map 1 is :::: ", myDetails1)

	delete(myDetails1, "state")
	fmt.Println("myDetails1 map after deletion ::: ", myDetails1)

	/* Syntax 2 */

	myDetails2 := map[string]string{
		"firstName": "Piyush",
		"lastName":  "Kumar",
	}

	fmt.Println("My map 2 is :::: ", myDetails2)

	/* Syntax 3 */

	myDetails3 := map[string]string{}
	myDetails3["firstName"] = "Piyush"
	myDetails3["state"] = "Karnataka"
	fmt.Println("My map 3 is :::: ", myDetails3)

	myDetails3Val1 := myDetails3["firstName"]
	fmt.Println("Value of firstName is :::: ", myDetails3Val1)

	myDetails3Val2, isExists := myDetails3["firstName"]
	fmt.Printf("Value of firstName is %v and isExists=%v \n", myDetails3Val2, isExists)

	_, isLastNameExists := myDetails3["lastName"]
	fmt.Printf("Value of lastName isExists=%v \n", isLastNameExists)

	// looping maps

	// Printing keys and values
	fmt.Println("::::: Printing keys and values :::::: ")
	for key, value := range myDetails3 {

		fmt.Printf("The key is %v and value is %v \n", key, value)
	}

	// Printing only values
	fmt.Println("::::: Printing only values :::::: ")
	for _, value := range myDetails2 {

		fmt.Printf("The value is %v \n", value)
	}

	// Printing only keys
	fmt.Println("::::: Printing only keys :::::: ")
	for key, _ := range myDetails1 {

		fmt.Printf("The key is %v \n", key)
	}

}
