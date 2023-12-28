package main

import (
	"fmt"
	"net/http"
	"sync"
)

var completionSignal = []string{"test"} // Its a shared resource

var wg sync.WaitGroup = sync.WaitGroup{}

var lock sync.Mutex = sync.Mutex{}

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

	fmt.Println(completionSignal)
}

func getStatusCode(w string) {
	defer wg.Done()

	response, err := http.Get(w)

	if err != nil {
		fmt.Println("Error while calling ::: ", w)
	} else {

		lock.Lock()
		completionSignal = append(completionSignal, w) // locking shared resource and updating it.
		lock.Unlock()

		fmt.Printf("Successful in calling website=%s and responseCode=%v \n", w, response.StatusCode)
	}
}
