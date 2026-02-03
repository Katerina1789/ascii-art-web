package main

import (
	"ascii-art-web/internal/utilities/ascii"
	"ascii-art-web/internal/utilities/server"
	"fmt"
	"net/http"
)

func main() {
	// Ensure all banner font files exist before starting the server
	if ascii.EnsureFontFiles() != nil {
		fmt.Printf("Error retrieving the files")
		return
	}

	// Route: Home page (GET /)
	http.HandleFunc("/", server.HandleHome)

	// Route: ASCII Art generation (POST /ascii-art)
	http.HandleFunc("/ascii-art", server.HandleAsciiArt)

	// Route: Static assets (CSS) under /static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start HTTP server
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
