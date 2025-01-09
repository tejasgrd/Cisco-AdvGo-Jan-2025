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
	ch := genNos()
	for no := range ch {
		fmt.Println(no)
	}
}

func genNos() <-chan int {
	ch := make(chan int)
	stopCh := time.After(5 * time.Second)
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
