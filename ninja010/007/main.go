package main

import (
	"fmt"
	"sync"
)

//     write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

func main() {
	c := make(chan int)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(start int) {
				start = start * 10
				for j := start; j <= start+10; j++ {
					c <- j
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
}
