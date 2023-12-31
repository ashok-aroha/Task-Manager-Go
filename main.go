package main

import (
	"fmt"
	"net/http"

	a "github.com/ashok-aroha/ACL-Go"
)

// Handler function for a sample endpoint
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is a public endpoint!")
}

func main() {
	// Attach the middleware to the default ServeMux
	http.Handle("/", a.RestrictEndpoints(http.HandlerFunc(handler)))

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
