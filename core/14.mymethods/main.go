package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Structs example ::::")

	student := Student{
		Id:     1,
		Name:   "Piyush",
		Course: "IT",
	}

	fmt.Println("The student1 is ::: ", student.Name)

	student.SetNameWithValueReciever("Kumar")
	fmt.Println("The student1 is ::: ", student.Name) // it will still print Piyush

	student.SetNameWithReferenceReciever("Kirad")
	fmt.Println("The student1 is ::: ", student.Name) // it will print Kirad

	fmt.Println("Getter method value is ::: ", student.GetName())

}

type Student struct {
	Id     int
	Name   string
	Course string
}

// NOTE : we should not mix the methods with value receiver and reference receiver -  this is not recommended.
// Either all should be value receivers or all should be reference receivers.
func (s Student) SetNameWithValueReciever(name string) {

	s.Name = name
	fmt.Println("The name set by value receiver is ::: ", s.Name)
}

func (s *Student) SetNameWithReferenceReciever(name string) {

	s.Name = name
	fmt.Println("The name set by value receiver is ::: ", s.Name)
}

func (s Student) GetName() string {
	return s.Name
}
