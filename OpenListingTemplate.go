package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenListingTemplate() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}
	dir = dir + "\\html\\"
	filePath := filepath.Join(dir, "listing.html")
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}

	return string(file)

}
