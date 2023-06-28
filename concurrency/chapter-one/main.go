package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Race Conditions
func checkZero() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}

// Memory Access Synchronization
func checkOne() {
	var memoryAccess sync.Mutex
	var data int

	go func() {
		memoryAccess.Lock()
		data++ // critical section
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 { // critical section
		fmt.Printf("the value is %v.\n", data) // critical section
	} else {
		fmt.Printf("the value is %v.\n", data) // critical section
	}
	memoryAccess.Unlock()
}

// DeadLock
type value struct {
	mu    sync.Mutex
	value int
}

func checkTwo() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         // attempt to enter critical section
		defer v1.mu.Unlock() // exit the critical section before printSum returns

		time.Sleep(2 * time.Second) // here we sleep for a period of time to simulate work
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

// LiveLocks
func checkThree() {
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}
	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}

// Starvation
func checkFour() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}

func main() {
	// checkZero()
	// checkOne()
	// checkTwo()
	// checkThree()
	checkFour()
}

/*
Keywords:
---------
- Race Conditions (to check: "go run -race main.go")
- Atomicity
- Memory Access Sync
- Deadlocks
	- Mutual Exclusion
	- Wait For Condition
	- No Preemption
	- Circular Wait
- Livelocks
- Starvation
*/
