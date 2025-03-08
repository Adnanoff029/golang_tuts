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

// Empty Struct
type emptyStruct struct{}

// Method for employee struct
func (emp employee) getDetails() {
	fmt.Println("Employee Name: ", emp.name)
	fmt.Println("Employee Age: ", emp.age)
	fmt.Println("Employee Salary: ", emp.salary)
	fmt.Println("Employee Desgination: ", emp.desgination)
	fmt.Println("Employee Department: ", emp.department)
	fmt.Println("Employee isWorking: ", emp.isWorking)
}

// Method for truck struct
func (t truck) getDetails() {
	fmt.Println("Brand: ", t.brand)
	fmt.Println("Model: ", t.model)
	fmt.Println("Bed Size: ", t.bedSize)
}

// Method for User struct
func (u User) getDetails() {
	fmt.Println("Name: ", u.name)
	fmt.Println("Age: ", u.age)
	fmt.Println("Membership Type: ", u.Type)
	fmt.Println("Message Char Limit: ", u.MessageCharLimit)
}

func (u User) checkMessage() bool {
	if u.MessageCharLimit > 5000 {
		return true
	}
	return false
}

type Membership struct {
	Type             string
	MessageCharLimit int64
}

type User struct {
	name string
	age  int
	Membership
}

func main() {
	// empty struct -> 0 byte
	// bool -> 1 byte
	// uint16 -> 2 bytes
	// int64 -> 8 bytes

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
		details: employee{
			name:        "John Doe",
			age:         30,
			salary:      50000,
			desgination: "Software Engineer",
			department:  "Engineering",
			isWorking:   true,
		},
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
	trk := truck{
		car: car{
			brand: "Ford",
			model: "F150",
		},
		bedSize: 5,
	}
	empty := emptyStruct{}
	user1 := User{
		name: "John Doe",
		age:  30,
		Membership: Membership{
			Type:             "Standard",
			MessageCharLimit: 5000,
		},
	}
	fmt.Println("User")
	user1.getDetails()
	fmt.Println("Valid user Message Limit:", user1.checkMessage())
	fmt.Println("Empty Struct \n", empty)
	fmt.Println("Truck \n", trk, trk.model)
	fmt.Println("Manager \n", manager1.details.department)
	fmt.Println("Animal \n", dog)
	fmt.Println("Employee")
	emp1.getDetails()
	fmt.Println("Truck")
	trk.getDetails()
}
