package main

import (
	"fmt"
	"os"
)

// basic struct: func functionName(argument argument type, ) return type {}
// func sumOfTwoNumber(a, b int) int {} ... same thing
func sumOfTwoNumber(a, b int) (res int) {
	res = a + b
	return
}

func sumOfAList(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func func2() int {
	fmt.Println("Return from func2()")
	return 100
}

func func1() int {
	fmt.Println("Return from func1()")
	return func2()
}

func multiValRetFunc() (int, float32) {
	return 5, 5.9
}

// variadic function, can take any number of arguments as the last parameter only
// By using ... before the type name of the last parameter we can indicate that it takes zero or more of those parameters.
// This is precisely how the fmt.Println function is implemented

func sumUsingVariadicFunc(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

// no built-int function for calculating minimum, variadic function can be used here to implement one
// it takes two or more arguments & calculates the minimum value among them, max can be implemented similarly

func min(a int, b int, args ...int) int {
	mn := a
	if b < mn {
		mn = b
	}
	for _, val := range args {
		if val < mn {
			mn = val
		}
	}
	return mn
}

func max(a int, b int, args ...int) int {
	mx := a
	if b > mx {
		mx = b
	}
	for _, val := range args {
		if val > mx {
			mx = val
		}
	}
	return mx
}

// a simple recursive function to calculate the factorial of a number
func recFunc(n int) int {
	if n < 0 {
		panic("negative number in factorial function")
	}
	if n <= 1 {
		return 1
	}
	return n * recFunc(n-1)
}

func intSeq() func() int {
	i := 0
	fmt.Println("intseq out, comes here only while definig the closure!")
	return func() int {
		fmt.Println("intseq in, comes here every time this func is called!")
		i++
		return i
	}
}

// Higher order function, see the def. in the main function
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func returnAnonymFunc() func(msg string) {
	return func(msg string) {
		fmt.Println("hello from return anonym func. " + msg)
	}
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {
	fmt.Println(sumOfTwoNumber(10, 20))

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sumOfAList(slice1))

	// functions are built-up in a stack like fashion
	fmt.Println(func1())

	// a function can return multiple values, types must given in similar fashion
	val1, val2 := multiValRetFunc()
	fmt.Println(val1, val2)

	fmt.Println(sumUsingVariadicFunc(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(min(5, 4, 2, 1, 3), max(1, 4, 5, 2, 3))

	fmt.Println(recFunc(5), recFunc(0))

	// Closure -> it is possible to create functions insides of functions, like this:
	// here, add is a local variable that has the type func(int, int) int (a function that takes two ints and returns an int).
	addTwo := func(x, y int) int {
		return x + y
	}
	fmt.Println(addTwo(10, 20))

	// When you create a local function like this it also has access to other local variables(scope!!), should print 1, 2
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment(), increment())

	swap := func(x, y int) (int, int) {
		return y, x
	}
	// swap two integers
	x, y := swap(10, 20)
	fmt.Println(x, y)
	// or we can just do this! cute! :p
	a := 100
	b := 200
	a, b = b, a
	fmt.Println(a, b)

	// added another closure example from go by exmple:
	nextInt := intSeq()    //  so, nextInt is not an integer, it's actually storing a function, that has access to the local scoped vars
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	newInts := intSeq()    // redifining i = 0
	fmt.Println(newInts()) // 1

	// Higher order functions: The definition of Higher-order function:
	// a function which does at least one of the following: i. takes one or more functions as arguments ii. returns a function as its result
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	// an anonymous function
	func(msg string) {
		fmt.Println("Hello from anonym. func " + msg)
	}("ishtiaq")
	// it returns an anonymous function
	anonymFunc := returnAnonymFunc()
	anonymFunc("hello")

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
