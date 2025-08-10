package main

import (
	"io"
	"net/http"
	"strings"
)

func JobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs := []Gig{
		{"Go Developer",
			"Acme Corp",
			"Remote",
			"Description"},
		{"Backend Engineer (Go)", "Tech Solutions", "Chicago, IL", "Description"},
	}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(jobs)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//get listing.html, and then fill the template
	cleanTemplate := OpenListingTemplate()
	dirtyTemplate := ""

	for _, g := range jobs {

		s := cleanTemplate
		s = strings.ReplaceAll(s, "$company", g.Company)
		s = strings.ReplaceAll(s, "$location", g.Location)
		s = strings.ReplaceAll(s, "$title", g.Title)
		s = strings.ReplaceAll(s, "$description", g.Description)

		s = s + "<br><br>==========================="
		dirtyTemplate = dirtyTemplate + s
	}

	io.WriteString(w, dirtyTemplate)
	//fmt.Fprint(w, len(jobs))

}
