package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"ascii-art-web/internal/utilities/ascii"
)

// HandleHome serves the main page (GET /)
// It validates the path, loads the index template, and renders it
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

// HandleAsciiArt processes ASCII art generation requests (POST /ascii-art)
// It validates the method, reads form values, checks the banner name, generates ASCII art, and re-renders the template with the result
func HandleAsciiArt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== Handler called ===")

	// Only POST is allowed for ASCII generation
	if r.Method != http.MethodPost {
		Send400(w, "Only POST method is allowed")
		return
	}

	// Extracts form values
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	fmt.Printf("Text: %q, Banner: %q\n", text, banner)

	// Validates banner selection
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		Send400(w, "Invalid banner name")
		return
	}

	// Generates ASCII art
	result, err := ascii.GenerateAsciiArt(text, banner)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		Send500(w, "Failed to generate ASCII art")
		return
	}
	fmt.Printf("Success: %d bytes\n", len(result))

	// Loads template for displaying result
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
	format := r.FormValue("format")

	if asciiResult == "" {
		Send400(w, "No ASCII art to export")
		return
	}

	var data []byte
	var contentType, filename string

	switch format {
	case "html":
		htmlContent := fmt.Sprintf(
			`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>ASCII Art</title>
    <style>pre { font-family: monospace; }</style>
</head>
<body>
<pre>%s</pre>
</body>
</html>`, asciiResult)

		data = []byte(htmlContent)
		contentType = "text/html; charset=utf-8"
		filename = "ascii-art.html"

	case "markdown":
		md := fmt.Sprintf("```\n%s\n```", asciiResult)
		data = []byte(md)
		contentType = "text/markdown; charset=utf-8"
		filename = "ascii-art.md"

	default: // txt
		data = []byte(asciiResult)
		contentType = "text/plain; charset=utf-8"
		filename = "ascii-art.txt"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
