package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Add(2)
	go sayHello("first")
	go sayHello("second")
	fmt.Println("Goroutines now\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("All done waitin'")
}

func sayHello(s string) {
	fmt.Println("Goroutine:\t", s)
	for i := 0; i < 10; i++ {
		fmt.Println("Hello stranger", i)
	}
	wg.Done()
}
