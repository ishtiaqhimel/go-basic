package main

import "fmt"

func main() {
	fmt.Println(Total(1, 3, 5, 6))
}

func Total(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}
