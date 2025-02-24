package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"search-server/handlers"
	"search-server/middleware"
)

func addMiddlewares(h http.HandlerFunc) http.HandlerFunc {
	return middleware.Logging(h)
}

// main is the entry point for the application.
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", addMiddlewares(handlers.Index))
	http.HandleFunc("/health-check", addMiddlewares(handlers.HealthCheck))
	http.HandleFunc("/hello-world", addMiddlewares(handlers.HelloWorld))

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
