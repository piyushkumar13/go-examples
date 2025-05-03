package main

import (
	"encoding/json"
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

	/*
	   Creating empty/nil8- slice : https://stackoverflow.com/q/29164375
	   Some notes from ChatGPT

	   In Go, when you declare a slice without initializing it like this:
	          var myslice []int
	          fmt.Println(myslice)

	   You're declaring a nil slice of int. This will print an empty slice representation — [].
	   Even though it's nil under the hood (i.e., myslice == nil is true), when printed with fmt.Println, it shows up as [].

	           Nil Slice
	              var a []int           // a is nil
	              fmt.Println(a)        // Output: []
	              fmt.Println(a == nil) // Output: true
	              fmt.Println(len(a))   // Output: 0
	              fmt.Println(cap(a))   // Output: 0

	            Empty Slice
	             b := []int{}          // b is not nil, just empty
	             fmt.Println(b)        // Output: []
	             fmt.Println(b == nil) // Output: false
	             fmt.Println(len(b))   // Output: 0
	             fmt.Println(cap(b))   // Output: 0


	   Whats the difference between []int{}   and make([]int,0) ?

	   Both []int{} and make([]int, 0) create empty slices, but there are subtle differences in how they are created under the hood.
	   In both the cases, slice contents is empty slice,length(len()) is 0, capacity(cap(0)), nil check is false, but in case of
	   []int{} underlying array will not be allocated but in case of make([]int, 0) underlying array will be allocated with length zero.


	   Recommendations :
	   For general empty slice creation, []int{} is the fastest and cleanest though you can also use make([]int, 0) but this might be marginally slower
	   due to backed array allocation(basically slightly more overhead) but its not something which make it very slow so both
	   []int{} and make([]int, 0) are approximately same in performance. But use []int{} since its more cleaner and less verbose.

	   But If you're optimizing for append-heavy workloads, use make([]T, 0, N) with a precomputed capacity i.e N here.

	   However, if you want a nil slice you can use var mySlice []int

	*/

	myEmptySlice1 := []int{}            // This is empty slice
	myEmptySlice2 := make([]int, 0)     // This is empty slice, with initial length zero
	myEmptySlice3 := make([]int, 0, 10) // This is empty slice, with initial length zero and capacity 10

	var myEmptySlice4 []int // This is nil slice

	fmt.Printf("myEmptySlice1 : %v, myEmptySlice2 : %v, myEmptySlice4 : %v,  myEmptySlice4 : %v \n",
		myEmptySlice1, myEmptySlice2, myEmptySlice3, myEmptySlice4) // myEmptySlice4 is actually nil slice but when we print it gives [] but when we compare with nil it will return true

	fmt.Println("myEmptySlice1 nil check", myEmptySlice1 == nil) // false
	fmt.Println("myEmptySlice2 nil check", myEmptySlice2 == nil) // false
	fmt.Println("myEmptySlice3 nil check", myEmptySlice3 == nil) // false
	fmt.Println("myEmptySlice4 nil check", myEmptySlice4 == nil) // true

	fmt.Println("Understanding json serialization of empty/nil slice")

	marshalJsonForSlice1, _ := json.Marshal(myEmptySlice1)
	marshalJsonForSlice2, _ := json.Marshal(myEmptySlice2)
	marshalJsonForSlice3, _ := json.Marshal(myEmptySlice3)
	marshalJsonForSlice4, _ := json.Marshal(myEmptySlice4)

	fmt.Println("Serialized myEmptySlice1 is ::: ", string(marshalJsonForSlice1)) // prints []
	fmt.Println("Serialized myEmptySlice2 is ::: ", string(marshalJsonForSlice2)) // prints []
	fmt.Println("Serialized myEmptySlice3 is ::: ", string(marshalJsonForSlice3)) // prints []
	fmt.Println("Serialized myEmptySlice4 is ::: ", string(marshalJsonForSlice4)) // prints null
}
