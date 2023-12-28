package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(":::: Routines example ::::")

	go greeting("World")
	greeting("Hello")

	// this will allow main method to wait so that thread(go routine) can complete. But, using sleep is not the good way. For this we have wait groups.
	time.Sleep(3 * time.Millisecond)
}

func greeting(s string) {

	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
