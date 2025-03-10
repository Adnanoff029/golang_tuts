// Type assertions are used to access the underlying type of an interface value

package main

import "fmt"

type j interface {
	interfaceMethod()
}

type x struct {
	property string
	luck     int
}

func (va x) interfaceMethod() {
	fmt.Println("This is the interface method")
}

// We can also use switch to do multiple type assertions
func checkType(c any) {
	switch v := c.(type) {
	case x:
		fmt.Println("This is a struct x")
		fmt.Println(c.(x).property)
		fmt.Println(v)
	case int:
		fmt.Println("This is an integer")
		fmt.Println(v)
	case string:
		fmt.Println("This is a string")
		fmt.Println(v)
	default:
		fmt.Println("This is not a struct x")
	}
}

func main() {
	// var c any = "This is a string"
	// d, ok := c.(string) // If c is a string then ok is true, else it is false
	// if ok {
	// 	fmt.Println(d)
	// } else {
	// 	fmt.Println("This is okay")
	// }

	// We declared an interface j and a struct x
	var y any
	y = x{
		property: "This is a property",
		luck:     10,
	}

	c, ok := y.(x)
	if ok {
		fmt.Printf("%T, %v \n", c, c)
		fmt.Println(c.property, c.luck)
		c.interfaceMethod()
	} else {
		fmt.Println("This is not okay")
	}

	checkType(10)
	checkType("Adnan")
	checkType(c)
	checkType(x{})
}
