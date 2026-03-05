package server_test

import (
	"ascii-art-web/internal/utilities/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test: Send400 should return status 400
func TestSend400(t *testing.T) {
	w := httptest.NewRecorder()
	server.Send400(w, "Bad request")

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

// Test: Send404 should return status 404
func TestSend404(t *testing.T) {
	w := httptest.NewRecorder()
	server.Send404(w, "Not found")

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

// Test: Send500 should return status 500
func TestSend500(t *testing.T) {
	w := httptest.NewRecorder()
	server.Send500(w, "Internal error")

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", w.Code)
	}
}
