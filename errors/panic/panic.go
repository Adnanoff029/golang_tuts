// Panic is another way of handling errors in go
// When the panic function is called, then the program crashes and prints a stack trace
// Note -> It is advised not to do this. It yeets control out of current function and up the call stack until it reaches a function that defers a recover.

// Defer

// Defer keyword pushes the function into a list that executes then after the surrounding function returns int a LIFO manner.
// Deferred function arguments are evaluated when the defer statement is evaluated.
// Deferred function calls are executed in a LIFO manner.

// If there is no panic, then recover simply returns nil.
// Panic is a built in function that stops the normal flow of program. If function f() calls panic, then the execution of f() stops and all the defers of f() are executed normally and then f() returns to its caller. f() then behaves as a call to panic and keeps returning to subsequent callers until the entire goroutine has returned.
// Recover is a built in function that regains control of a panicking goroutine. It is used inside deferred functions to stop the goroutine from crashing if a panic is called.

package main

import "fmt"

func a() {
	i := 1
	fmt.Println("This is a.")
	defer fmt.Println(2 * i)
	defer func() {
		i++
		fmt.Println("This is a -> b.")
		defer func() {
			fmt.Println(2 * i)
			fmt.Println("This is a -> b -> 1")
		}()
		fmt.Println("Exiting from a -> b")
	}()
	defer func() {
		// fmt.Println(2 * i)
		fmt.Println("This is a -> c.")
	}()
	i++
	fmt.Println(i)
	fmt.Println("Exiting from a.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f")
		}
	}()
	fmt.Println("Calling g()")
	g(0)
	fmt.Println("Returned normally from g()")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!!")
		panic(fmt.Sprintf("Panicked because i > 3 : i = %v", i))
	}
	defer fmt.Println("Defer in g()", i)
	fmt.Println("Printing in g()", i)
	g(i + 1)
}

func main() {
	a()
	fmt.Printf("\n \n")
	fmt.Println("Calling f()")
	f()
	fmt.Println("Returned normally from f()")
}
