package main

import "fmt"

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

	// recursion
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
