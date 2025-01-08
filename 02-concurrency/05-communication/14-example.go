package main

import (
	"fmt"
	"sync"
)

// share memory by communicating

func main() {
	// var ch chan int = make(chan int)
	// var ch = make(chan int)
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch //receive the result from the channel
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result //send the result through the channel
}
