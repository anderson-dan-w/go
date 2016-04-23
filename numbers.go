package main

import (
	"fmt"
	"math"
)

// can do any block-like operation in block-fmt
// and also, type how you want; gofmt will neaten it up
// (and vim-go does it automatically on save!)
var (
	yes         = true
	f32 float32 = 0.5
	f64 float64 = 0.5
)

// apparently can't overload function names: try making it "square"
// and compilation fails
func squareInt(x int) (y int) {
	return x * x
}

func square(x float64) float64 {
	return x * x
}

func newtonSqrt_10(x float64) float64 {
	z := float64(x / 2) // randomish starting point
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func newtonSqrt_delta(x float64) (z float64) {
	z = x
	newz := z / 2
	for math.Abs(z-newz) > square(.00001) {
		z = newz
		newz = z - ((z*z - x) / (2 * z))
	}
	return
}

func main() {
	// can't embrace roll over to be lazy:doesn't compile
	//var max8 uint8 = -1

	// but I guess it makes types AFTER the math?
	var max8 uint8 = 1<<8 - 1
	fmt.Println(max8)
	fmt.Println(yes, f32, f64)

	// based on running this, looks like truncation, or floor
	casted := int(f32)

	// v=value; q=quoted; T=Type;
	// seems lame to write varname 3 times, not sure if anything better yet
	fmt.Printf("I cast ye out! %v (quoted=%q) (type=%T)\n",
		casted, casted, casted)

	// as generally expected: [start, stop) : includes 0, not 10
	for i := 0; i < 10; i++ {
		fmt.Println(squareInt(i))
	}

	sum := 1
	// fun fact: I typed "for ; sum < 100 ; {"
	// and it reformatted to drop the semicolons
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum) // should print 128

	for i := 0; i < 10; i++ {
		x := float64(i)
		fmt.Printf("%v\n\tsqrt: %v\n\tnewton_10: %v\n\tnewton_delta: %v\n",
			i, math.Sqrt(x), newtonSqrt_10(x), newtonSqrt_delta(x))
	}
}
