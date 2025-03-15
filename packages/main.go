// Here we declared a package "main" whose entry point is the main() function.
// On running this, the main package is compiled into an executable program.
// Packages from any other name are library packages like "fmt" or "slices" or "math" etc.
// By convention package's name is the same as the last element of its import path. Like -> "math/rand"
// A directory of go can have atmost one package i.e. all go files in a folder must belong to the same package.
// Path to a package's directory determines its import path and where it can be downloaded from.
// $GOPATH env. variable is set by default on the machne
package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.111
	fmt.Println(math.Round(a))
}
