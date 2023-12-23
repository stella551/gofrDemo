package main

import (
	"fmt"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/", func(c *gofr.Context) (interface{}, error) {
		responseText := "Hello, Go Web Server!"
		return responseText, nil
	})

	// Define the HTTP port
	app.Server.HTTP.Port = 8080

	// Inform the user that the server is running
	fmt.Printf("Server is running on http://localhost:%d\n", app.Server.HTTP.Port)

	app.Start()
}
