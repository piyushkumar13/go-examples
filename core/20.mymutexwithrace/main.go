package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println(":::: Routines with race example ::::")

	scores := []int{}

	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}

	wg.Add(3)
	go func(w *sync.WaitGroup, l *sync.Mutex) {
		fmt.Println("First routine")

		l.Lock()
		scores = append(scores, 1)
		l.Unlock()

		w.Done()

	}(wg, lock)

	go func(w *sync.WaitGroup, l *sync.Mutex) {
		fmt.Println("Second routine")

		l.Lock()
		scores = append(scores, 2)
		l.Unlock()

		w.Done()

	}(wg, lock)

	go func(w *sync.WaitGroup, l *sync.Mutex) {
		fmt.Println("Third routine")

		l.Lock()
		scores = append(scores, 3)
		l.Unlock()

		w.Done()

	}(wg, lock)

	wg.Wait()

	fmt.Println("The scores is :::: ", scores)
}
