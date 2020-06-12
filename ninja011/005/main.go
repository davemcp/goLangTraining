package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	sr, err := SqrRoot(-42)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sr)
}

// SqrRoot comment
func SqrRoot(i float64) (float64, error) {
	if i < 0 {
		return 0, fmt.Errorf("The vale '%v' must be greater than zero", i)
	}
	return math.Sqrt(i), nil
}
