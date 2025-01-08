package main

import (
	"fmt"
	"sync"
)

func main() {
	var result int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = add(100, 200)
	}()
	wg.Wait()
	fmt.Println("result :", result)
}

// 3rd party library api (cannot change the code)
func add(x, y int) int {
	return x + y
}
