package api

import (
	"net/http"
)

// Loads the mainpage.html file
func LoadFrontEndFiles() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, "frontend/mainpage.html")
	})
	//Loads the style.css file
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "frontend/style.css")
	})

	//Loads the script.js file
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		http.ServeFile(w, r, "frontend/script.js")
	})
}
