package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

/*
Converting the error to the underlying type and retrieving more information from the struct fields

type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
*/

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// The simplest way to create a custom error is to use the New function of the errors package.
		// return 0, errors.New("Radius is negative")

		// Let's use the Errorf function and make the program better.
		// return 0, fmt.Errorf("Radius %0.2f is negative", radius)

		// more info with struct
		return 0, &areaError{
			err:    "radius is negative",
			radius: radius,
		}
	}
	return math.Pi * radius * radius, nil
}

func fileErr() {
	f, err := os.Open("/test.txt") // func Open(name string) (*File, error)
	// The idiomatic way of handling errors in Go is to compare the returned error to nil.
	// A nil value indicates that no error has occurred and a non-nil value indicates the presence of an error.

	if err != nil {
		// We can use the 'As' function from errors package to convert the error to it's underlying type.
		// A simple description of 'As' is that it tries to convert the error to a error type and returns either
		// true or false indicating whether the conversion is successful or not.
		var pErr *os.PathError
		if errors.As(err, &pErr) {
			fmt.Println("Failed to open file at path:", pErr.Path)
			return
		}

		fmt.Println("Generic Error:", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")

	// Any type which implements this interface can be used as an error
	/*
		Error type representation:
		type error interface {
			Error() string
		}
	*/
}

func customErr() {
	// Custom Error
	radius := -5.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func provideMoreInfo() {
	// Error using struct type and fields
	radius := -5.0
	area, err := circleArea(radius)
	if err != nil {
		var areaErr *areaError
		// target must be a non-nil pointer
		if errors.As(err, &areaErr) {
			fmt.Println(areaErr.err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(area)
	return
}

func main() {
	fileErr()
	customErr()
	provideMoreInfo()
	// we can also provide more information about the error using methods on struct types

	// Panic and Recover
	defer fmt.Println("deferred call in main")
	firstName := "Ishtiaq"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

// Panic and Recover
/*
The idiomatic way of handling abnormal conditions in a Go program is using errors.
Errors are sufficient for most of the abnormal conditions arising in the program.
But there are some situations where the program cannot continue execution after an
abnormal condition. In this case, we use panic to prematurely terminate the program.
It is possible to regain control of a panicking program using recover.

panic and recover can be considered similar to try-catch-finally idiom in other
languages such as Java except that they are rarely used in Go.
*/

/*
When should panic be used?
> One important factor is that you should avoid panic and recover and use errors
where ever possible. Only in cases where the program just cannot continue execution
should panic and recover mechanism be used.
There are two valid use cases for panic.
1. An unrecoverable error where the program cannot simply continue its execution.
	- One example is a web server that fails to bind to the required port. In this case,
	it's reasonable to panic as there is nothing else to do if the port binding itself fails.
2. A programmer error.
	- Let's say we have a method that accepts a pointer as a parameter and someone
	calls this method using a nil argument. In this case, we can panic as it's a
	programmer error to call a method with nil argument which was expecting a valid pointer.


When a function encounters a panic, its execution is stopped, any deferred functions are
executed and then the control returns to its caller.
*/

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

/*
To Do:
- Getting Stack Trace after Recover
- Panic, Recover and Goroutines
*/
