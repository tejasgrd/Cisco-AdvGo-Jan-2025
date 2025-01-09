/*
Modify to follow "Share memory by communicating" aka chanels
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 100)
	for no := range primes {
		fmt.Println("Prime No :", no)
	}
}

func genPrimes(start, end int) <-chan int {
	// share memory by communicating
	var result chan int = make(chan int)
	go func() {
		var wg sync.WaitGroup
		for no := start; no <= end; no++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if isPrime(no) {
					result <- no
				}
			}()
		}
		wg.Wait() //block
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
