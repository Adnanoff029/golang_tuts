// Concept of functional programming langugages that involves partial application of functions.

// It allows function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.

package main

import "fmt"

// Takes a function as argument and returns a function as argument
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	a := squareFunc(10)
	fmt.Println(a)
	b := doubleFunc(10)
	fmt.Println(b)
}
