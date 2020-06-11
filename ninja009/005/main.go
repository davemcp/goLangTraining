package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Using goroutines, create an incrementer program
func main() {
	// have a variable to hold the incrementer value
	var incrementor int64

	// use waitgroups to wait for all of your goroutines to finish
	var wg sync.WaitGroup

	// launch a bunch of goroutines
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			// Add 1 to incrementor atomically
			atomic.AddInt64(&incrementor, 1)

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count:", incrementor)
}
