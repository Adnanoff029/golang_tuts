// Interfaces are implemented implicitly.
// If a type contains all the methods of an interface, it automatically fulfills the interface.

package main

import (
	"fmt"
	"math"
)

type quadrilateral interface {
	area() float64
	perimeter() float64
	diagonal() float64
}

type square struct {
	side float64
}

type rect struct {
	length float64
	width  float64
}

type trapezius struct {
	base1  float64
	base2  float64
	height float64
	side1  float64
	side2  float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

func (s square) diagonal() float64 {
	return s.side * math.Sqrt(2)
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r rect) diagonal() float64 {
	return math.Sqrt((r.length * r.length) + (r.width * r.width))
}

func printData(q quadrilateral) {
	fmt.Println("Area: ", q.area())
	fmt.Println("Perimeter: ", q.perimeter())
	fmt.Println("Diagonal: ", q.diagonal())
}

func main() {
	r := rect{length: 10, width: 5}
	printData(r)
}
