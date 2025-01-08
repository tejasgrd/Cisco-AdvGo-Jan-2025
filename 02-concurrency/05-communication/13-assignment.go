/*
Modify so that the logic for checking if a number is prime or not is executed concurrently
*/
package main

import "fmt"

func main() {
	primes := genPrimes(2, 100)
	for _, no := range primes {
		fmt.Println("Prime No :", no)
	}
}

func genPrimes(start, end int) []int {
	var result []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			result = append(result, no)
		}
	}
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
