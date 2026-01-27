package api

import (
	"ascii-art-web/internal/ascii/utilities/fontHandler"
	"ascii-art-web/internal/ascii/utilities/output"
	"log"
	"net/http"
	"path/filepath"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	//Get the message and banner from the form
	message := r.FormValue("message")
	banner := r.FormValue("banner")

	if message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}

	// Validate banner choice to prevent path traversal or invalid files
	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {
		banner = "standard" // default fallback
	}
	//Loads the font file
	fontPath := filepath.Join("internal/ascii/banners", banner+".txt")
	fontLines, err := fontHandler.LoadAsciiFile(fontPath)
	if err != nil {
		log.Printf("Error loading font file: %v", err)
		http.Error(w, "Error loading font", http.StatusInternalServerError)
		return
	}
	//Renders the ascii art
	result := output.RenderAscii(fontLines, message)
	//Writes the result to the response
	w.Write([]byte(result))
}
