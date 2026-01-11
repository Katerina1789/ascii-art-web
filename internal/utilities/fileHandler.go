/*
Ensures that font file is ready for use and load the file else it downloads this file
*/
package ascii

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// URL for font files
const asciiUrlStandard = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/standard.txt"
const asciiUrlShadow = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/shadow.txt"
const asciiUrlThinkertoy = "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/thinkertoy.txt"

// File names
const fontStandard = "standard.txt"
const fontShadow = "shadow.txt"
const fontThinkertoy = "thinkertoy.txt"

// Local path to save font files
var fontPath = "internal/utilities/fonts/"

// Slice of Link and file data
var urlList = []string{asciiUrlStandard, asciiUrlShadow, asciiUrlThinkertoy}
var fileNames = []string{fontStandard, fontShadow, fontThinkertoy}

// Make sure the file exist if not it downloads the file to the project
func EnsureFontFiles() error {
	for index, file := range urlList {
		if FileExists(file) {
			continue
		}

		data, err := DownloadFile(file)
		if err != nil {
			return err
		}

		if SaveFile(fontPath+fileNames[index], data) != nil {
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
	return os.WriteFile(path, data, 0o644)
}

// Read and split the file contents
func LoadAsciiFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
