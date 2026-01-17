package main

import (
	"html/template"
	"net/http"
)

func main() {
	//Root path Parsing to the path of the mainpage.html and sending data via template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("internal/template/mainpage.html")
		data := struct {
			Welcome string
		}{
			Welcome: "Hello from Go server!",
		}
		tmpl.Execute(w, data)
	})
	//Loads the style.css file
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "internal/template/style.css")
	})
	//Starts the server on port 8080
	http.ListenAndServe(":8080", nil)
}
