package main

import (
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>gungstuh</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}