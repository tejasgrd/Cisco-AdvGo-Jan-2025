package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = 99.99
	x = "Dolore exercitation nisi veniam cupidatat enim pariatur qui ut ut."
	x = true
	x = struct{}{}
	fmt.Println(x)

	// x = 100
	// fmt.Println(x * 2)
	// fmt.Println(x.(int) * 2)

	// type assertion
	x = true
	if val, ok := x.(int); ok {
		fmt.Println(val * 2)
	} else {
		fmt.Println("x is not an int")
	}
}
