package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(":::: Slices example ::::")

	/* Syntax 1 */
	var fruitSlice = []string{"Apple", "Mango", "Pineapple"}
	fmt.Println("Fruit slice is ::: ", fruitSlice)
	fmt.Printf("Fruit slice type is %T\n", fruitSlice)

	fruitSlice = append(fruitSlice, "Stawberry", "Pear", "Watermelon")

	fmt.Println("Updated fruit slice is ::: ", fruitSlice)

	/* Syntax 2 */

	scores := make([]int, 3)
	scores[0] = 0
	scores[1] = 1
	scores[2] = 2
	//scores[3] = 4 // if I add this, I will get an error saying  index out of range [3] with length 3 but we can still add using append
	fmt.Println("Scores is ::: ", scores)
	fmt.Println("Scores len is ::: ", len(scores))

	scores = append(scores, 3, 4, 5, 7)

	fmt.Println("Updated scores is ::: ", scores)
	fmt.Println("Updated scores len is ::: ", len(scores))

	/* Fetch from slices */

	fmt.Println("Slice1 of score ::: ", scores[:3])  // it will fetch elements from slice at index 0 till 2 index, 3 is exclusive
	fmt.Println("Slice2 of score ::: ", scores[1:3]) // it will fetch elements from slice at index 1 till 2 index, 3 is exclusive

	/* Syntax 3 and Sorting of slices */

	marks := []int{}
	marks = append(marks, 10, 8, 9, 2, 4, 2, 1)

	sort.Ints(marks)
	fmt.Println("Sorted slice is ::: ", marks)

	/* Removing elements from slices */

	subjects := []string{"English", "Hindi", "Geography", "Maths", "Computer Science"}
	fmt.Println("The subjects are ::: ", subjects)

	removeSubjectFromindex := 2
	subjects = append(subjects[:removeSubjectFromindex], subjects[removeSubjectFromindex+1:]...)

	fmt.Println("The subjects slice after removal ::: ", subjects)

}
