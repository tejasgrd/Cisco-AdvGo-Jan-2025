/*
Modify to follow "Share memory by communicating" aka chanels
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	primes := genPrimes(2, 1000)
	for _, no := range primes {
		fmt.Println("Prime No :", no)
	}
}

func genPrimes(start, end int) []int {
	// communicate by sharing memory
	var result []int
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				mutex.Lock()
				{
					result = append(result, no)
				}
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
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
