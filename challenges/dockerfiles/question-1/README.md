# Dockerfile Creation Challenge for Go Application

## Objective

Create a Dockerfile to build a Docker image for a simple Go application that serves an HTTP server. The goal is to ensure that the Dockerfile correctly sets up the environment, compiles the Go application, and configures the container to run the application.

## Go Application Code

Here is the Go application that you need to containerize:

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
