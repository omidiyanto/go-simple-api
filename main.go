package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define handlers for each endpoint
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/version/print", versionHandler)

	// Start the server
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ACTIVE")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "v.1.0.0")
}
