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

	fmt.Println("The student is :::: ", student)
	fmt.Println("The student name is :::: ", student.Name)
	fmt.Println("The student id is :::: ", student.Id)
	fmt.Println("The student course is :::: ", student.Course)

	student1 := new(Student)
	student1.Id = 2
	student1.Name = "ABC"
	student1.Course = "CS"

	fmt.Println("The student1 is ::: ", student1)
	fmt.Println("The student1 is ::: ", *student1)
	fmt.Println("The student1 address is ::: ", &student1)
}

type Student struct {
	Id     int
	Name   string
	Course string
}
