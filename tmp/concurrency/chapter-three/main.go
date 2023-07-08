package main

import (
	"fmt"
	"sync"
)

func functionA() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func functionB() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Three!")

	// Goroutines: just put 'go' keywork before a function!!!

	functionA()
	functionB()

}

/*
			GOROUTINES
			----------
- How do goroutines actually work?
- Are they OS threads?
- Green threads?
- How many can we create?

--> Goroutines are unique to Go are their deep integration with Go's runtime
--> They are not OS threads & they are not exactly green threads
--> Coroutines and goroutines
--> Go's mechanism for hosting goroutines is an implementation of what’s called
	an M:N scheduler, which means it maps M green threads to N OS threads.
--> Go follows a model of concurrency called the fork-join model.

*/
