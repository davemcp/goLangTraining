package main

import (
	"fmt"
	"github.com/davemcp/goLandTraining/ninja012/001/dog"
)

func main()  {
	humanYears := 10
	fmt.Println("Human years", humanYears)
	fmt.Println("Dog years", dog.Years(humanYears))
}
