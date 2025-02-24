package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func readBody(r *http.Request) (string, error) {
	text, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return "", err
	}
	return string(text), nil
}

func findLine(text, term string) (int, bool) {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(term)) {
			return i + 1, true
		}
	}
	return 0, false
}

func Index(w http.ResponseWriter, r *http.Request) {
	if err := AssertMethod(r, http.MethodPost); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	term := r.URL.Query().Get("term")
	if term == "" {
		http.Error(w, "missing term parameter", http.StatusBadRequest)
		return
	}

	text, err := readBody(r)
	if err != nil {
		err = fmt.Errorf("error reading request body: %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	line, found := findLine(text, term)
	if found {
		fmt.Fprintf(w, "Found %s on line %d\n", term, line)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
