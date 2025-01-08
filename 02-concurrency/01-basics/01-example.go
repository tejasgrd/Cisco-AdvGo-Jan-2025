package main

import (
	"fmt"
	"time"
)

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}

func main() {
	go f1() //schedule the execution of f1() through the scheduler
	f2()

	// Poor man's synchronization techniques (DO NOT USE THEM)

	time.Sleep(1 * time.Second) // block the main() so that the scheduler can look for other goroutines scheduled and execute them (cooperative multitasking)

	// fmt.Scanln()
}
