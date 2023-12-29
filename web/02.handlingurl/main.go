package main

import (
	"fmt"
	"net/url"
)

const uri = "https://example.com/learn?course=golang&username=piyush&isloggedin=true"

func main() {

	fmt.Println("Handling URL example")

	parsedUrl, _ := url.Parse(uri)

	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)
	fmt.Println(parsedUrl.Query())

	// Building URI

	builtUrl := &url.URL{
		Scheme:   "https",
		Host:     "myexample",
		Path:     "somemorelearning",
		RawQuery: "course=golang",
	}

	fmt.Println("The built url is ::: ", builtUrl.String())

}
