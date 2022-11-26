package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

/*
The following are the basic types in Go

- bool
- numeric types
	- int8, int16, int32, int64, int (signed integers)
	- uint8, uint16, uint32, uint32, uint64, uint, uintptr (unsigned integers)
	- float32, float64
	- complex64, complex128
	- byte (alias for uint8)
	- rune (alias for int32)
-string
*/

func main() {
	// bool
	// && (and), || (or), ! (not)
	a := true // create a variable with starting value
	b := false
	fmt.Println("a:", a, "b:", b)
	c := (a && b)
	fmt.Println("c:", c)
	d := (a || b)
	fmt.Println("d:", d)
	e := !a
	fmt.Println("e:", e)

	// complex number
	x := 6 + 7i
	y := complex(2, 6) // 2 + 6i
	fmt.Println("complex sum:", x+y)

	// string
	// A string is a slice of bytes in Go. Strings in Go are Unicode compliant and are UTF-8 Encoded.
	// Strings can be created by enclosing a set of characters inside double quotes " ".
	firstName := "Ishtiaq"
	lastName := "Islam"
	name := firstName + " " + lastName
	fmt.Println(name)

	// Since a string is a slice of bytes, it's possible to access each byte of a string.
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	// rune
	name = "Señor"

	// gives wrong output as 'ñ' has unicode value 'U+00F1'
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
	fmt.Printf("\n\n")

	// A rune is a builtin type in Go and it's the alias of int32. Rune represents a Unicode code point in Go.
	runes := []rune(name)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n\n")

	// 'for range' loop returns each position of the bytes and it access each individual runes
	for index, char := range name {
		fmt.Printf("%c starts at byte %d\n", char, index)
	}

	// string length
	/*
		The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string.
		This method takes a string as an argument and returns the number of runes in it.
		As we discussed earlier, len(s) is used to find the number of bytes in the string and it doesn't return the string length.
	*/
	fmt.Println("Runes in name:", utf8.RuneCountInString(name))
	fmt.Println("Bytes in name:", len(name))

	// strings are immutable
	/*
		error: cannot assign to name[0]
		name[0] = 'p'
		fmt.Println(name)
	*/
	// To work with it we need to work with runes
	runes[0] = 'P'
	name = string(runes) // type conversion rune to string
	fmt.Println(name)

	// Type Conversion
	// Go is very strict about explicit typing. There is no automatic type promotion or conversion.
	p := 10
	q := float32(p)
	fmt.Println(reflect.TypeOf(q)) // Printing type of 'q'

	// Type uintptr in Golang is an integer type that is large enough to contain the bit pattern of any pointer. In other words, uintptr is an integer representation of a memory address.
	// Below is how we declare a variable of type uintptr:
	var u uintptr = 0xc82000c290
	fmt.Println(u)
	fmt.Println(reflect.TypeOf(u))
	// https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang
}
