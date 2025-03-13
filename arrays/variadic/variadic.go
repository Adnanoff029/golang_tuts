// VARIADIC FUNCTIONS
// Many functions take arbitrary number of final arguments which is facilitated by "..."

// Println() and Sprintf() are also variadic functions.
package main

import "fmt"

// Custom variadic function using spread operator
func concat(strs ...string) string {
	final := ""
	for _, value := range strs {
		final += value
	}
	return final
}

func main() {
	ans := concat("Apple", "Ball", "Cat", "Dog")
	n, err := fmt.Println(ans)
	fmt.Println(n, err)
}
