package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqStr := "Request: " + r.Method + " " + r.URL.Path
		if r.URL.RawQuery != "" {
			reqStr += "?" + r.URL.RawQuery
		}

		log.Println(reqStr)
		next.ServeHTTP(w, r)
	})
}
