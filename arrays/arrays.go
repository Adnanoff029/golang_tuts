// Array are fixed sized
// We use slices -> Dynamic sized, flexible view of elements of an array
package main

import "fmt"

func main() {
	var arr []int

	for i := range 10 {
		arr = append(arr, i)
	}

	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
	fmt.Printf("\n")
	strArr := [3]string{"one", "two", "three"}
	fmt.Println(strArr[0])
}
