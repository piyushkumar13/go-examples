package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	const uri = "https://jsonplaceholder.typicode.com/posts/1"

	fmt.Println("Making a GET web request")

	response, err := http.Get(uri)
	handleError(err, "Call to GET URI failed")

	defer response.Body.Close()

	fmt.Println("The response status is ::: ", response.StatusCode)

	contentInBytes, err := io.ReadAll(response.Body)
	handleError(err, "Reading reponseBody failed")

	var responseString strings.Builder // response builder has advantage of having several methods.

	byteCount, err := responseString.Write(contentInBytes)
	handleError(err, "response string builder write failed")

	fmt.Println("The response body byte count is :::: ", byteCount)
	fmt.Println("The response body is :::: ", responseString.String())

}

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
