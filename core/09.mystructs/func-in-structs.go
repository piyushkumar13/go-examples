package main

import "fmt"

/* Way 1 */
type Person1 struct {
	Id         int
	FirstName  string
	Age        int
	LastName   string
	FullName   func() string
	IsEligible func(int) bool
}

/* Way 2 */

type myFullNameFunc func() string
type myIsElibibleFunc func(int) bool

type Person2 struct {
	Id         int
	FirstName  string
	Age        int
	LastName   string
	FullName   myFullNameFunc
	IsEligible myIsElibibleFunc
}

func main() {

	person1 := Person1{

		Id:        1,
		FirstName: "Piyush",
		Age:       35,
		LastName:  "Kumar",

		//FullName: func() string { // We cannot do this in golang. You need to create the instance first and then use dot notation to define function.
		//    return FirstName + LastName
		//},
	}

	person1.FullName = func() string {

		return person1.FirstName + " " + person1.LastName
	}

	person1.IsEligible = func(age int) bool {

		return age > 30
	}

	fmt.Println("Person1 is ::: ", person1)
	fmt.Println("Person1 full name is ::: ", person1.FullName())
	fmt.Println("Person1 is eligible ::: ", person1.IsEligible(35))

	/* Defining type Person 2 in different ways */

	/* Way 1 */
	person2 := Person2{

		Id:        2,
		FirstName: "Gaurav",
		Age:       31,
		LastName:  "Kumar",
	}

	person2.FullName = func() string {
		return person2.FirstName + " " + person2.LastName
	}

	person2.IsEligible = func(age int) bool {
		return age > 31
	}

	/* Way 2 */
	person2a := Person2{

		Id:         2,
		FirstName:  "Gaurav",
		Age:        31,
		LastName:   "Kumar",
		FullName:   fullNameFunc,
		IsEligible: isEligibleFunc,
	}

	fmt.Println("Person 2a fullname ::: ", person2a.FullName())
	fmt.Println("Person 2a isEligible ::: ", person2a.IsEligible(30))

	/* Way 3 */

	person2b := Person2{

		Id:        2,
		FirstName: "Gaurav",
		Age:       31,
		LastName:  "Kumar",
	}

	person2b.FullName = fullNameFunc
	person2b.IsEligible = isEligibleFunc

	fmt.Println("Person 2b fullname ::: ", person2a.FullName())
	fmt.Println("Person 2b isEligible ::: ", person2a.IsEligible(30))

}

func fullNameFunc() string {
	return "Dummy" + "Name"
}

func isEligibleFunc(age int) bool {
	return age > 31
}
