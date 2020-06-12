// Create a struct “customErr” which implements the builtin.error interface.
// Create a func “foo” that has a value of type error as a parameter.
// Create a value of type “customErr” and pass it into “foo”.
package main

import "fmt"

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return ce.err
}

func main() {
	c1 := customErr{
		err: "I am broken",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
