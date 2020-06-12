package main

import (
	"fmt"
)

// Get this code: https://play.golang.org/p/j-EA6003P0

func main() {

	// Anon function
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()
	fmt.Println("Anon Solution", <-c)

	// Buffered solution
	// Only allow one entry to channel, then close
	cb := make(chan int, 1)

	cb <- 42

	fmt.Println("Buffered Solution", <-cb)

}
