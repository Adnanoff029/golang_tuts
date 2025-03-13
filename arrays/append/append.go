// builtin append() dynamically adds elements to the slice

// append is a variadic function

// during append, new array is only created if the capacity of the slice falls short.

// Note -> Use append on the same slice the result is assigned to.

package main

import "fmt"

// Custom append function for interger arrays
func append(slice []int, args ...int) []int {
	ans := make([]int, (len(slice) + len(args)))
	for i, value := range slice {
		ans[i] = value
	}
	for i, value := range args {
		ans[len(slice)+i] = value
	}
	return ans
}

func main() {
	slice := make([]int, 5)
	fmt.Println(slice)
	slice = append(slice, 111, 111, 111)
	slice2 := make([]int, 3)
	slice = append(slice, slice2...)
	fmt.Println(slice)

	i := make([]int, 3, 8)
	fmt.Println("len of i:", len(i))
	fmt.Println("cap of i:", cap(i))
	// len of i: 3
	// cap of i: 8

	j := append(i, 4)
	fmt.Println("appending 4 to j from i")
	fmt.Println("j:", j)
	fmt.Println("addr of j:", &j[0])
	// appending 4 to j from i
	// j: [0 0 0 4]
	// addr of j: 0x454000

	g := append(i, 5)
	fmt.Println("appending 5 to g from i")
	fmt.Println("addr of g:", &g[0])
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("g:", g)
	// appending 5 to g from i
	// addr of g: 0x454000
	// i: [0 0 0]
	// j: [0 0 0 5]
	// g: [0 0 0 5]
}
