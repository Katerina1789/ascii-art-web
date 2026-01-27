package output

import (
	"strings"
)

func RenderAscii(asciiLines []string, userInput string) string {
	var result strings.Builder
	lines := strings.Split(userInput, "\\n")
	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		result.WriteString(BuildAsciiLine(asciiLines, line))
	}
	return result.String()
}

// Prints the ascii art line by line
func BuildAsciiLine(asciiLines []string, line string) string {
	asciiHeight := 8
	blockSize := asciiHeight + 1 // 9
	var finalString strings.Builder
	for row := 0; row < asciiHeight; row++ {
		for _, char := range line {
			charIndex := int(char) - 32
			start := charIndex * blockSize

			finalString.WriteString(asciiLines[(start+row)+1])
		}
		finalString.WriteString("\n")
	}
	return finalString.String()
}
