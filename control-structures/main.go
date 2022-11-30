package main

import "fmt"

func main() {
	// for loop
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	// for loop range
	arr := []int{1, 2, 4}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// switch
	i := 4
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
