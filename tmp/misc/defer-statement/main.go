package main

import (
	"fmt"
	"sync"
)

func printA(a int) {
	fmt.Println("Value of a:", a)
}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {

	defer wg.Done() // remove 3 wg.Done() with one defer wg.Done()

	if r.length < 0 {
		fmt.Println("Length should be greater than Zero!")
		// wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Println("Width should be greater than Zero!")
		// wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Println("Area:", area)
	// wg.Done()
}

func practicalFunc() {
	var wg sync.WaitGroup

	r1 := rect{-5, 3}
	r2 := rect{5, -3}
	r3 := rect{5, 3}
	rects := []rect{r1, r2, r3}
	for _, r := range rects {
		wg.Add(1)
		go r.area(&wg)
	}

	wg.Wait()
	fmt.Println("All go routines done executing...")
}

func main() {
	// The arguments of a deferred function are evaluated when the defer statement is executed
	// and not when the actual function call is done.
	a := 5
	defer printA(a)
	a = 10
	printA(a)

	// When a function has multiple defer calls, they are pushed on to a stack and executed in
	// Last In First Out (LIFO) order.
	name := "Ishtiaq"
	for _, v := range name {
		defer fmt.Printf("%c\n", v) // "qaithsI"
	}

	// Practical use of defer!!!
	// Defer is used in places where a function call should be executed irrespective of the code flow.
	practicalFunc()
}
