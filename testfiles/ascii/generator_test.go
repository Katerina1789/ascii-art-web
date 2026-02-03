package ascii_test

import (
	"ascii-art-web/internal/utilities/ascii"
	"testing"
)

// Test: invalid banner should return an error
func TestGenerateAsciiArt_InvalidBanner(t *testing.T) {
	_, err := ascii.GenerateAsciiArt("Hello", "invalid")
	if err == nil {
		t.Fatalf("expected error for invalid banner")
	}
}

// Test: empty text should produce an empty result
func TestGenerateAsciiArt_EmptyText(t *testing.T) {
	result, _ := ascii.GenerateAsciiArt("", "standard")
	if result != "" {
		t.Fatalf("expected empty result for empty text")
	}
}
