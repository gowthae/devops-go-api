package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d",
			http.StatusOK, rec.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d",
			http.StatusOK, rec.Code)
	}
}

func TestUnknownPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/heath", nil)
	rec := httptest.NewRecorder()

	homeHandler(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d",
			http.StatusNotFound, rec.Code)
	}
}
