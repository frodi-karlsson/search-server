package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"search-server/handlers"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		term       string
		statusCode int
		want       string
	}{
		{
			name:       "found",
			body:       "hello world",
			term:       "world",
			statusCode: http.StatusOK,
			want:       "Found world on line 1\n",
		},
		{
			name:       "case insensitive",
			body:       "hello world",
			term:       "WORLD",
			statusCode: http.StatusOK,
			want:       "Found WORLD on line 1\n",
		},
		{
			name:       "multiple lines",
			body:       "hello\nworld",
			term:       "world",
			statusCode: http.StatusOK,
			want:       "Found world on line 2\n",
		},
		{
			name:       "not found",
			body:       "hello world",
			term:       "gopher",
			statusCode: http.StatusNotFound,
			want:       "",
		},
		{
			name:       "missing term",
			body:       "hello world",
			term:       "",
			statusCode: http.StatusBadRequest,
			want:       "missing term parameter\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/?term="+tt.term, strings.NewReader(tt.body))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handlers.Index(rr, req)

			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, rr.Code)
			}

			if got := rr.Body.String(); got != tt.want {
				t.Errorf("expected body %q, got %q", tt.want, got)
			}
		})
	}

	t.Run("should not allow GET requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handlers.Index(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status code %d, got %d", http.StatusMethodNotAllowed, rr.Code)
		}
	})
}

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handlers.HealthCheck(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	if got := rr.Body.String(); got != "OK\n" {
		t.Errorf("expected body %q, got %q", "OK\n", got)
	}
}
