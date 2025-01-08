package main

import (
	"fmt"
	"time"
)

// share memory by communicating

// consumer
func main() {

	ch := add(100, 200)
	/*
		go func() {
			ch <- 10000
		}()
	*/
	result := <-ch
	fmt.Println(result)
}

// producer
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
