package main

import "fmt"

// create a type person struct
type person struct {
	Name string
	Age  int
}

// attach a method speak to type person using a pointer receiver
// *person
func (p *person) speak() {
	fmt.Println(p.Name, "says hello!")
}

// create a type human interface
type human interface {
	speak()
}

// show the following in your code
// you CAN pass a value of type *person into saySomething
// you CANNOT pass a value of type person into saySomething
func main() {
	person1 := person{
		Name: "Dave",
		Age:  40,
	}
	person2 := person{
		Name: "Kate",
		Age:  40,
	}
	// Works
	saySomething(&person1)
	saySomething(&person2)
	// Does NOT work
	// saySomething(person1)
	// saySomething(person2)
}

// create func “saySomething”
// have it take in a human as a parameter
// have it call the speak method
func saySomething(h human) {
	h.speak()
}
