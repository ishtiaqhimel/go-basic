package main

import (
	"fmt"
	"os"
)

func printName(s string) {
	fmt.Println("My name is", s)
}

// func rectangleProperties(length, width float64) (float64, float64) {
// 	area := length * width
// 	perimeter := (length + width) * 2
// 	return area, perimeter
// }

func rectangleProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // also called naked return
}

// variadic can only be used in final parameter
func sum(nums ...int) int {
	// here num is []int type
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

var (
	prod = func(a, b int) int {
		return a * b
	}
)

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {
	fmt.Println("Fucntions")

	// sample function
	printName("Ishtiaq")

	// multiple return values
	area, perimeter := rectangleProperties(5.9, 2.3)
	fmt.Println(area, perimeter)

	// Variadic functions can be called with any number of trailing arguments.
	// For example, fmt.Println is a common variadic function.
	nums := []int{2, 5, 1, 3}
	s := sum(nums...)
	fmt.Println(s)

	// anonymous functions
	// An anonymous function is a function that was declared without any named identifier to refer to it.
	// Anonymous functions can accept inputs and return outputs, just as standard functions do.
	fmt.Println(prod(6, 8))

	// or assign result to a variable
	res := func(a, b int) int {
		return a * b
	}(8, 9)
	fmt.Println(res)

	// closures functions
	// Closures are a special case of anonymous functions. Closures are anonymous functions which access the
	// variables defined outside the body of the function.
	l, w := 20, 30
	func() {
		area := l * w // can't access area out of this function
		fmt.Println(area)
	}()

	// higher order functions
	// Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.
	partial := partialSum(3)
	fmt.Println(partial(7))

	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	// recursion
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))

	// Defer : defer schedules a function call to be run after the function completes. Consider the following example:
	defer second()
	first()
	// This program prints 1st followed by 2nd. Basically defer moves the call to second to the end of the function:

	// defer is often used when resources need to be freed in some way. For example when we open a file we need to make sure to close it later. With defer:
	myFile, err := os.Open("myFile.txt")
	defer myFile.Close()
	if err != nil {
		panic(err) //  the panic function basically causes a run time error
	}
	// This has 3 advantages:
	// (1) it keeps our Close call near our Open call so it's easier to understand,
	// (2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them and
	// (3) deferred functions are run even if a run-time panic occurs.
}
