// Ensures that font files are ready for use and loads the files, otherwise it downloads them
package ascii

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Font struct for creating any pair you want with the filename and the URL from which you get this file
type FontResource struct {
	Name string
	Url  string
}

// URLs for font files
const (
	asciiUrlStandard   = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/standard.txt"
	asciiUrlShadow     = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/shadow.txt"
	asciiUrlThinkertoy = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/thinkertoy.txt"
)

// Font slice to iterate in order to check if the fonts exist
var fonts = []FontResource{
	{Name: "standard.txt", Url: asciiUrlStandard},
	{Name: "shadow.txt", Url: asciiUrlShadow},
	{Name: "thinkertoy.txt", Url: asciiUrlThinkertoy},
}

// Local path to save font files (relative to project root)
var fontPath = "internal/utilities/ascii/banners"

// Makes sure the files exist; if not, downloads them into the project
func EnsureFontFiles() error {
	// Creates the directory if it doesn't exist (self-healing)
	// 0755 is standard permission: owner can write/read/execute, others can read/execute
	if err := os.MkdirAll(fontPath, 0o755); err != nil {
		return fmt.Errorf("could not create font directory: %w", err)
	}

	for _, font := range fonts {
		// Constructs the safe, full path once
		fullPath := filepath.Join(fontPath, font.Name)

		// Checks the specific file path, not just the name
		if FileExists(fullPath) {
			continue
		}

		fmt.Printf("Downloading %s...\n", font.Name)
		data, err := DownloadFile(font.Url)
		if err != nil {
			return err
		}

		// Saves to the same path we checked
		if err := SaveFile(fullPath, data); err != nil {
			return err
		}
	}
	return nil
}

// Checks if a file exists at the given path
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Downloads ASCII font file
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

// Saves the downloaded file to disk
func SaveFile(path string, data []byte) error {
	// 0644: owner can read/write, everyone else can read
	return os.WriteFile(path, data, 0o644)
}

// Reads and splits the file contents
func LoadAsciiFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Windows/Unix compatibility for line breaks
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}
