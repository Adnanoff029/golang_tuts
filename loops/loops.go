// There is no while loop in go because we can omit init and after sections in for loop, we can make it work like a while loop also.
// We also have continue and break statements in go
package main

import "fmt"

func main() {
	name := "Adnan Khan"
	// Old C-like syntax of a "For loop"
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(name)
	// }

	// Simulating a while loop
	i := 0
	for i < 5 {
		if i == 3 {
			i++
			continue
		}
		fmt.Println(name, i)
		i++
	}
}
