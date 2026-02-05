package ascii_test

import (
	"ascii-art-web/internal/utilities/ascii"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

// Test: file creation, saving, loading, and existence checks
func TestFileExists_SaveLoad(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.txt")

	// File should not exist before saving
	if ascii.FileExists(path) {
		t.Fatalf("file %s should not exist yet", path)
	}

	// Save file to disk
	data := []byte("line1\nline2")
	if err := ascii.SaveFile(path, data); err != nil {
		t.Fatalf("SaveFile error: %v", err)
	}

	// File should exist after saving
	if !ascii.FileExists(path) {
		t.Fatalf("file %s should exist after SaveFile", path)
	}

	// Load file and verify contents
	lines, err := ascii.LoadAsciiFile(path)
	if err != nil {
		t.Fatalf("LoadAsciiFile error: %v", err)
	}
	if len(lines) != 2 || lines[0] != "line1" || lines[1] != "line2" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

// Test: downloading a file using a mock HTTP server
func TestDownloadFile_HTTPServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	}))
	defer srv.Close()

	// Download from mock server
	data, err := ascii.DownloadFile(srv.URL)
	if err != nil {
		t.Fatalf("DownloadFile error: %v", err)
	}

	// Validate downloaded content
	if string(data) != "hello" {
		t.Fatalf("expected %q, got %q", "hello", string(data))
	}
}
