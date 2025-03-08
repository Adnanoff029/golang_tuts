package main

import ("math" 
		"fmt")

type quadrilateral interface {
	area() float64
	perimeter() float64
	diagonal() float64
}

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

func printData(q quadrilateral) {
	fmt.Println("Area: ", q.area())
	fmt.Println("Perimeter: ", q.perimeter())
	fmt.Println("Diagonal: ", q.diagonal())
}

func main() {
	r := rect{length: 10, width: 5}
	printData(r)
}
