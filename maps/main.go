package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	// assign key values
	// emails["bob"] = "bob@gmail.com"
	// emails["michelle"] = "michelle@gmail.com"

	// declare map and add key values
	emails := map[string]string{"bob":"bob@gmail.com", "michelle": "michelle@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	// delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}