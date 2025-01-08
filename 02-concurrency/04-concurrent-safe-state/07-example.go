package main

import "fmt"

var count int

func main() {
	for range 200 {
		go increment()
	}
	fmt.Println("count :", count)
}

func increment() {
	count++
}
