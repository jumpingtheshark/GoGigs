package main

import (
	"fmt"
	"io"
	"net/http"
)

func JobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs := []Gig{
		{"Go Developer", "Acme Corp", "Remote"},
		{"Backend Engineer (Go)", "Tech Solutions", "Chicago, IL"},
	}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(jobs)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//get listing.html, and then fill the template
	for _, g := range jobs {
		io.WriteString(w, g.Title)
		io.WriteString(w, "<br>")
		io.WriteString(w, g.Company)
		io.WriteString(w, "<br>")
		io.WriteString(w, g.Location)
		io.WriteString(w, "<br>")
		io.WriteString(w, "<br>")
		io.WriteString(w, "========================================")
		io.WriteString(w, "<br>")
	}

	fmt.Fprint(w, len(jobs))
	w.Write([]byte("<br>"))
	fmt.Fprint(w, "hello world")
	w.Write([]byte("<br>"))
	w.Write([]byte("blah"))
	io.WriteString(w, "fried eggs and kosher veggie ham")
}
