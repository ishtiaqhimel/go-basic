// this is a single line comment
/*
	this a multiple line comment
*/

// package declaration
// packages are Go's way of organizing and reusing code
package main

// import keyword is how we include code from other packages to use with our program
import "fmt"

// function declaration
// 'main' is special because it's the function that gets called when you execute the program
func main() {
	// access another function inside of the fmt package called Println
	fmt.Println("Hello World")
}
