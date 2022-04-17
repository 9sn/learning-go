package main

import "fmt"

func main() {
	// var fruitArr [2]string

	// assiging values

	// fruitArr[0] = "orange"
	// fruitArr[1] = "apple"

	// declare and assign
	// fruitArr := [2]string{"apple", "orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := [2]string{"apple", "orange"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
}