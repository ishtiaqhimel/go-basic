package main

import "fmt"

func main() {
	number := 3

	// switch case default
	switch number {

	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")

	case 4:
		fmt.Println("Four")

	case 5:
		fmt.Println("Five")

	default:
		fmt.Println("Invalid Number!")
	}

	// switch case with fallthrough
	// A fallthrough statement is used to transfer control to the first statement of the case that is present immediately after the case which has been executed.
	switch number {

	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")

	case 3:
		fmt.Println("Three")
		fallthrough

	case 4:
		fmt.Println("Four")
		fallthrough

	case 5:
		fmt.Println("Five")

	default:
		fmt.Println("Invalid Number!")
	}

	// Go switch with multiple cases
	switch number {
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
	case 1 == number:
		fmt.Println("One")
	case 3 == number:
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
}
