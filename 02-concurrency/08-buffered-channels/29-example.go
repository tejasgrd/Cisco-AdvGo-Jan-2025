package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	/*
		errCh := doSomething(wg)
		if err := <-errCh; err != nil {
			fmt.Println("[main] :", err)
		}
	*/

	// ignoring the error doesnt result in a deadlock when a buffered channel is used
	doSomething(wg)
	wg.Wait()
}

func doSomething(wg *sync.WaitGroup) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		if rand.Intn(20)%2 == 0 {
			fmt.Println("[doSomething] - reporting error")
			errCh <- errors.New("something went wrong")
			return
		}
		errCh <- nil
	}()
	return errCh
}
