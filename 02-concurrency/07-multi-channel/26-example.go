/*
modify the below to stop when the user hits ENTER key
DO NOT accept user input in the genNos() function
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stopCh := stop()
	ch := genNos(stopCh)
	for data := range ch {
		fmt.Println(data)
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

func genNos(stopCh <-chan struct{}) <-chan string {
	resultCh := make(chan string)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go genOddNos(stopCh, resultCh, wg)

		wg.Add(1)
		go genEvenNos(stopCh, resultCh, wg)

		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func genEvenNos(stopCh <-chan struct{}, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 0; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Even : %d", no):
			time.Sleep(500 * time.Millisecond)
			no++
		case <-stopCh:
			break LOOP
		}
	}
}

func genOddNos(stopCh <-chan struct{}, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 1; ; no += 2 {
		select {
		case resultCh <- fmt.Sprintf("Odd : %d", no):
			time.Sleep(300 * time.Millisecond)
			no++
		case <-stopCh:
			break LOOP
		}
	}
}
