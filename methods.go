package main

import (
	"fmt"
	"math"
)

// this exists in types.go but I can't figure out how to import it
type SpaceCoord struct {
	x float64
	y float64
	z float64
}

// this "wierd syntax" means first supply the SpaceCoord, and then .Distance()
// to call it, just like a "class method" even though go doesn't have classes
func (s SpaceCoord) Distance() float64 {
	x := s.x*s.x + s.y*s.y + s.z*s.z
	return math.Sqrt(x)
}

func main() {
	fmt.Println("Methods")
	flatZ := SpaceCoord{x: 3, y: 4}
	fmt.Println(flatZ.Distance())
	s := SpaceCoord{6, 6, 7}
	fmt.Println(s.Distance())
}
