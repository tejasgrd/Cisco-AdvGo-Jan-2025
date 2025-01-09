package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("data from ch3 :", <-ch3)
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch3 <- 300
	}()

	wg.Wait()
}
