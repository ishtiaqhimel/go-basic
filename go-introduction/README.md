## About Go
Go is an open-source programming language developed by Google. It was created with these things in mind:

- efficient execution
- efficient compilation
- ease of programming

That's why Go is very fast, expressive (easy to read and write code) and lightweight.

## Running Go
The simplest way of running Go code is by using an online Go compiler like [The Go Playground](https://go.dev/play/).

You can also easily install and run Go programming locally on your computer. For that, visit [Install Go on Windows, Linux, and macOS](https://go.dev/doc/install).

From CLI we can run go program with `go run <file-name>.go`. The `file-name` file must contain `main()` function from `main` package. If we have many file in same package, we must run `go run .`.

## Hello World
Basic Structure of a Go Program
Here are the things to take away from this tutorial.

1. All Go programs start from the main() function.
2. The bare minimum code structure of a Go program is:
```go
package main

fun main() {
    // code starts here
}
```

## Variables
There are 3 ways to assign values to a variable.
- Method 1: `var x int = 10`
- Method 2: `var x = 10`
</br>(compiler automatically detects the type)
- Method 3: `x := 10`
</br>(shorthand notation to assign values to variables)

Important Notes on Go Variables $-$
- If a variable is not assigned any value, a default value is assigned to it. 
- In Go, every variable must have a data type associated with it. If not, the program throws an error.

Rules of naming Variables $-$
- A variable name consists of alphabets, digits, and an underscore.
Variables cannot have other symbols ( $, @, #, etc).
- Variable names cannot begin with a number.
- A variable name cannot be a reserved word as they are part of the Go syntax like int, type, for, etc.

By the way, we should always try to give meaningful variable names. This makes your code easier to read and understand.

## Data Types
The basic data types in Golang are $-$
| Data Types | Description | Variations |
|:--- |:--- |:---|
| `int` | Integer numbers | int8, int16, int32, int64, int (signed integers), uint8, uint16, uint32, uint32, uint64, uint, uintptr (unsigned integers) |
| `float` | Numbers with decimal points | float32, float64 |
| `complex` | Complex numbers | complex64, complex128 |
| `string` | Sequence of characters (a slice of bytes in Go) |  |
| `bool` | Either true or false | |
| `byte` | A byte (8 bits) of non-negative integers (alias for uint8) | |
| `rune` | Used for characters. Internally used as 32-bit integers (alias for int32) |  |

## I/O in Go
- Using `fmt` package
- Using the bufio.NewReader(os.Stdin)
- File I/O

## Type Casting
Golang provides various predefined functions like `int()`, `float32()`, `string()`, etc to perform explicit type casting.
```go
package main
import "fmt"

func main() {
  
  var floatValue float32 = 5.45

  // type conversion from float to int
  var intValue int = int(floatValue)
 
  fmt.Printf("Float Value is %f\n", floatValue)
  fmt.Printf("Integer Value is %d", intValue)
}
```
 Go doesn't support implicit type casting. 
 ```go
 package main
import "fmt"

func main() {
  
  // initialize integer variable to a floating-point number
  var number int = 4.34 // Error

  fmt.Printf("Number is %f", number)
}
 ```