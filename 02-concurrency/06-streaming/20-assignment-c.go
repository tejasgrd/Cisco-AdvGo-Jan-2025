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

func dataProducer(start, end int) <-chan int {
	dataCh := make(chan int)

	// producer
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	return dataCh
}

func worker(wg *sync.WaitGroup, dataCh <-chan int, result chan<- int) {
	defer wg.Done()
	for no := range dataCh {
		if isPrime(no) {
			result <- no
		}
	}
}

func doWork(workerCount int, start, end int, resultCh chan<- int) {
	wg := &sync.WaitGroup{}
	dataCh := dataProducer(start, end)
	// spin up the workers
	for workerId := range workerCount {
		wg.Add(1)
		fmt.Println("Starting worker - ", workerId)
		go worker(wg, dataCh, resultCh)
	}
	wg.Wait()
	close(resultCh)
}

func genPrimes(start, end int, workerCount int) <-chan int {
	// share memory by communicating
	resultCh := make(chan int)
	// Work Manager
	go doWork(workerCount, start, end, resultCh)
	return resultCh
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
