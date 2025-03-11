// Error is also an interface in go
// Go has a standard library package called errors for error creation, comparison and type assertions.
// Error handling in go is more controlled rather than JS, where all errors of a block are handled at a single place

package main

import (
	"errors"
	"fmt"
	"strconv"
)

type error interface {
	Error() string
}

type badGatewayError struct {
	code     int
	response string
}

func (err badGatewayError) Error() string {
	return fmt.Sprintf("%v BAD GATEWAY : %v", err.code, err.response)
}

type notFoundError struct {
	code    int
	message string
}

func (err notFoundError) Error() string {
	return fmt.Sprintf("%v NOT FOUND : %v", err.code, err.message)
}

func printError(err error) error {
	switch v := err.(type) {
	case notFoundError:
		return v
	case badGatewayError:
		return v
	default:
		return fmt.Errorf("Unknown error")
	}
}

func getUserId() (int, error) {
	// This is a dummy function that returns an integer and an error
	return 0, nil
	// return -1, fmt.Errorf("This is an error")
}

func divide(l, m int) (float64, error) {
	if m == 0 {
		return -1, errors.New("Divisor can't be 0")
	}
	return float64(l / m), nil
}

func main() {
	err1 := notFoundError{
		code:    404,
		message: "Not found",
	}
	err2 := badGatewayError{
		code:     500,
		response: "Bad Gateway",
	}
	fmt.Println(printError(err1))
	fmt.Println(printError(err2))
	var a string
	a = "101"
	// Atoi converts a string to an integer and returns 2 values the integer and error
	// If the string is convertible then the error is nil
	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("This is not an integer")
		fmt.Println(err)
	} else {
		fmt.Println("This is an integer", b)
	}

	x, otherwise := getUserId()
	if otherwise != nil {
		fmt.Println("This is an error")
	} else {
		fmt.Println("This is the user id", x)
	}
	// errors package
	// errors.New(msg string) error
	temp := errors.New("This is an error made of errors package")
	fmt.Printf("%T, %v \n", temp, temp)
	// errors.Is(err, target error) bool
	// For error specific comparisons
	temp2 := errors.Is(temp, errors.New("This is an error made of errors package"))
	fmt.Printf("%v \n", temp2)
	// errors.As(err, target interface{}) bool Checks if the err can be cast into the target type
	// For error types
	var temp3 error
	temp4 := errors.As(temp, &temp3)
	fmt.Printf("%v \n", temp4)

	l := 10
	m := 1
	// m := 0 // (returns an error)
	c, err10 := divide(l, m)
	if err10 != nil {
		fmt.Println(err10)
	} else {
		fmt.Println(c)
	}
}
