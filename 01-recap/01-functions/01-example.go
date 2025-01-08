package main

import "fmt"

func main() {
	/*
		add := func(x, y int) int {
			return x + y
		}
	*/
	/*
		var add func(int, int) int
		add = func(x, y int) int {
			return x + y
		}
		fmt.Println(add(100, 200))

		var mutiply func(int, int) int
		mutiply = func(x, y int) int {
			return x * y
		}
		fmt.Println(mutiply(100, 200))
	*/

	// function UDT
	type ArithmaticOperation func(int, int) int

	var add ArithmaticOperation
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	var mutiply ArithmaticOperation
	mutiply = func(x, y int) int {
		return x * y
	}
	fmt.Println(mutiply(100, 200))
}
