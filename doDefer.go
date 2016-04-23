package main

import "fmt"

func sayHi() string {
	fmt.Println("inside sayHi()")
	return "sayHi() is returning now"
}

// it'll evaluate sayHi() - and hence print Hi
// then print Inside doSomeStuff
// then, upon function-finish, it'll print return of sayHi()
func doSomeStuff() {
	defer fmt.Println(sayHi())
	defer fmt.Println("This prints before sayHi() return value - LIFO")
	fmt.Println("Think of defer along the lines of 'finally'")
}

func main() {
	doSomeStuff()
}
