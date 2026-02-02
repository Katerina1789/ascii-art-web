package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound, "Page not found")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "Template error")
		return
	}

	data := PageData{}
	tmpl.Execute(w, data)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorHandler(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		errorHandler(w, http.StatusBadRequest, "Text cannot be empty")
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		errorHandler(w, http.StatusBadRequest, "Invalid banner type")
		return
	}

	// Simple ASCII art generation (placeholder)
	result := generateAsciiArt(text, banner)

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError, "Template error")
		return
	}

	data := PageData{
		Text:   text,
		Banner: banner,
		Result: result,
	}

	tmpl.Execute(w, data)
}

func generateAsciiArt(text, banner string) string {
	// Placeholder ASCII art generation
	lines := []string{}
	for i := 0; i < 8; i++ {
		line := ""
		for _, char := range text {
			switch banner {
			case "shadow":
				line += fmt.Sprintf("[%c]", char)
			case "thinkertoy":
				line += fmt.Sprintf("<%c>", char)
			default:
				line += fmt.Sprintf("|%c|", char)
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func errorHandler(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		StatusCode int
		Message    string
	}{
		StatusCode: statusCode,
		Message:    message,
	}

	tmpl.Execute(w, data)
}