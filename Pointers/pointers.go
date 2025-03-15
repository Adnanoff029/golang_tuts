// Pointers are variables that basically store memory address of some other variable.
// * gives the value present at the address the pointer is pointing to.
// & gives the address of current variable
// In go, functions are called by value, except for types like slices.
// Dereferencing a nil pointer will cause the program to panic
// Methods with pointer receiver don't need a pointer of that type to call it.
// Only use pointers if you have a need, use it when performance optimization is extremely important. Otherwise use simple values.
package main

import "fmt"

func help(s *string) {
	if len(*s) > 100 {
		return
	}
	*s += "\nThis was added using pointer."
	help(s)
}

func handleSlice(slice []int) {
	slice[0] = 100
}

// Pass by refernece
func randomFunc(rs *randomStruct) {
	// Below is the shorthand for (*rs).a = "I changed it using refernce!"
	rs.a = "I changed it using refernce!"
}

// Pointers in struct
type node struct {
	val  int
	next *node
	prev *node
}

// Method with pointers as receivers are able to modify its values.
func (n *node) getDetails() {
	fmt.Println(n.val, n.prev, n.next)
}

type randomStruct struct {
	a string
}

func main() {
	// Pointer using var
	var p *int
	var x int
	x = 100
	p = &x
	fmt.Println("Address :", p, "Pointing to value :", *p)

	// Declaring and assigning pointers
	p2 := &x
	fmt.Println("Pointing to same location as p :", p2)

	// Modifying values via pointer
	*p2 = 20
	fmt.Println(*p, *p2)

	// Simulating pass by refernce using pointers as there is only pass by value in go.
	s := "My name is Adnan Khan."
	p3 := &s
	help(p3)
	fmt.Println(s)

	// Using new operator
	p4 := new(node)
	p4.val = 100
	fmt.Println(*p4)

	// Just a simple doubly linked list
	head := node{
		val:  10,
		next: nil,
		prev: nil,
	}
	leaf := node{
		val:  20,
		next: nil,
		prev: nil,
	}
	head.next = &leaf
	leaf.prev = &head
	temp := &head
	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.next
	}
	head.getDetails()

	// Pointer to pointer
	y := "Value"
	p5 := &y
	pp5 := &p5
	fmt.Println(y, &y, p5, *p5, &p5, pp5, *pp5, &pp5)

	// Slices act as pass by reference
	randomSlice := make([]int, 5) // {0, 0, 0, 0, 0}
	randomSlice[0] = -1
	randomSlice[1] = -20
	randomSlice[2] = -300
	fmt.Printf("Before : %v \n", randomSlice)
	handleSlice(randomSlice)
	fmt.Printf("After : %v \n", randomSlice)

	// Passing struct by reference
	rs := randomStruct{
		a: "This is now",
	}
	fmt.Printf("Before : %v \n", rs)
	randomFunc(&rs)
	fmt.Printf("After : %v \n", rs)

}
