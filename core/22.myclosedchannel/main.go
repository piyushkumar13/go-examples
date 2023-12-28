package main

import (
	"fmt"
	"sync"
)

///*
//Following code will give error and thats the expected error saying send on closed channel. As we have closed the channel and then we are trying to push value to the channel.
//And then we are consuming it.
//*/
//func main() {
//
//	fmt.Println(":::: Closed channel example ::::")
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
//		fmt.Println("The channel value is ::: ", <-ch)
//
//		w.Done()
//
//	}(wg, channel)
//
//	// Producer channel go routine
//	go func(w *sync.WaitGroup, ch chan int) {
//
//		close(ch)
//		ch <- 5
//		ch <- 6
//
//		w.Done()
//	}(wg, channel)
//
//	wg.Wait()
//}

///*
//This code is fine as after closing the channel, we are not sending any value to the channel.
//But, in this you will observe receive value 0. So, what if before closing channel, we send ch <- 0 value. It will become difficult to find out
//whether its coming due to closed channel or the producer actually send 0 value. This can be solved by using comma, ok syntax.
//*/
//func main() {
//
//	fmt.Println(":::: Closed channel example ::::")
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
//		w.Done()
//	}(wg, channel)
//
//	// Producer channel go routine
//	go func(w *sync.WaitGroup, ch chan int) {
//
//		close(ch)
//
//		w.Done()
//	}(wg, channel)
//
//	wg.Wait()
//
//}

/*
The problem to find out whether zero coming due to closed channel or the producer actually send 0 value. This can be solved by using comma, ok syntax
as illustrated below.
*/
func main() {

	fmt.Println(":::: Closed channel example ::::")

	channel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// Listener channel go routine
	go func(w *sync.WaitGroup, ch chan int) {

		val, isChannelOpen := <-ch
		fmt.Println("The channel value and isChannelOpen ::: ", val, isChannelOpen)
		w.Done()
	}(wg, channel)

	// Producer channel go routine
	go func(w *sync.WaitGroup, ch chan int) {

		ch <- 0
		close(ch)

		w.Done()
	}(wg, channel)

	wg.Wait()

}
