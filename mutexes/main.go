// Mutexes are used to lock and unlock the resources.
// It can be used to limit the access of data to a particular goroutine.
// Maps are not thread safe, so we need to use mutexes to lock and unlock the access of map so that only one goroutine can access the map at a time.
package main

import (
	"fmt"
	"sync"
)

var mu *sync.Mutex

func mutexFunc() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Mutex locked")
}

func a() {
	fmt.Println("a")
	mutexFunc()
	fmt.Println("a done")
}

func b() {
	fmt.Println("b")
	mutexFunc()
	fmt.Println("b done")
}

func main() {
	go a()
	go b()
}
