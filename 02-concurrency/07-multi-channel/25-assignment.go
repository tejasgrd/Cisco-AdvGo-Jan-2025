/*
modify the below to stop when the user hits ENTER key
DO NOT accept user input in the genNos() function
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := stop()
	ch := genNos(stopCh)
	for no := range ch {
		fmt.Println(no)
	}
}

func stop() <-chan struct{} {
	stopCh := make(chan struct{})
	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	return stopCh
}

func genNos(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
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
