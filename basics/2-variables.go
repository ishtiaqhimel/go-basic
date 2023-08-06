package main

import "fmt"

// Global variable
var X int
var Y int = 7

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

/*
Constant in Go
Constants are the fixed values that cannot be changed once declared.
In Go, we use the const keyword to create constant variables. For example,

initial Value:
const employeeId = 10028

Error! Constants cannot be changed:
emloyeeId = 10028

By the way, we cannot use the shorthand notation := to create constants. For example,
Error! Code
const employee := 10028
*/

func main() {
	fmt.Println(X, Y)

	// Declaring a single variable:
	// declarring first, then assigning value, default is always zero
	var i int
	i = 5

	// Declaring a variable with an initial value:
	// we're declaring a variable & assigning the value at the same time
	var j int = 10

	// Type inference:
	// if we don't mention the type, then it'll see the value & assign the type, string for this instance
	var name = "Ishtiaq"
	fmt.Println(i, j, name)

	// shorthand for declarign & assigning values to a new variable
	k := 15
	// we can't re-declare it like this: k := 10, what we can do is assign a new value to it like this: k = 10

	fmt.Println(k)
}
