package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the SQL database
	db, err := sql.Open("mysql", "root:password@tcp(localhost:2001)/Students")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	// Define a handler function to respond to HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Query returns all matching rows as a Rows struct your code can loop over
		rows, err := db.Query("SELECT name from students;")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Generate a dynamic response based on the retrieved data
		responseText := "Hello from Go Web Server!\n"

		// Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
		for rows.Next() {
			var name string

			// Scan copies the columns from the matched row into the values pointed at by dest
			err := rows.Scan(&name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			responseText += fmt.Sprintf("User: %s\n", name)
		}

		// Write the generated response
		_, err = fmt.Fprint(w, responseText)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)

	// Inform the user that the server is running
	fmt.Printf("Server is running on http://localhost%s\n", serverAddress)

	// Start the HTTP server
	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		fmt.Printf("Server encountered an error: %s\n", err)
	}
}
