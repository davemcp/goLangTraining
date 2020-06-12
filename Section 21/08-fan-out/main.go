package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go populate(ch1)

	go fanOutIn(ch1, ch2)

	for v := range ch2 {
		fmt.Println("ch2", v)
	}

	fmt.Println("Exiting you fool.")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(ch1, ch2 chan int) {
	var wg sync.WaitGroup
	for v := range ch1 {
		wg.Add(1)
		go func(val int) {
			ch2 <- timeConsumingWork(val)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(ch2)
}

func timeConsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
