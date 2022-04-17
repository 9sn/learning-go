package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	// var name = "Zimzy" same thing with adding string
	// var name string = "Zimzy"
	var age int = 14
	var isCool bool = true

	// short-hand variable
	name, email := "zimzy", "visualstudiocode@riseup.net"
	// name := "Zimzy" 
	size := 1.3

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size) // tells you what the type is

}