package main

import "fmt"

func main() {
	// declaration of a map [key, value]: [string, int] using the make function
	mp := make(map[string]int)
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
}
