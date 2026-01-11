package ascii

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestFileExists_SaveLoad(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.txt")

	if FileExists(path) {
		t.Fatalf("file %s should not exist yet", path)
	}

	data := []byte("line1\nline2")
	if err := SaveFile(path, data); err != nil {
		t.Fatalf("SaveFile error: %v", err)
	}

	if !FileExists(path) {
		t.Fatalf("file %s should exist after SaveFile", path)
	}

	lines, err := LoadAsciiFile(path)
	if err != nil {
		t.Fatalf("LoadAsciiFile error: %v", err)
	}
	if len(lines) != 2 || lines[0] != "line1" || lines[1] != "line2" {
		t.Fatalf("unexpected lines: %v", lines)
	}
}

func TestDownloadFile_HTTPServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	}))
	defer srv.Close()

	data, err := DownloadFile(srv.URL)
	if err != nil {
		t.Fatalf("DownloadFile error: %v", err)
	}
	if string(data) != "hello" {
		t.Fatalf("expected %q, got %q", "hello", string(data))
	}
}
