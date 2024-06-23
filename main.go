package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId,omitempy"`
	ID        int    `json:"id,omitempy"`
	Title     string `json:"title,omitempy"`
	Completed bool   `json:"completed,omitempy"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2/"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(body))

	if response.StatusCode == http.StatusOK {
		todoItem := Todo{}
		json.NewDecoder(response.Body).Decode(&todoItem)

		fmt.Printf("Data from API: %+v\n", todoItem)

	}
}
