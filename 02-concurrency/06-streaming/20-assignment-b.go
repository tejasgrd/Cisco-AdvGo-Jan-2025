/*
Modify to follow "Share memory by communicating" aka chanels
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100, 5)
	for no := range primes {
		fmt.Println("Prime No :", no)
	}
}

func genPrimes(start, end int, workerCount int) <-chan int {
	// share memory by communicating
	result := make(chan int)
	dataCh := make(chan int)

	// producer
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()

	// Work Manager
	go func() {
		var wg sync.WaitGroup
		// spin up the workers
		for workerId := range workerCount {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for no := range dataCh {
					fmt.Printf("Worker id - %d, processing no %d\n", workerId, no)
					if isPrime(no) {
						result <- no
					}
				}
			}()
		}
		wg.Wait()
		close(result)
	}()
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
