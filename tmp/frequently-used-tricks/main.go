package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	// The sort package contains functions for sorting arbitrary data
	// Slice Sorting
	val := []int{1, 3, 4, 2}
	sort.Ints(val)
	fmt.Println(val)

	// to sort an array []a, use it like a slice [l, r) -> r is exclusive, array indices 1 to n
	a := [10]int{2, 3, 4}
	n := 3
	sort.Ints(a[0:n])
	fmt.Println(a)

	// Struct
	personList := []Person{
		{1, "Ishtiaq", 25},
		{2, "Misir", 24},
		{5, "Islam", 26},
		{4, "Abc", 26},
		{3, "Islam", 26},
	}

	// Struct sorting: SliceStable keep i < j stable, if val[i] is same as val[j]
	sort.SliceStable(personList, func(i, j int) bool {
		if personList[i].Age == personList[j].Age {
			return personList[i].Name > personList[j].Name
		} else {
			return personList[i].Age < personList[j].Age
		}
	})

	/*
		less function:
		Returning true from the less function will cause the element at
		index i to be sorted to a lower position than
		index j (The element at index i will come first in the sorted slice).
		Otherwise, the element at index j will come first if false is returned.
	*/

	// not stable sorting
	sort.Slice(personList, func(i, j int) bool {
		if personList[i].Age == personList[j].Age {
			return personList[i].Name > personList[j].Name
		} else {
			return personList[i].Age < personList[j].Age
		}
	})

	fmt.Println(personList)

	// map sorting
	m := map[string]int{
		"sally": 45,
		"john":  5,
		"marcy": 15,
		"duff":  10,
		"tom":   20,
	}
	fmt.Println(m)
	// map is sorted by key, if we need to sort by value then we should make key-value pair slice and sort it
}
