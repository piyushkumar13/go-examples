package main

import "fmt"

type MyGetter[T any] interface {
	Get() T
}

type MyModel[T any] struct {
	Data T
}

func (m *MyModel[T]) Get() T {
	return m.Data
}

func main() {

	var getter1 MyGetter[int]
	getter1 = &MyModel[int]{Data: 2}
	fmt.Println("Getter1: ", getter1)

	var getter2 MyGetter[[]int]
	getter2 = &MyModel[[]int]{Data: []int{1, 2, 3}}
	fmt.Println("Getter2: ", getter2)
}
