package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func mult(x int, y int) int {
	return x * y
}

func wrapTwice(fn func(int, int) int, x int, y int) int {
	return fn(fn(x, y), fn(x, y))
}

// closures: return made func that reference outside vars
// adder takes no arge, and returns a func of form f(int)->int
func adder() func(int) int {
	// sum is not local to the func we are returning, but it's bound
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// a closure that takes a func as a param, cause why not
// who cares if you can read it or reason about it
// this is a function, that takes a func like mult or add, and an initial value
// and returns a function that will apply the func to value and the new param,
// updating the result as it goes. Sort of like a weirdly contrived, can stop
// in the middle, reduce
func doFunc(fn func(int, int) int, initial int) func(int) int {
	val := initial
	return func(x int) int {
		val = fn(val, x)
		return val
	}
}

func main() {
	var x int = 4
	var y int = 3
	fmt.Println("Add ", x, y)
	fmt.Println(add(x, y))
	fmt.Println(" ... and twice")
	fmt.Println(wrapTwice(add, x, y))

	fmt.Println("Mult ", x, y)
	fmt.Println(mult(x, y))
	fmt.Println("... and twice")
	fmt.Println(wrapTwice(mult, x, y))

	every, evens := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(every(i), evens(2*i))
	}

	expon := doFunc(mult, 1)
	for i := 1; i < 5; i++ {
		fmt.Println(expon(i))
	}
}
