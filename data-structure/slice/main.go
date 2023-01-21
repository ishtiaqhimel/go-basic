package main

import "fmt"

func main() {
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
	// This creates a slice that is associated with an underlying int array of length 10.
	// Slices are always associated with some array, and although they can never be longer than the array, they can be smaller.
	// The make function also allows a 3rd parameter, like this:
	newSlice := make([]int, 10, 20)
	// Here, 20 represents the capacity of the underlying array which the slice points to:
	fmt.Println(newSlice)

	// From this example we can easily understand how slice is created from an underlying array
	numbers := [5]int{1, 2, 3, 4, 5}
	sliceNumbers := numbers[2:4] // numbers[low : high] , high is exclusive
	numbers[3] = 7               // changing the array, slice also changes, as slice points to the array
	fmt.Println(sliceNumbers)
	// length & capacity -> they aren't same
	fmt.Println(len(sliceNumbers), cap(sliceNumbers))

	/*
		The length of the slice is the number of elements in the slice.
		The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.

		fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
		fruitslice := fruitarray[1:3]
		fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
	*/

	// Clearing a slice
	a := []string{"A", "B", "C", "D", "E"}
	a = nil
	fmt.Println(a, len(a), cap(a)) // [] 0 0
	// In practice, nil slices and empty slices can often be treated in the same way: they have zero length and capacity,
	// they can be used with the same effect in for loops and append functions, and they even look the same when printed.

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
		graph[i] = make([]int, i+1)
		fmt.Println(graph[i])
	}

	// appending at the end of a slice, we can simply append a single value, but to insert a slice we need to use the (...) spread operator/syntax
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
	// what if we make change in addSlice...will testSlice change?
	addSlice[0] = 50
	fmt.Println(testSlice)
	fmt.Println(addSlice)
}

// More Read:
// Go Slices: usage and internals: https://go.dev/blog/slices-intro
// SliceTricks: https://github.com/golang/go/wiki/SliceTricks
// Iterating Over Slices In Go: https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html
// Arrays, slices (and strings): The mechanics of 'append': https://go.dev/blog/slices
// Using goroutines on loop iterator variables: https://github.com/golang/go/wiki/CommonMistakes
