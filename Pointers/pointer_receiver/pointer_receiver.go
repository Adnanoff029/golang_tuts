package main

import "fmt"

type circle struct {
	x      int
	y      int
	radius int
}

func (c *circle) initCircle(cx, cy, r int) {
	c.x = cx
	c.y = cy
	c.radius = r
}

func (c *circle) translateCircle(i, j int) {
	c.x += i
	c.y += j
}

func main() {
	tempCircle := circle{}
	fmt.Println(tempCircle)
	tempCircle.initCircle(0, 0, 5)
	fmt.Println(tempCircle)
	tempCircle.translateCircle(4, 9)
	fmt.Println(tempCircle)
}
