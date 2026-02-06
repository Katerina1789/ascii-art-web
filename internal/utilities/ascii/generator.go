// ASCII art generator using the helper functions (from ascii-art project)
package ascii

import (
	"path/filepath"
)

// GenerateAsciiArt loads the banner, filters input, splits lines, and renders ASCII output
func GenerateAsciiArt(text, banner string) (string, error) {
	// Load banner file
	bannerPath := filepath.Join("internal/utilities/ascii/banners", banner+".txt")
	asciiLines, err := LoadAsciiFile(bannerPath)
	if err != nil {
		return "", err
	}

	// Filter text first (preserves newlines), then split into lines
	filteredText, _ := AsciiFilter(text)
	lines := SplitInputLines(filteredText)

	// Render ASCII output using helper functions
	return RenderAscii(lines, asciiLines), nil
}
