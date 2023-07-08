package main

import "fmt"

func setValueWithoutRef(val int) {
	val = 0
}

func setValueWithRef(val *int) {
	*val = 0
}

func setNewVal(val *int) {
	*val = 10
}

func square(x *int) {
	*x = *x * *x
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	// In Go a pointer is represented using the * (asterisk) character followed by the type of the stored value.
	// In the zero function xPtr is a pointer to an int.

	// * is also used to “dereference” pointer variables. Dereferencing a pointer gives us access to the value the
	// pointer points to. When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers to”.
	// If we try xPtr = 0 instead we will get a compiler error because xPtr is not an int it's a *int, which can only be given another *int.

	//Finally we use the & operator to find the address of a variable. &x returns a *int (pointer to an int) because x is an int.

	currentValue := 10
	fmt.Println(currentValue)
	setValueWithoutRef(currentValue)
	fmt.Println("Without Ref: ", currentValue)
	setValueWithRef(&currentValue)
	fmt.Println("With Ref: ", currentValue)

	// using the new operator
	// new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	// In some programming languages there is a significant difference between using new and &, with great care being
	// needed to eventually delete anything created with new. Go is not like this, it's a garbage collected programming
	// language which means memory is cleaned up automatically when nothing refers to it anymore.
	newVal := new(int)
	fmt.Println(*newVal)
	setNewVal(newVal)
	fmt.Println(*newVal)

	// square calculation
	val := 5
	square(&val)
	fmt.Println(val)

	// swap two values
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	// Do not pass a pointer to an array as an argument to a function. Use slice instead.
}
