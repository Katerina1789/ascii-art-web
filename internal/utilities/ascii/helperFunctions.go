package ascii

import "strings"

// Handles newline normalization and splitting
func SplitInputLines(input string) []string {
	input = strings.ReplaceAll(input, "\r\n", "\n") // Windows CRLF
	input = strings.ReplaceAll(input, "\\n", "\n")  // Literal \n
	return strings.Split(input, "\n")
}

// Filters unsupported characters while preserving newlines
func AsciiFilter(text string) (string, []string) {
	filtered := ""
	var removed []string

	for _, char := range text {
		if char == '\n' || char == '\r' {
			filtered += string(char)
		} else if char < 32 || char > 126 {
			removed = append(removed, string(char))
		} else {
			filtered += string(char)
		}
	}
	return filtered, removed
}

// Renders one line of ASCII art (8 rows)
func PrintAsciiLine(line string, asciiLines []string) string {
	var result strings.Builder
	asciiHeight := 8
	blockSize := 9

	for row := 0; row < asciiHeight; row++ {
		for _, char := range line {
			if char < 32 || char > 126 {
				continue
			}
			charIndex := int(char) - 32
			start := charIndex * blockSize
			result.WriteString(asciiLines[(start+row)+1])
		}
		result.WriteString("\n")
	}
	return result.String()
}

// Renders multiple lines with blank lines between blocks
func RenderAscii(lines []string, asciiLines []string) string {
	var result strings.Builder

	// Removes trailing empty line from final newline
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for i, line := range lines {
		if line == "" {
			continue
		}

		result.WriteString(PrintAsciiLine(line, asciiLines))

		// Adds one blank line between blocks
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}
