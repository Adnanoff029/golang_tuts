// Interfaces are custom types that other types can implement.
// Interfaces are implemented implicitly.
// If a type contains all the methods of an interface, it automatically fulfills that interface.
// A type can have multiple interfaces (below the square and rectangles have 2 interfaces -> dimensions and quadrilateral)
// We can add names to the interface methods' parameters.
// Note - Errors in go are also interfaces

// Best Practices
// 1. Keep thee small
// 2. Interfaces should have no knowledge of satisfying types
// 3. Interfaces are not classes (no constructor or destructor, won't DRY the code as we still need to define custom methods with same same for each struct, however they would be accessible with one function)

package main

import (
	"fmt"
	"math"
)

type scale interface {
	scaleX(scaleParameter float64) (xDimension float64)
	scaleY(scaleParameter float64) (yDimension float64)
	scaleXY(scaleParameter float64) (xDimension float64, yDimensions float64)
}

type dimensions interface {
	getDimensions() (float64, float64)
}

// A quadrilateral interface
type quadrilateral interface {
	area() float64
	perimeter() float64
	diagonal() float64
}

// A square struct
type square struct {
	side float64
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

func (s square) getDimensions() (float64, float64) {
	return s.side, s.side
}

// A rect struct
type rect struct {
	length float64
	width  float64
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

func (r rect) getDimensions() (float64, float64) {
	return r.length, r.width
}

func (r rect) scaleX(scaleParameter float64) (xDimension float64) {
	return r.length * scaleParameter
}

func (r rect) scaleY(scaleParameter float64) (yDimension float64) {
	return r.width * scaleParameter
}

func (r rect) scaleXY(scaleParameter float64) (xDimension float64, yDimensions float64) {
	return r.length * scaleParameter, r.width * scaleParameter
}

// Functions that validate that the interfaces have been automatically assigned. They take the interface as a parameter.

func printData(q quadrilateral) {
	fmt.Println("Area: ", q.area())
	fmt.Println("Perimeter: ", q.perimeter())
	fmt.Println("Diagonal: ", q.diagonal())
}

func scaleData(s scale) {
	fmt.Printf("\n")
	fmt.Println("Scaled X: ", s.scaleX(5.0))
	fmt.Println("Scaled Y: ", s.scaleY(5.0))
	a, b := s.scaleXY(5.0)
	fmt.Printf("Scaled XY: %v %v", a, b)
}

func printDimensions(d dimensions) {
	l, w := d.getDimensions()
	fmt.Println("Length: ", l)
	fmt.Println("Width: ", w)
}

func main() {
	r := rect{length: 10, width: 5}
	s := square{side: 5}
	fmt.Printf("Below are the details of the square\n")
	printData(s)
	printDimensions(s)
	fmt.Printf("\n")
	fmt.Printf("Below are the details of the rectangle\n")
	printData(r)
	printDimensions(r)
	fmt.Printf("I am scaling the rectangle by 5 \n")
	scaleData(r)
}
