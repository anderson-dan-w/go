package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello, Dan")
	fmt.Println("Time is: ", time.Now())
	fmt.Println("Random number is", rand.Intn(10))
}
