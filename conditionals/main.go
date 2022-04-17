package main

import "fmt"

func main() {
	x := 10
	y := 10

	// if else
	if x <= y {
		fmt.Printf("%d is less then or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less then %d", y, x)
	}

	// else if

	color := "purple"

	if color == "blue" {
		fmt.Println("Color is blue")
	} else if color == "purple" {
		fmt.Println("Color is purple")
	} else {
		fmt.Println("Color isnt blue or purple")
	}

	// switch

	switch color {
	case "purple":
		fmt.Println("color is purple")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or purple")
	}
}