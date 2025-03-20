// This is a simple example of channels being ranged over
package main

import "fmt"

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for range n {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	n := 10
	ch := make(chan int)
	go fibonacci(n, ch)
	for value := range ch {
		fmt.Println(value)
	}
}
