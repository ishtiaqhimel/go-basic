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

	// assigning values while declaring the array
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
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 1, 0}
	matrix[2] = [3]int{0, 0, 1}
	fmt.Println(matrix)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	// for slices the size need not to be defined, it's more like C++ container vector <> in terms of usage but only better
	// unlike array, slices actually points to the array a, from which b is assigned, so both go changed, like pointers do
	var mSlice []int // array like declaration, just without assigning the size of the array -> makes it a slice
	fmt.Println(mSlice)
	aslices := []int{1, 2, 3}
	bslices := aslices
	bslices[1] = 10
	fmt.Println(aslices) // 1, 10, 3
	fmt.Println(bslices) // 1, 10, 3

	// if you want to create a slice you should use the builtin function "make" like this:
	mySlice := make([]int, 10)
	fmt.Println(mySlice)
	// This creates a slice that is associated with an underlying float64 array of length 5.
	// Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.
	// The make function also allows a 3rd parameter, like this:
	newSlice := make([]int, 10, 20)
	// Here, 20 represents the capacity of the underlying array which the slice points to:
	fmt.Println(newSlice)
	// Clearing a slice
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a)) // [] 0 0
	// In practice, nil slices and empty slices can often be treated in the same way: they have zero length and capacity,
	// they can be used with the same effect in for loops and append functions, and they even look the same when printed.

	// length & capacity -> they aren't same
	fmt.Println(len(a), cap(a))

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bb := aa[:]   // slice of all elements
	cc := aa[3:]  // slice from 4th element to end
	dd := aa[:6]  // slice first 6 elements, so basically [left, right), left is inclusive & right is exclusive
	ee := aa[3:6] // slice the 4th, 5th, and 6th elements, aa, bb, cc, dd, ee all point to same underlying data
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)
	fmt.Println(ee)

	// 2D slices, how we store graph. e.g: vector <int> graph[N], each of the indices from [0, 10) have an empty slice with them
	var graph [10][]int
	for i := 0; i < 10; i++ {
		fmt.Println(graph[i])
	}

	/*
			The length of the slice is the number of elements in the slice.
			The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.

			fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		    fruitslice := fruitarray[1:3]
		    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	*/

	// appending at the end of a slice, we can simply append a single value, but to insert a slide we need to use the (...) spread operator/syntax
	var testSlice []int
	testSlice = append(testSlice, 10)
	testSlice = append(testSlice, 20)
	fmt.Println("testSlice: ", testSlice)
	var addSlice []int
	addSlice = append(addSlice, 30)
	addSlice = append(addSlice, 40)
	fmt.Println("addSlice", addSlice)
	testSlice = append(testSlice, addSlice...) // see the three dot's (spread operator/syntax)
	fmt.Println("updated testSlice", testSlice)
	fmt.Println("see if the addSlice has changed: ", addSlice) // no, it hasn't

	// declaring slice like this will generate a slice of size 5, initialized with all zeros [0, 0, 0, 0, 0]
	xSlice := make([]int, 5)
	fmt.Println(xSlice)

	// declaration of a map [key, value]: [string, int] using the make function
	mp := make(map[string]int)
	mp["abc"] = 1
	mp["def"] = 2
	mp["someone"] = 3
	// printing the map using key, val pair & ranging over the map
	for key, val := range mp {
		fmt.Println(key, val)
	}
	fmt.Println(mp["non_existant_key"]) // so, non existent key gives zero, so how do we check if it's actually there or not!!!

	// var myMap = make(map[int]int) // declaring map in package level

	// so, to ensure that the key exists in our map, ok must be equal to true
	val, ok := mp["someone"]
	if ok {
		fmt.Println("key exists & it's value is", val)
	} else {
		fmt.Println("key doesn't exist in the map mp")
	}
	// this is cute!! so, here's a semicolon!! we can use this condition to check the existance & print the value if it's present
	// remember, them scope of this value is inside the curly braces only
	if value, ok := mp["someone"]; ok {
		fmt.Println("val: ", value)
	} else {
		fmt.Println("...not present")
	}
	// for some reason if we don't actually need the value & we just only need to check the existance then we can use underscore(_)
	if _, ok := mp["non_existant_key"]; ok {
		fmt.Println("the key exists")
	} else {
		fmt.Println("...not present")
	}
}
