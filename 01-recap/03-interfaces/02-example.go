/*
Interfaces need not be "explicitly" implemented
Interfaces can be composed
*/

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Scale(times float64) {
	r.Height = r.Height * times
	r.Width = r.Width * times
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Argument is neither a Circle nor a Rectangle")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }: // any object with Area() method
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Argument does not IMPLEMENT the Area() method")
	}
}
*/

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

// ver 2.0 (introduce Perimeter() methods)

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

func PrintPerimeter(x interface{ Perimeter() float64 }) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// ver 3.0

/*
func PrintStats(x interface {
	Area() float64
	Perimeter() float64
}) {
	fmt.Println("Stats")
	fmt.Println("--------------")
	PrintArea(x)      //x should be interface{ Area() float64 }
	PrintPerimeter(x) //x should be interface{ Perimeter() float64 }
	fmt.Println("--------------")
}
*/

func PrintStats(x interface {
	interface{ Area() float64 }
	interface{ Perimeter() float64 }
}) {
	fmt.Println("Stats")
	fmt.Println("--------------")
	PrintArea(x)      //x should be interface{ Area() float64 }
	PrintPerimeter(x) //x should be interface{ Perimeter() float64 }
	fmt.Println("--------------")
}

func main() {
	c := Circle{Radius: 15}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/

	PrintStats(c)

	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/

	PrintStats(r)

	/*
		r.Scale(0.5)
		fmt.Println(r)
	*/

	s := Square{Side: 14}
	/*
		PrintArea(s)
		PrintPerimeter(s)
	*/

	PrintStats(s)

	// PrintArea(100)
}
