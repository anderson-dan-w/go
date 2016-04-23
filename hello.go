package main

import (
	"fmt"
	"math/rand"
	"time"
)

// can't : out here
// also, Caps variables get exported
var MyNumber = 37

// and lowercase don't
var myOtherNumber = 73

func square(x int) int {
	return x * x
}

// No named returns, so need NEW variables - the :
func squareAndCube(x int) (int, int) {
	y := square(x)
	z := x * y
	return y, z
}

// Named returns figure out what you probably meant to return
// and no need to "declare" them - already done in the func-decl
func squareAndQuad(x int) (y int, z int) {
	y = square(x)
	z = square(y)
	return
}

func main() {
	fmt.Println("Hello, Dan")
	fmt.Println("Time is: ", time.Now())
	fmt.Println("Random number is", rand.Intn(10))
	fmt.Println("My first function: square(9)=", square(9))

	// new local variables, so need to : them
	x, y := squareAndCube(9)
	fmt.Println("andCube", x, y)
	// overwriting previous values, so don't redeclare
	x, y = squareAndQuad(9)
	fmt.Println("andQuad", x, y)
}
