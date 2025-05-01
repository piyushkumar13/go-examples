package main

import "fmt"

type Model[T any] struct {
	Data T
}

/*
If you declare methods on a generic type, you must repeat the type parameter declaration on the receiver,
even if the type parameters are not used in the method scope, however in this case you may use the blank identifier _ to make it obvious
*/
func (m *Model[T]) PrintWithInput(input T) {

	fmt.Printf("Complete data : %v-%v \n", m.Data, input)
}

/*
	Here, I have used _ to make it obvious that I am not using type parameter in the method. Even, if I use T also it will work fine.

But _ makes it more clear that I am not using any type parameter in the method.
*/
func (m *Model[_]) Print() {

	fmt.Println("Data : ", m.Data)
}

func main() {

	mInt := Model[int]{Data: 3}

	mInt.PrintWithInput(2)
	mInt.Print()

	mString := Model[string]{Data: "Hello"}

	mString.PrintWithInput("World")
	mString.Print()
}
