package main

import "net/http"

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1/"

	response, err := http.Get(url)
}
