package server

import (
	"ascii-art-web/internal/utilities/ascii"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// HandleHome serves the main page (GET /).
// It validates the path, loads the index template, and renders it.
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Send404(w, "Page not found")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Send404(w, "Template not found")
		return
	}

	tmpl.Execute(w, nil)
}

// HandleAsciiArt processes ASCII art generation requests (POST /ascii-art).
// It validates the method, reads form values, checks the banner name, generates ASCII art, and re-renders the template with the result.
func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== Handler called ===")

	// Only POST is allowed for ASCII generation
	if r.Method != http.MethodPost {
		Send400(w, "Only POST method is allowed")
		return
	}

	// Extract form values
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	fmt.Printf("Text: %q, Banner: %q\n", text, banner)

	// Validate banner selection
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		Send400(w, "Invalid banner name")
		return
	}

	// Generate ASCII art
	result, err := ascii.GenerateAsciiArt(text, banner)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		Send500(w, "Failed to generate ASCII art")
		return
	}
	fmt.Printf("Success: %d bytes\n", len(result))

	// Load template for displaying result
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Send404(w, "Template not found")
		return
	}

	// Data passed to the template
	data := struct {
		Text   string
		Banner string
		Result string
	}{
		Text:   text,
		Banner: banner,
		Result: result,
	}

	tmpl.Execute(w, data)
}

// ExportAsciiArt handles file download (POST /export).
func ExportAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Send400(w, "Only POST method is allowed")
		return
	}

	asciiResult := r.FormValue("ascii")
	if asciiResult == "" {
		Send400(w, "No ASCII art to export")
		return
	}

	data := []byte(asciiResult)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Disposition", "attachment; filename=\"ascii-art.txt\"")

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
