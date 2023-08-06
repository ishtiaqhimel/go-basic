package main

import "fmt"

func main() {
	// Basic Syntax
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Break
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	// Continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	// Nested
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Labels
Outer:
	for i := 0; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println(i, j)
			if i == j {
				break Outer
			}
		}
	}

	// Act as While loop
	i := 1
	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	for {
		fmt.Println("loop")
		break
	}
}
