package main

import "fmt"

type spacecoord struct {
	x int
	y int
	z int
}

func main() {
	origin := spacecoord{0, 0, 0}
	lazyOrigin := spacecoord{}
	fmt.Println(origin, lazyOrigin)

	pythagoras3d := spacecoord{3, 4, 5}
	// create pointers like C ...
	p := &pythagoras3d
	// ... but pointer dereferencing is ALSO .
	fmt.Printf("x:%v y:%v z:%v (%T)\n", pythagoras3d.x, p.y, p.z, p)

	// do you...
	noY := spacecoord{x: 1, z: 1}
	fmt.Println("Do you know why?", noY)

	// onto arrays
	xs := [4]int{1, 2, 3, 4}
	var ys []int = xs[0:3]
	fmt.Printf("xs: %v and ys: %v\n", xs, ys)

	fmt.Println("let's change y[0]")
	ys[0] = 4
	fmt.Printf("xs: %v and ys: %v\n", xs, ys)

	// dynamic arrays, len, capacity
	a := make([]int, 5) // len=5, cap=5
	printSlice(a)
	a = a[:3] // len=3, cap still 5
	printSlice(a)
	a = a[:cap(a)] // back to len=5, cap=5
	//a = a[:cap(a)+1] // out of bounds
	a = a[2:] // len=3, cap=3 -- we chucked the first two, can't get back
	printSlice(a)
	a = append(a, 4, 5) // len=5 but ...cap=6? dunno why yet..
	printSlice(a)

	for i, x := range a {
		fmt.Println(i, x)
	}

	xy := make([][]uint8, 10)
	for i := range xy {
		xy[i] = make([]uint8, 10)
		for j := range xy[i] {
			xy[i][j] = uint8(i) * uint8(j)
		}
	}
	for _, row := range xy {
		fmt.Println(row)
	}

	locations := make(map[string]spacecoord)
	locations["origin"] = origin
	locations["me"] = *p
	for k, v := range locations {
		fmt.Println(k, v)
	}
	fmt.Println(locations)

	v, ok := locations["here"]
	fmt.Println("The value:", v, "Present?", ok)
}

func printSlice(a []int) {
	fmt.Printf("a:%v len(%v) cap(%v)\n", a, len(a), cap(a))
}
