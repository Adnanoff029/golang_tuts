package main

import (
	"fmt"
)

// Nested Struct
type manager struct {
	details         employee
	numberOfReports int
}

// Basic struct
type employee struct {
	name        string
	age         int
	salary      int
	desgination string
	department  string
	isWorking   bool
}

// Nested anonymous structs
// Description is an anonymous struct field
type animal struct {
	breed       string
	age         int
	description struct {
		color  string
		weight float64
		height float64
		lenth  float64
	}
}

// Embedded Structs (All the fields of the car struct are available at the top leve of the truck struct unlike nested structs)
type truck struct {
	car
	bedSize int
}

type car struct {
	brand string
	model string
}

func main() {
	// Create a new employee
	emp1 := employee{
		name:        "John Doe",
		age:         30,
		salary:      50000,
		desgination: "Software Engineer",
		department:  "Engineering",
		isWorking:   true,
	}
	manager1 := manager{
		details:         emp1,
		numberOfReports: 5,
	}
	dog := animal{
		breed: "German Shepherd",
		age:   5,
		description: struct {
			color  string
			weight float64
			height float64
			lenth  float64
		}{
			color:  "Black",
			weight: 30.5,
			height: 2.5,
			lenth:  3.5,
		},
	}
	fmt.Println("Hello, World! \n", manager1)
	fmt.Println("Hello, World! \n", emp1.desgination)
	fmt.Println(dog)
}
