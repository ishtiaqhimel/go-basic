package main

import "fmt"

// we can move the variable outside of the main function
// This means that other functions can access this variable
// var x string = "Hello World"

/*
we can enclose them in parenthesis like this.
& ofcourse follow better naming convention, don't just put x, y etc as variable name
use meaningful & considerable lenghth considering lifecycle of that particular variables
*/
var (
	firstName string = "Ishtiaq"
	lastName  string = "Islam"
	company   string = "appscode"
	id        int    = 10028
)

func main() {
	// declarring first, then assigning value, default is always zero
	var i int
	i = 5

	// we're declaring a variable & assigning the value at the same time
	var j int = 10

	// if we don't mention the type, then it'll see the value & assign the type, string for this instance
	var name = "Ishtiaq"
	fmt.Println(i, j, name)

	// shorthand for declarign & assigning values to a new variable
	k := 15
	// we can't re-declare it like this: k := 10, what we can do is assign a new value to it like this: k = 10

	fmt.Println(k)
}
