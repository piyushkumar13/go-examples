package main

import (
	"fmt"
	"sync"
)

/*
For channels when we are producing into the channel, there should be a listener for it else we will get an error.

	If we are putting two values in the channel, then the listener should be popping out these two values else we will get error.
	Or we can use buffered channel if we dont want to get error if poping statement is less
*/

/* Syntax 1 */
//func main() {
//
//	fmt.Println(":::: Channels example ::::")
//
//	channel := make(chan int)
//	wg := &sync.WaitGroup{}
//
//	wg.Add(2)
//
//	// Listener channel go routine
//	go func(w *sync.WaitGroup, ch chan int) {
//
//		fmt.Println("The channel value is ::: ", <-ch)
//		fmt.Println("The channel value is ::: ", <-ch) // this is required as we are pushing two values in the channel. If we dont give, then we will get error
//		w.Done()
//
//	}(wg, channel)
//
//	// Producer channel go routine
//	go func(w *sync.WaitGroup, ch chan int) {
//
//		ch <- 5
//		ch <- 6
//
//		w.Done()
//	}(wg, channel)
//
//	wg.Wait()
//
//}

/* Syntax 2 */
func main() {

	fmt.Println(":::: Channels example ::::")

	channel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Listener channel go routine - receive only channel go routine
	go func(w *sync.WaitGroup, ch <-chan int) {

		fmt.Println("The channel value is ::: ", <-ch)
		fmt.Println("The channel value is ::: ", <-ch) // this is required as we are pushing two values in the channel. If we dont give, then we will get error
		w.Done()

	}(wg, channel)

	// Producer channel go routine - send only channel go routine
	go func(w *sync.WaitGroup, ch chan<- int) {

		ch <- 5
		ch <- 6

		w.Done()
	}(wg, channel)

	wg.Wait()

}
