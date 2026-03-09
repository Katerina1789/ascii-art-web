package server

import (
	"html/template"
	"net/http"
)

// ErrorData holds the status code and message passed to the error template
type ErrorData struct {
	StatusCode int
	Message    string
}

// Send400 returns a 400 Bad Request response
// Used when the client sends an invalid or incomplete request
func Send400(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	tmpl.Execute(w, ErrorData{StatusCode: 400, Message: message})
}

// Send404 returns a 404 Not Found response
// Used when the requested route or resource does not exist
func Send404(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, ErrorData{StatusCode: 404, Message: message})
}

// Send500 returns a 500 Internal Server Error response
// Used when an unexpected server-side error occurs
func Send500(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, ErrorData{StatusCode: 500, Message: message})
}
