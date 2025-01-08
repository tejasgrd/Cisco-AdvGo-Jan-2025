package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}

func main() {
	var count int
	wg := &sync.WaitGroup{}
	flag.IntVar(&count, "count", 0, "Number of goroutines to spin up!")
	flag.Parse()
	fmt.Printf("Starting %d goroutines.. Hit ENTER to start....!\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1) //increment the counter by 1
		go fn(wg, id)
	}
	wg.Wait() //block the execution of this function until the counter becomes 0 (default = 0)
	fmt.Println("Done")
}
