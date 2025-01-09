package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for no := range ch {
		fmt.Println(no)
	}
}

func timeout(d time.Duration) <-chan time.Time {
	stopCh := make(chan time.Time)
	go func() {
		time.Sleep(5 * time.Second)
		stopCh <- time.Now()
	}()
	return stopCh
}

func genNos() <-chan int {
	ch := make(chan int)
	stopCh := timeout(5 * time.Second)
	go func() {
		var no int
	LOOP:
		for {
			select {
			case ch <- no:
				time.Sleep(500 * time.Millisecond)
				no++
			case <-stopCh:
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}
