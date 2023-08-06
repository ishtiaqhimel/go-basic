package main

import (
	"fmt"
	"math/rand"
)

func isEven1(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func isEven2(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func main() {
	num := 6
	// if statement
	if num%2 == 0 {
		fmt.Println("The number is even.")
	}

	// if else statement
	// The else statement should start in the same line after
	// the closing curly brace } of the if statement. If not
	// the compiler will complain.
	if num%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	// if ... else if ... else statement
	if num < 5 {
		fmt.Println("The number is less than five.")
	} else if num >= 5 && num < 10 {
		fmt.Println("The number is greater or equal to five but less than ten.")
	} else {
		fmt.Println("The number is greater or equal to ten.")
	}

	// if with assignment
	// One thing to be noted is that x is available only for access
	// from inside the if and else. i.e. the scope of x is limited
	// to the if else blocks. If we try to access x from outside the
	// if or else, the compiler will complain.
	if x := 10; x%2 == 0 {
		fmt.Println(x, "is even.")
	} else {
		fmt.Println(x, "is odd.")
	}

	// Idiomatic Go
	// In Go's philosophy, it is better to avoid unnecessary branches and indentation of code.
	// It is also considered better to return as early as possible.
	// Here isEven2() is better than isEven1()
	fmt.Println(isEven1(num))
	fmt.Println(isEven2(num))

	// switch
	// A switch is a conditional statement that evaluates an expression and compares it against
	// a list of possible matches and executes the corresponding block of code.
	// It can be considered as an idiomatic way of replacing complex if else clauses.
	n := 2
	switch n {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	// Note: Duplicate cases are not allowed
	switch n {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	// case 2: // Not allowed
	//	fmt.Println("Another Two")
	case 3:
		fmt.Println("Three")
	}

	n = 5
	// switch case default
	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Invalid Number!")
	}

	n = 2
	// switch case with fallthrough
	// A fallthrough statement is used to transfer control to the first statement of the case
	// that is present immediately after the case which has been executed.
	// Fallthrough happens even when the case evaluates to false
	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Invalid Number!")
	}

	// Go switch with multiple cases
	switch n {
	case 1, 3, 5:
		fmt.Println("Odd")
	case 2, 4:
		fmt.Println("Even")
	default:
		fmt.Println("Invalid Number!!")
	}

	// switch without any expression
	// If we don't use the expression, the switch statement is true by default.
	switch {
	case 1 == n:
		fmt.Println("One")
	case 3 == n:
		fmt.Println("Three")
	default:
		fmt.Println("Invalid Number!!")
	}

	// switch with statement
	switch num := 3; num {
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Invalid")
	}
	// can't access num out of switch...only accessible in switch scope
	// fmt.Println(num) // Error

	// Break switch
	switch num := -5; {
	case num < 50:
		if num < 0 {
			fmt.Printf("%d is less than 0\n", num)
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	// Breaking the outer for loop
	// When the switch case is inside a for loop, there might be a need to terminate
	// the for loop early. This can be done by labeling the for loop and breaking the
	// for loop using that label inside the switch statement.
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d\n", i)
			break randloop
		}
	}

	// Type Switch
	// data must be interface{}
	var data interface{} = "ishtiaq"
	switch x := data.(type) {
	case nil:
		fmt.Println("num is nil")
	case int:
		fmt.Println(x)
	case float64:
		fmt.Println(x)
	case bool, string:
		fmt.Println("num is bool or string")
	default:
		fmt.Println("don't know the type")
	}
}
