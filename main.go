package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// logRequest logs the details of an HTTP request.
func logRequest(r *http.Request) {
	// Log the request method and URL
	fmt.Printf("Method: %s, URL: %s\n", r.Method, r.URL.String())

	// Log headers
	fmt.Println("Headers:")
	for name, values := range r.Header {
		fmt.Printf("  %s: %s\n", name, strings.Join(values, ", "))
	}

	// Log query parameters
	fmt.Println("Query Parameters:")
	for name, values := range r.URL.Query() {
		fmt.Printf("  %s: %s\n", name, strings.Join(values, ", "))
	}

	// Log the body (if any)
	if r.Body != nil {
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil && err.Error() != "EOF" {
			fmt.Printf("Error reading body: %v\n", err)
			return
		}

		fmt.Printf("Body: %s\n", string(body))
	}
}

// handler is the HTTP handler that logs the request and sends a response.
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request:")
	logRequest(r)

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Request logged successfully!")
}

func main() {
	http.HandleFunc("/", handler)

	// Start the server
	port := "1313"
	log.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
