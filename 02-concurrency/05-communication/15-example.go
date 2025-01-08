package main

import "fmt"

func main() {

	/*
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		go func() { //scheduling through the scheduler
			ch <- 100 // (2-NonBlocking)
		}()
		data := <-ch // (1-Blocking, 3-Unblocked)
		fmt.Println(data)
	*/

	ch := make(chan int)
	go func() { //scheduling through the scheduler
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100
}
