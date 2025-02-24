package middleware_test

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"search-server/middleware"
	"testing"
)

const logTimestampLen = 20

func TestLogging(t *testing.T) {
	tests := []struct {
		name    string
		method  string
		query   string
		wantLog string
	}{
		{
			name:    "GET request",
			method:  http.MethodGet,
			wantLog: "Request: GET /",
		},
		{
			name:    "POST request",
			method:  http.MethodPost,
			wantLog: "Request: POST /",
		},
		{
			name:    "GET request with query",
			method:  http.MethodGet,
			query:   "term=gopher",
			wantLog: "Request: GET /?term=gopher",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqString := "/"
			if tt.query != "" {
				reqString += "?" + tt.query
			}

			req, err := http.NewRequest(tt.method, reqString, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			scanner, reader, writer := mockLogger(t)
			defer resetLogger(reader, writer)

			handler := http.HandlerFunc(middleware.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
			handler.ServeHTTP(httptest.NewRecorder(), req)

			if !scanner.Scan() {
				t.Fatalf("expected log entry, got none")
			}

			gotLog := scanner.Text()[logTimestampLen:]
			if gotLog != tt.wantLog {
				t.Errorf("expected log entry %q, got %q", tt.wantLog, gotLog)
			}

		})
	}
}

func mockLogger(t *testing.T) (*bufio.Scanner, *os.File, *os.File) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("could not create pipe: %v", err)
	}
	log.SetOutput(writer)

	return bufio.NewScanner(reader), reader, writer
}

func resetLogger(reader *os.File, writer *os.File) {
	err := reader.Close()
	if err != nil {
		fmt.Println("error closing reader was ", err)
	}
	if err = writer.Close(); err != nil {
		fmt.Println("error closing writer was ", err)
	}
	log.SetOutput(os.Stderr)
}
