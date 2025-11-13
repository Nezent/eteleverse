package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nezent/eteleverse/internal"
	"github.com/go-chi/chi/v5"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"Service is healthy","status":"ok"}`
	if rr.Body.String() != expected+"\n" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(internal.HelloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"Hello, World!","status":"success"}`
	if rr.Body.String() != expected+"\n" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestRouter(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/health", internal.HealthHandler)
	r.Get("/hello", internal.HelloHandler)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{"Health endpoint", "/health", http.StatusOK},
		{"Hello endpoint", "/hello", http.StatusOK},
		{"Not found", "/unknown", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code for %s: got %v want %v",
					tt.path, status, tt.expectedStatus)
			}
		})
	}
}
