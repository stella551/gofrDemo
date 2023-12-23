package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function to respond to HTTP requests
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// Write a plain text response
		responseText := "Hello, Go Web Server!"
		_, err := fmt.Fprint(w, responseText)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Define the HTTP port
	port := 9000
	serverAddress := fmt.Sprintf(":%d", port)

	// Inform the user that the server is running
	fmt.Printf("Server is running on http://localhost%s\n", serverAddress)

	// Start the HTTP server
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		fmt.Printf("Server encountered an error: %s\n", err)
	}
}
