package main

import (
	"fmt"
	"github.com/davemcp/goLandTraining/ninja013/002/quote"
	"github.com/davemcp/goLandTraining/ninja013/002/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
