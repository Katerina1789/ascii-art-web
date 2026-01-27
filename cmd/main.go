package main

import (
	"ascii-art-web/internal/api"
	"ascii-art-web/internal/ascii/utilities/fontHandler"
	"fmt"
	"net/http"
)

func main() {
	//Loads the font files
	if fontHandler.EnsureFontFiles() != nil {
		fmt.Printf("Error retriving the files")
		return
	}
	//Loads the front end files
	api.LoadFrontEndFiles()
	//Handlers for the API
	http.HandleFunc("/ascii-art/api", api.AsciiArtHandler)

	//Starts the server on port 8080
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
