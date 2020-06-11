package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Using goroutines, create an incrementer program
func main() {
	// have a variable to hold the incrementer value
	incrementer := 0

	// use waitgroups to wait for all of your goroutines to finish
	var wg sync.WaitGroup

	// launch a bunch of goroutines
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		// each goroutine should
		go func() {
			// read the incrementer value
			// store it in a new variable
			v := incrementer
			// yield the processor with runtime.Gosched()
			runtime.Gosched()
			// increment the new variable
			v++
			// write the value in the new variable back to the incrementer variable
			incrementer = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Count:", incrementer)
}
