package server_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"ascii-art-web/internal/utilities/server"
)

// Test: home handler should return 200 or 404
func TestHandleHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	server.HandleHome(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Fatalf("expected 200 or 404, got %d", w.Code)
	}
}

// Test: ASCII handler should reject invalid method
func TestHandleAsciiArt_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	w := httptest.NewRecorder()

	server.HandleAsciiArt(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

// Test: ASCII handler should reject invalid banner
func TestHandleAsciiArt_InvalidBanner(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Hello")
	form.Add("banner", "invalid")

	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	server.HandleAsciiArt(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

// Test: ASCII export
