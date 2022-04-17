package main

import "fmt"

func main() {
	// long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// i = i + 1
		// shorter way of i = i + 1
		i++


	}
	// short method
	for i :=1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// fizzbuzz
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else if i % 5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}