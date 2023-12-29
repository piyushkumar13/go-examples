package main

import (
	"fmt"
	"io"
	"net/http"
)

const uri = "https://google.com"

func main() {

	fmt.Println("Making a GET web request")

	response, err := http.Get(uri)
	handleError(err, "Call to GET URI failed")

	defer response.Body.Close()

	responseBodyBytes, err := io.ReadAll(response.Body)
	handleError(err, "Reading reponseBody failed")

	fmt.Println("The response is ::: ", string(responseBodyBytes))
	fmt.Println("The response status is ::: ", response.StatusCode)
}

func handleError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
