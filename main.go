package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// this was generated using https://mholt.github.io/json-to-go/
type Data []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Data

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	for _, curData := range record {
		fmt.Println("User Id. = ", curData.UserID)
		fmt.Println("ID   = ", curData.ID)
		fmt.Println("Title  = ", curData.Title)
		fmt.Println("Body   = ", curData.Body)

	}

}
