/*
Modify to follow "Share memory by communicating" aka chanels
*/
package main

import (
	"fmt"
	"sync"
)

type PrimeResult struct {
	no      int
	isPrime bool
}

func main() {
	primes := genPrimes(2, 100, 5)
	for result := range primes {
		fmt.Printf("%+v\n", result)
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

func worker(wg *sync.WaitGroup, dataCh <-chan int, result chan<- PrimeResult) {
	defer wg.Done()
	for no := range dataCh {
		if isPrime(no) {
			result <- PrimeResult{no: no, isPrime: true}
			continue
		}
		result <- PrimeResult{no: no, isPrime: false}
	}
}

func doWork(workerCount int, start, end int, resultCh chan<- PrimeResult) {
	wg := &sync.WaitGroup{}
	dataCh := dataProducer(start, end)
	// spin up the workers
	for workerId := range workerCount {
		wg.Add(1)
		fmt.Println("Starting worker - ", workerId)

		// start the workers
		go worker(wg, dataCh, resultCh)
	}
	wg.Wait()
	close(resultCh)
}

func genPrimes(start, end int, workerCount int) <-chan PrimeResult {
	// share memory by communicating
	resultCh := make(chan PrimeResult)
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
