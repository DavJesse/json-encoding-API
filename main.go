package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int	`json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2/"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		todoItem := Todo{}
		decoder := json.NewDecoder(response.Body)

		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal("Decode error:", err)
		}

		//Convert back to json
		todo, err := json.Marshal(todoItem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(todo))
	}
}
