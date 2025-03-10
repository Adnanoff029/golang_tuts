// Error is also an interface in go
package main

import (
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
	return err.response
}

type notFoundError struct {
	code    int
	message string
}

func (err notFoundError) Error() string {
	return err.message
}

func printError(err error) {
	switch v := err.(type) {
	case notFoundError:
		fmt.Println(v.message)
	case badGatewayError:
		fmt.Println(v.response)
	default:
		fmt.Println("Unknown error")
	}
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
	printError(err1)
	printError(err2)

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
}
