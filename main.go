package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/jobs", JobsHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go Jobs API")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple API for Go job listings.")
}
