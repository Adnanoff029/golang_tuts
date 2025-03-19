// Concurrency Vs Parallelism
// Parallelism means that all the tasks are being carried out at the same time due to multiple cores. Each core assigned with a particular task.
// Concurrency is the illusion of parallelism.

// Go is designed to be concurrent.
// We use "go" keyword is used to spawn a new goroutine and run code faster with concurrency.
// A Goroutine is a lightweight thread of execution managed by the Go runtime. It allows concurrent execution without creating OS-level threads, making Go highly efficient in handling multiple tasks.

// Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
// Watchout for deadlocks when working with goroutines.

package main

import (
	"fmt"
)

var o int

func concurrentFunc(ch chan int) {
	v := <-ch
	fmt.Println("This is a concurrent function:", v)
}

func helper(ch chan int) {
	for i := range 10 {
		ch <- i
	}
}

func helper2(ch chan string, mp *map[int]string) {
	str := <-ch
	fmt.Println(str)
	(*mp)[o] = str
	o++
}

func help() {
	// Making a channel
	ch := make(chan int)
	ch2 := make(chan string)
	// Send data to channel (this will block until another goroutine is ready to receive the value.)
	mp := make(map[int]string)
	mpPtr := &mp
	go helper(ch)
	var i int
	for i < 10 {
		concurrentFunc(ch)
		i++
	}
	i = 0
	go helper2(ch2, mpPtr)
	for i < 10 {
		ch2 <- "This is a string"
		i++
	}
	for key, value := range mp {
		fmt.Println(key, value)
	}
	// Receive data from channel (reads and removes a value from the channel and saves it into a variable v. Operation will block until the value is read.)
	// v := <-ch
	// sometimes we don't care what is passed through a channel. We care when and if it is passed. Hence we use empty channels.
	// <- ch will block until it pops a single item off the channel, then continue discardin the item.

}

func main() {
	help()
}
