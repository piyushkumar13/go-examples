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
	fmt.Println("Json unmarshalling/deserialization")

	unmarshalJson()
}

func unmarshalJson() {

	jsonVal := []byte(`

        {
                "id": 1,
                "name": "Piyush",
                "course": "IT",
                "subjects": [
                        "DS",
                        "OS",
                        "DBMS"
                ],
                "city": "Bengaluru"
        }

`)

	var student Student

	isValidJson := json.Valid(jsonVal)

	if !isValidJson {
		panic("Json is not valid")
	}

	err := json.Unmarshal(jsonVal, &student)
	//jsonValue, err := json.MarshalIndent(students, "", "\t") // this will print pretty jsonValue
	handleError(err, "Error while unmarshalling")

	fmt.Println("The unmarshalled object is ::: ", student)
}

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
