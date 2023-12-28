package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(":::: Routines example ::::")

	websites := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, website := range websites {

		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
}

func getStatusCode(w string) {
	defer wg.Done()

	response, err := http.Get(w)

	if err != nil {
		fmt.Println("Error while calling ::: ", w)
	} else {
		fmt.Printf("Successful in calling website=%s and responseCode=%v \n", w, response.StatusCode)
	}
}
