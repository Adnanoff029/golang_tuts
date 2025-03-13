// Array are fixed sized

// We use slices -> Dynamic sized, flexible view of elements of an array

// Slices hold references to an underlying array.

// Capacity of a slice is the number of elements in the underlying array.

// Syntax -> arrayname[startIndex:endIndex + 1] / arrayname[startIndex:] (complete array from startIndex) / arrayname[:endIndex] (array upto endIndex - 1) / arrayname[:] (full array)

// We also have slices package -> It has Contains(slice, value), Index(slice, value), Clone(slice), Reverse(slice), Sort(slice), SortFunc(slice, cmpFunc), Delete(slice, i, j), Equal(slice1, slice2).
package main

import (
	"fmt"
	"slices"
)

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) {
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}

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
	strArrTemp := strArr[:]
	for i := range strArrTemp {
		fmt.Printf("%v ", strArrTemp[i])
	}
	fmt.Printf("\n")

	// Making a slice make([]int, size, capacity)
	makedSlice := make([]int, 5) // The default values would be 0
	makedSlice[4] = 100
	fmt.Printf("The size of makedSlice is : %d \n", len(makedSlice))
	fmt.Printf("The capacity of the makedSlice is: %d \n", cap(makedSlice))
	for i, value := range makedSlice {
		fmt.Printf("For index %d : %d \n", i, value)
	}

	// Copying slices (see the output this is interesting)
	src := []int{1, 2, 3, 4, 5, 6}
	dest := []int{7, 8, 9}
	copy(src, dest)
	fmt.Println(dest, src)

	// Some tricks
	// 1. Reslicing a slice
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[:3]
	slice3 := slice2[1:]
	fmt.Println(slice1, slice2, slice3)
	// 2. Removing an element from the slice (in this case it is at index 2)
	// slice1 = append(slice1[:2], slice1[3:]...)
	slice1 = slices.Delete(slice1, 2, 3)
	fmt.Printf("This is the slice1 with element at index 2 removed : %v \n", slice1)

	arr2 := []string{"Adnan", "Saquib", "Sufia", "Muzib"}
	slice4 := arr2[:]
	slice5 := arr2[:]
	slice4[0] = "Khan"
	fmt.Println(slice4, slice5)

	twoDArr := [3][3]int{{1, 2, 3}, {3, 4, 5}, {6, 7, 8}}
	for i := range twoDArr {
		for j := range twoDArr[i] {
			fmt.Printf("%d ", twoDArr[i][j])
		}
		fmt.Printf("\n")
	}

	twoDArr2 := twoDArr[0:2]
	twoDArr2[0][0] = 100
	fmt.Println(cap(twoDArr2)) // It returns the capacity of the underlying array of the slice unlike len which returns the size of current slice.

}
