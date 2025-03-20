// Concurrency Vs Parallelism
// Parallelism means that all the tasks are being carried out at the same time due to multiple cores. Each core assigned with a particular task.
// Concurrency is the illusion of parallelism.

// Go is designed to be concurrent.
// We use "go" keyword is used to spawn a new goroutine and run code faster with concurrency.
// A Goroutine is a lightweight thread of execution managed by the Go runtime. It allows concurrent execution without creating OS-level threads, making Go highly efficient in handling multiple tasks.

// Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.
// Watchout for deadlocks when working with goroutines.

// Some points to remember:
// 1. A declared but uninitialized channel is nil.
// 2. A send to a nil channel blocks forever.
// 3. A receive from a nil channel blocks forever.
// 4. Send to close channel will panic.
// 5. Receive from closed channel returns zero value immediatey.


package main

import (
	"fmt"
	"time"
)

// producer demonstrates sending values on an unbuffered channel
func producer(ch chan int) {
	for i := range 3 {
		fmt.Printf("Sending: %d\n", i)
		ch <- i // This will block until receiver is ready
		time.Sleep(time.Millisecond * 1000)
	}
}

// bufferProducer demonstrates sending values to a buffered channel
func bufferProducer(ch chan int) {
	// These won't block until buffer is full
	ch <- 10
	ch <- 20
}

// sendOnly demonstrates a channel that can only send
func sendOnly(ch chan int) {
	ch <- 42
}

// receiveOnly demonstrates a channel that can only receive
func receiveOnly(ch chan int) {
	val := <-ch
	fmt.Printf("Received in receiveOnly: %d\n", val)
}

// selectExample demonstrates the select statement with multiple channels
func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 1000)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 2000)
		ch2 <- "message from ch2"
	}()

	// Select statement with timeout
	fmt.Println("\nSelect statement example:")
	for range 2 {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Millisecond * 3000):
			fmt.Println("Timeout!")
			return
		}
	}
}

func main() {
	// Declare an unbuffered channel of integers
	// Unbuffered channels block until both sender and receiver are ready
	ch := make(chan int)

	// Declare a buffered channel with capacity of 2
	// Buffered channels block only when buffer is full (for sends) or empty (for receives)
	bufferedCh := make(chan int, 2)

	// Start a goroutine to demonstrate channel communication
	go producer(ch)
	go bufferProducer(bufferedCh)

	// Receive values from unbuffered channel (recieved from producer())
	fmt.Println("Receiving from unbuffered channel:")
	for range 3 {
		value := <-ch
		fmt.Printf("Received: %d\n", value)
	}
	close(ch) // This will close the channel

	// Demonstrate buffered channel operations (recieved from bufferProducer())
	fmt.Println("\nBuffered channel operations:")
	for range 2 {
		value := <-bufferedCh
		fmt.Printf("Received from buffer: %d\n", value)
	}
	close(bufferedCh) // This will close the channel
	// bufferedCh <- 30 // This will panic because the channel is closed
	v, ok := <-bufferedCh // 0, false because the channel is closed
	fmt.Printf("Checking if bufferedCh is closed: %d, %t\n", v, ok)

	// Channel direction example (recieved from sendOnly() and receiveOnly())
	sendOnlyCh := make(chan int)
	go sendOnly(sendOnlyCh)    // Write only channel
	go receiveOnly(sendOnlyCh) // Read only channel

	// Select statement example with timeout
	selectExample()
}
