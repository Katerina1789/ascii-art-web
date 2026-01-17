/*
Ensures that font file is ready for use and load the file else it downloads this file
*/
package fontHandler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Font struct for creating any pair you want with the filename and the Url that you get this file
type FontResource struct {
	Name string
	Url  string
}

// URL for font files
const asciiUrlStandard = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/standard.txt"
const asciiUrlShadow = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/shadow.txt"
const asciiUrlThinkertoy = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/thinkertoy.txt"

// Font Slice to iterate in order to check if the fonts exist
var fonts = []FontResource{
	{Name: "standard.txt", Url: asciiUrlStandard},
	{Name: "shadow.txt", Url: asciiUrlShadow},
	{Name: "thinkertoy.txt", Url: asciiUrlThinkertoy},
}

// Local path to save font files
// Note: filepath.Join handles the separators, so we don't strictly need the trailing "/" here anymore.
var fontPath = "internal/ascii/banners"

// Make sure the file exist if not it downloads the file to the project
func EnsureFontFiles() error {

	//Create the directory if it doesn't exist (Self-healing)
	// 0755 is standard permission: Owner can write/read/execute, others can read/execute.
	if err := os.MkdirAll(fontPath, 0755); err != nil {
		return fmt.Errorf("could not create font directory: %w", err)
	}

	for _, font := range fonts {
		//Construct the safe, full path once
		fullPath := filepath.Join(fontPath, font.Name)

		// Check the specific file path, not just the name
		if FileExists(fullPath) {
			continue
		}

		fmt.Printf("Downloading %s...\n", font.Name)
		data, err := DownloadFile(font.Url)
		if err != nil {
			return err
		}

		// Save to the same path we checked
		err = SaveFile(fullPath, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// downloads ascii font file
func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status code: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func SaveFile(path string, data []byte) error {
	// 0644: Owner can read/write, everyone else can read.
	return os.WriteFile(path, data, 0644)
}

// Read and split the file contents
func LoadAsciiFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Windows/Unix compatibility for line breaks
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}
