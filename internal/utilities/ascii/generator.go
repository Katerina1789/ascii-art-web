// ASCII art generator using the helper functions (from ascii-art project)
package ascii

import (
	"path/filepath"
)

// GenerateAsciiArt loads the banner, filters input, splits lines, and renders ASCII output
func GenerateAsciiArt(text, banner string) (string, error) {
	// Loads banner file
	bannerPath := filepath.Join("internal/utilities/ascii/banners", banner+".txt")
	asciiLines, err := LoadAsciiFile(bannerPath)
	if err != nil {
		return "", err
	}

	// Filters text first (preserves newlines), then splits into lines
	filteredText, _ := AsciiFilter(text)
	lines := SplitInputLines(filteredText)

	// Renders ASCII output using helper functions
	return RenderAscii(lines, asciiLines), nil
}
