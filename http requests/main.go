package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// HTTP GET REQUEST
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil { // error handler
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil { 
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb) // outputs the request from url

	// HTTP POST REQUEST
	postBody, _ := json.Marshal(map[string]string{
		"name":"bob",
		"email":"bob@gmail.com",
	})
	resBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://postman-echo.com/post", "applications/json", resBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil { 
		log.Fatalln(err)
	}
	s := string(b)
	log.Printf(s)
}