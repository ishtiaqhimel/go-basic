package main

import (
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func booleanType() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	fmt.Println("a && b:", a && b)
	fmt.Println("a || b:", a || b)
	fmt.Println("!a", !a)
}

func numericType() {
	// int, float
	a := 6   // int
	b := 9.9 // float64
	fmt.Printf("%T %T\n", a, b)

	var c int32 = int32(a) // can't directly assign
	fmt.Printf("%T\n", c)

	// Complex Number
	x := 6 + 7i // complex128
	fmt.Printf("%T: %v\n", x, x)
	y := complex(2.2, 6.7) // complex128
	fmt.Printf("%T: %v\n", y, y)
	fmt.Println("complex sum:", x+y)

	// Rune
	name := "Señor"

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
		fmt.Printf("%c starts at index %d\n", char, index)
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

	// uintptr
	// Type uintptr in Golang is an integer type that is large enough to contain the bit pattern of any pointer. In other words, uintptr is an integer representation of a memory address.
	// Below is how we declare a variable of type uintptr:
	var u uintptr = 0xc82000c290
	fmt.Println(u)
	fmt.Println(reflect.TypeOf(u))
	// https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang
}

func stringType() {

	firstName := "Ishtiaq"
	lastName := "Islam"
	name := firstName + " " + lastName
	fmt.Println(name)

	// access each byte of string
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
}

func arrayType() {
	x := [5]int{} // var x [5]int
	x[3] = 45
	fmt.Println(x, len(x))

	// multi-dimensional
	var y [3][2][5]int // same as [3]([2]([5]int))
	fmt.Println(y, len(y), len(y[0]), len(y[0][0]))
}

func sliceType() {
	// Here x same as y
	x := make([]int, 5, 10)
	fmt.Println(x, len(x), cap(x))

	y := new([10]int)[0:5]
	fmt.Println(y, len(y), cap(y))

	// if capacity is not mentioned, then the capacity is set to length
	z := []int{2, 3, 4}
	fmt.Println(z, len(z), cap(z))

	// Go includes two built-in functions to assist with slices: append and copy
	// example of append()
	s1 := []int{1, 3, 5}
	s2 := []int{2, 4, 6}
	s1 = append(s1, s2...) // to append a slice
	fmt.Println(s1, s2)

	s1 = append(s1, 7, 8) // to append any value
	fmt.Println(s1)

	// example of copy()
	v1 := []int{1, 2, 3, 4, 5}
	v2 := make([]int, 3)
	fmt.Println(v1, v2)
	copy(v2, v1) // v2 has room for only 3 elements, so copy only first 3 elements
	fmt.Println(v1, v2)

	// only copies the values, here we can see no one is affected by updating another
	v1[0] = 9
	v2[1] = 6
	fmt.Println(v1, v2)

	// if we copy by assigning then the address will be assigned
	c1 := []int{1, 2, 3, 4, 5}
	c2 := make([]int, 3)
	fmt.Println(c1, c2)
	c2 = c1 // c2 holds the address of c1
	fmt.Println(c1, c2)
	c1[0] = 9 // changing c1 will affect c2
	c2[2] = 6 // changing c2 will affect c1
	fmt.Println(c1, c2)

	// By append we changed the address so now c1 will not be affected
	c2 = append(c2, 10)
	c2[1] = 11
	fmt.Println(c1, c2)
}

func typeConversion() {
	// Go is very strict about explicit typing. There is no automatic type promotion or conversion.
	p := 10
	q := float32(p)
	fmt.Println(reflect.TypeOf(q)) // Printing type of 'q'

	// string to int converting, make sure that you import "strconv"
	s := "123"
	fmt.Printf("%T, %v\n", s, s)
	// num, err := strconv.Atoi(s)
	num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		fmt.Println("some error occurred while conversion")
	} else {
		fmt.Println("no error, successfully converted!")
	}
	fmt.Printf("%T, %v\n", num, num)

	// int to string conversion, import "strconv" & look that Itoa returns only one value
	h := 123
	fmt.Printf("%T, %v\n", h, h)
	// t := strconv.Itoa(h)
	t := strconv.FormatInt(int64(h), 10)
	fmt.Printf("%T, %v\n", t, t)
	// another way using 'sprintf'
	ss := fmt.Sprintf("%d", h)
	fmt.Println(reflect.TypeOf(ss), ss)
}

func main() {
	fmt.Println("Boolean Type:")
	booleanType()
	fmt.Println()

	fmt.Println("Numeric Type:")
	numericType()
	fmt.Println()

	fmt.Println("String Type:")
	stringType()
	fmt.Println()

	fmt.Println("Array Type:")
	arrayType()
	fmt.Println()

	fmt.Println("Slice Type:")
	sliceType()
	fmt.Println()

	// Type Conversion
	fmt.Println("Type Conversion:")
	typeConversion()
}
