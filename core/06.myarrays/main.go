package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Arrays example ::::")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Papaya"
	fruitList[3] = "Mango"

	fmt.Println("Fruit list is :::: ", fruitList)
	fmt.Println("Fruit list len is :::: ", len(fruitList))
	fmt.Printf("Fruit list type is :::: %T\n", fruitList)

	var fruitList2 = [3]string{"Pineapple", "Stawberry", "Pear"}

	fmt.Println("Fruit list2 is :::: ", fruitList2)
	fmt.Println("Fruit list2 len is :::: ", len(fruitList2))

	vegList := [4]string{}
	vegList[0] = "Potato"
	vegList[1] = "Mushroom"
	vegList[2] = "Peas"

	fmt.Println("Veg list is :::: ", vegList)
	fmt.Println("Veg list len is :::: ", len(vegList))

}
