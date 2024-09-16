# Dockerfile Challenge: Parameterized Environments for Go Application

## Objective

Create a Dockerfile that supports parameterized environments, allowing configuration through environment variables. Your goal is to design a Dockerfile that can accommodate different runtime configurations by utilizing environment variables effectively.

## Scenario

You are tasked with containerizing a simple Go application that requires environment-specific configurations, such as different server ports or custom messages. The Dockerfile should be flexible enough to handle these configurations at runtime through environment variables.

## Go Application Code

Here’s the Go application that you need to containerize:

```go
// main.go
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
