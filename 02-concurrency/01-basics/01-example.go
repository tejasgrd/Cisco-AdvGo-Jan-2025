package main

import (
	"fmt"
)

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func main() {
	go f1() //schedule the execution of f1() through the scheduler
	f2()

	// Poor man's synchronization techniques (DO NOT USE THEM)

	// time.Sleep(1 * time.Second) // block the main() so that the scheduler can look for other goroutines scheduled and execute them (cooperative multitasking)

	// fmt.Scanln()
}
