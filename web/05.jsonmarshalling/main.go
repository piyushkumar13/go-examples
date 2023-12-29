package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Course   string   `json:"course"`
	Subjects []string `json:"subjects,omitempty"`
	City     string   `json:"city,omitempty"`
	Password string   `json:"-"` // this will not be serialized.
}

func main() {
	fmt.Println("Json marshalling/serialization")

	marshalObject()
}

func marshalObject() {

	//students := []Student{
	//	{Id: 1, Name: "Piyush", Course: "IT", Subjects: []string{"DS", "OS", "DBMS"}, City: "Bengaluru"},
	//	{Id: 1, Name: "Deepak", Course: "Commerce", Subjects: []string{"Maths", "BS"}, City: ""},
	//	{Id: 1, Name: "Baban", Course: "Commerce", Subjects: []string{"Maths", "Accounts", "HR"}},
	//	{Id: 1, Name: "Anand", Course: "IT", Subjects: []string{"DS", "OS"}},
	//}

	students := []Student{
		{1, "Piyush", "IT", []string{"DS", "OS", "DBMS"}, "Bengaluru", "123"}, // as I am providing all values so keys can be skipped
		{1, "Deepak", "Commerce", []string{"Maths", "BS"}, "", "345"},
		{Id: 1, Name: "Baban", Course: "Commerce", Subjects: []string{"Maths", "Accounts", "HR"}}, // since I didnt provide the values for all fields thats y key has to be mentioned
		{1, "Anand", "IT", nil, "", "987"},
	}

	//jsonValue, err := json.Marshal(students)
	jsonValue, err := json.MarshalIndent(students, "", "\t") // this will print pretty jsonValue
	handleError(err, "Error while parshing")

	fmt.Println("The jsonValue bytes is ::: ", jsonValue)   // this will print bytes
	fmt.Printf("The jsonValue is ::: %s\n", jsonValue)      // but this will print jsonValue string as we used %s
	fmt.Println("The jsonValue is ::: ", string(jsonValue)) // this will print jsonValue string
}

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
