package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		multiply(100, 200)
	*/

	// ver 2.0
	/*
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed!")

		log.Println("Operation started")
		multiply(100, 200)
		log.Println("Operation completed!")
	*/

	/*
		logAdd(100, 200)
		logMultiply(100, 200)
	*/

	/*
		logOperation(add, 100, 200)
		logOperation(multiply, 100, 200)
		logOperation(func(i1, i2 int) {
			fmt.Println("Subtract Result :", i1-i2)
		}, 100, 200)
	*/

	// ver 3.0 (logging)
	/*
		logAdd := wrapLogger(add)
		logAdd(100, 200)

		logMultiply := wrapLogger(multiply)
		logMultiply(100, 200)
	*/

	// ver 4.0 (profiling)
	/*
		profileAdd := wrapProfiler(add)
		profileAdd(100, 200)

		profileMultiply := wrapProfiler(multiply)
		profileMultiply(100, 200)
	*/

	// ver 5.0 (logging & profiling)
	/*
	   logAdd := wrapLogger(add)
	   	profileLogAdd := wrapProfiler(logAdd)
	*/
	profileLogAdd := wrapProfiler(wrapLogger(add))
	profileLogAdd(100, 200)
}

// ver 4.0 (adding profiling)
func wrapProfiler(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

// ver 3.0 (adding logging)
func wrapLogger(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		operation(x, y)
		log.Println("Operation completed!")
	}
}

// ver 2.0
/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed!")
}

func logMultiply(x, y int) {
	log.Println("Operation started")
	multiply(x, y)
	log.Println("Operation completed!")
}
*/

func logOperation(operation func(int, int), x, y int) {
	log.Println("Operation started")
	operation(x, y)
	log.Println("Operation completed!")
}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func multiply(x, y int) {
	fmt.Println("Multiply Result :", x*y)
}
