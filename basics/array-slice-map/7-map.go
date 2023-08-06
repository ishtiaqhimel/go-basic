package main

import "fmt"

func main() {
	// declaration of a map [key, value]: [string, int] using the make function
	mp := make(map[string]int) // initialize with empty map[]
	// "var mp map[string]int" here mp is nil, adding any value here will cause panic
	// we need to initialize first to add any value

	// adding items to map
	mp["abc"] = 1
	mp["def"] = 2
	mp["someone"] = 3
	// similar to
	// mp := map[string]int{
	// 	"abc":     1,
	// 	"def":     2,
	// 	"someone": 3,
	// }

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
	// remember, the scope of this value is inside the if-else curly braces only
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

	// delete items from a map
	fmt.Println("Before delete:", mp)
	delete(mp, "def")
	fmt.Println("After delete:", mp)

	// length of the map
	fmt.Println(len(mp))

	// maps are reference types
	fmt.Println("Original mp:", mp)
	mmp := mp
	mmp["someone"] = 5
	fmt.Println("Modified mp:", mp)

	// map's equality check
	// Maps can't be compared using the == operator. The == can be only used to check if a map is nil.
	// One way to check whether two maps are equal is to compare each one's individual elements one by one. The other way is using reflection.
}
