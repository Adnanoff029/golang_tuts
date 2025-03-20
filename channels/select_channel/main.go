package main

import (
	"fmt"
	"time"
)

func a(ch chan int) {
	ch <- 1
	close(ch)
}

func b(ch chan int) {
	ch <- 2
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go a(ch1)
	go b(ch2)
	time.Sleep(time.Millisecond * 1000)
	for value := range ch1 {
		fmt.Println("Received from ch1:", value)
	}
	for value := range ch2 {
		fmt.Println("Received from ch2:", value)
	}
	for range 2 {
		select {
		case value1, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", value1)
			} else {
				fmt.Println("Channel 1 is closed")
			}
		case value2, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", value2)
			} else {
				fmt.Println("Channel 2 is closed")
			}
		default:
			fmt.Println("No channel is ready")
		}
	}
}
