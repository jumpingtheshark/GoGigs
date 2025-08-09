package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/jobs", jobsHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go Jobs API")
}

func jobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs := []Gig{
		{"Go Developer", "Acme Corp", "Remote"},
		{"Backend Engineer (Go)", "Tech Solutions", "Chicago, IL"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple API for Go job listings.")
}
