package main

import (
	"fmt"
	"log"
)

// https://play.golang.org/p/wlEM1tgfQD

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("Your value was less than 42")
		return 0, sqrtError{
			lat:  "0",
			long: "0",
			err:  fmt.Errorf("- your value '%v' is less than zero -", f),
		}
	}
	return 42, nil
}
