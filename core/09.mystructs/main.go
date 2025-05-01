package main

import (
	"fmt"
)

func main() {

	fmt.Println(":::: Structs example ::::")

	// Syntax 1 to create struct object.
	student1 := Student{
		Id:     1,
		Name:   "Piyush",
		Course: "IT",
	}

	fmt.Println("The student1 is :::: ", student1)
	fmt.Println("The student1 name is :::: ", student1.Name)
	fmt.Println("The student1 id is :::: ", student1.Id)
	fmt.Println("The student1 course is :::: ", student1.Course)

	// Syntax 2 to create struct object.
	student2 := Student{1, "Piyush", "IT"}

	fmt.Println("The student2 is :::: ", student2)
	fmt.Println("The student2 name is :::: ", student2.Name)
	fmt.Println("The student2 id is :::: ", student2.Id)
	fmt.Println("The student2 course is :::: ", student2.Course)

	// Syntax 3 to create struct object.
	student3 := new(Student)
	student3.Id = 2
	student3.Name = "ABC"
	student3.Course = "CS"

	fmt.Println("The student3 is ::: ", student3)
	fmt.Println("The student3 is ::: ", *student3)
	fmt.Println("The student3 address is ::: ", &student3)
}

type Student struct {
	Id     int
	Name   string
	Course string
}
