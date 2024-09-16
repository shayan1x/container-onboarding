# Dockerfile Optimization Challenge: Reducing Image Size

## Objective

Optimize the Dockerfile to minimize the size of the final Docker image for the provided Go application. Your goal is to implement best practices for reducing the image size while ensuring that the application runs correctly.

## Go Application Code

The Go application to be optimized is a simple HTTP server:

```go
// main.go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
