package handlers

import (
	"fmt"
	"net/http"
)

func AssertMethod(r *http.Request, method string) error {
	if r.Method != method {
		return fmt.Errorf("method not allowed: %s", r.Method)
	}

	return nil
}
