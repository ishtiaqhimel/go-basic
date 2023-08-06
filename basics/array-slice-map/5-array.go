package main

import "fmt"

func main() {
	// declaring an array of int with size 5
	var ara [5]int
	for index, value := range ara {
		ara[index] = index
		// value did not change
		fmt.Println(index, value)
	}
	fmt.Println(ara)

	// assigning values while declaring the array (short hand declaration)
	arr := [10]int{20, 33, 30, 10, 34}
	fmt.Println(arr)

	// we can also break them down in separate lines, in that case an ending comma is a must
	words := [5]string{
		"gone",
		"with",
		"the",
		"wind",
		"again",
	}
	fmt.Println(words)

	// we can declare an array regarding of it's size by using three dots like the code below
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers, " & lenghth is ", len(numbers))

	/*
		The size of the array is a part of the type. Hence [5]int and [25]int are distinct types.
		Because of this, arrays cannot be resized.
		a := [3]int{5, 78, 8}
		var b [5]int
		b = a
		not possible since [3]int and [5]int are distinct types
	*/

	// copying an array is just like assigning a value to a variable, doesn't point to the originl array while assigning, it just copies it!
	aarray := [...]int{1, 2, 3}
	carray := aarray
	carray[1] = 10
	// although we've changed the content on carray, the value on aarray remains the same as before
	fmt.Println(aarray) // prints: [1, 2, 3]
	fmt.Println(carray) // prints: [1, 10, 3]

	// 2-D array declaration, basically created r by c empty arrays, now initiating 0 to r - 1 cells to assign new values
	var matrix [3][3]int
	// matrix[0] = [3]int{1, 0, 0}
	// matrix[1] = [3]int{0, 1, 0}
	// matrix[2] = [3]int{0, 0, 1}
	// matrix := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	matrix[1][1] = 1
	fmt.Println(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
