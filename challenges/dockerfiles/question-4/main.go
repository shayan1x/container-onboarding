package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello, Docker!"
	}
	fmt.Fprintf(w, message)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("Starting server on port %s...\n", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
