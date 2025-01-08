package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(<-ch)
	}
}

func genNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Println("[genNos] count =", count)
	for no := range count {
		ch <- (no + 1) * 10
	}
	close(ch)
}
