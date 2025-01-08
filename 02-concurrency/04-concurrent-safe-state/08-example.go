/*
Encapsulate concurrent safe state management in a custom type
*/
package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	sync.Mutex //composition
	count      int
}

func (sc *SafeCounter) Add(delta int) {
	sc.Lock()
	{
		sc.count += delta
	}
	sc.Unlock()
}

var sf SafeCounter

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", sf.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	sf.Add(1)
}
