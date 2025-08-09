package main

// Job represents a job posting
type Gig struct {
	Title    string `json:"title"`
	Company  string `json:"company"`
	Location string `json:"location"`
}
