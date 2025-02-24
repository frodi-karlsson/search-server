package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"search-server/handlers"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name            string
		method          string
		statusCode      int
		wantTextContent string
	}{
		{
			name:            "get",
			method:          http.MethodGet,
			statusCode:      http.StatusOK,
			wantTextContent: "<h1>Hello World</h1>",
		},
		{
			name:            "post",
			method:          http.MethodPost,
			statusCode:      http.StatusMethodNotAllowed,
			wantTextContent: "method not allowed: POST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/", nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.HelloWorld)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d, got %d", tt.statusCode, rr.Code)
			}

			body := rr.Body.String()
			matched, err := regexp.MatchString(tt.wantTextContent, body)
			if err != nil {
				t.Fatalf("could not match string: %v", err)
			}

			if !matched {
				t.Errorf("expected body to contain %q, got %q", tt.wantTextContent, body)
			}
		})
	}
}
