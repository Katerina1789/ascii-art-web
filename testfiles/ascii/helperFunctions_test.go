package ascii_test

import (
	"ascii-art-web/internal/utilities/ascii"
	"testing"
)

// Test: splitting input into lines
func TestSplitInputLines(t *testing.T) {
	input := "Hello\\nWorld"
	lines := ascii.SplitInputLines(input)
	if len(lines) != 2 || lines[0] != "Hello" || lines[1] != "World" {
		t.Fatalf("expected [Hello World], got %v", lines)
	}
}

// Test: filtering unsupported characters
func TestAsciiFilter(t *testing.T) {
	filtered, removed := ascii.AsciiFilter("Hello\x01World")
	if filtered != "HelloWorld" {
		t.Fatalf("expected HelloWorld, got %s", filtered)
	}
	if len(removed) != 1 {
		t.Fatalf("expected 1 removed char, got %d", len(removed))
	}
}

// Test: rendering a single ASCII line
func TestPrintAsciiLine(t *testing.T) {
	asciiLines := make([]string, 855)
	for i := range asciiLines {
		asciiLines[i] = "X"
	}
	result := ascii.PrintAsciiLine("A", asciiLines)
	if len(result) == 0 {
		t.Fatalf("expected non-empty result")
	}
}

// Test: rendering multiple ASCII blocks
func TestRenderAscii(t *testing.T) {
	asciiLines := make([]string, 855)
	for i := range asciiLines {
		asciiLines[i] = "X"
	}
	lines := []string{"A", "", "B"}
	result := ascii.RenderAscii(lines, asciiLines)
	if len(result) == 0 {
		t.Fatalf("expected non-empty result")
	}
}
