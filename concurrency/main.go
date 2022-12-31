package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	// channel !!!
	done := make(chan bool)

	workersName := []string{"Tommy", "Johnny", "Bobby", "Andy"}

	for _, name := range workersName {
		go task(name, done)
	}
	fmt.Println(<-done)

	// this will waste our time
	// time.Sleep(time.Second * 2)
}

func task(target string, done chan bool) {
	fmt.Println("Working", target)
	time.Sleep(time.Second)
	done <- true
}
