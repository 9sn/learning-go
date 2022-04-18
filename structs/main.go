package main

import (
	"fmt"
	"strconv"
)

// define person struct
type person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// greeting method
func (p person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
	// init person using struct
	person1 := person{firstName: "Bob", lastName: "Smith", city: "san diego", gender: "m", age: 25}
	// alternative
	// person1 := person{"Bob", "Smith", "san diego", "m", 25}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	fmt.Println(person1.greet())
}