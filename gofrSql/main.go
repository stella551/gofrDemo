package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/count", func(ctx *gofr.Context) (interface{}, error) {

		rows, err := app.DB().Query("SELECT name from students1;")
		if err != nil {
			log.Fatal(err)
		}

		var (
			name  string
			names []string
		)

		// Next prepares the next result row for reading with the Scan method. It returns true on success, or false if there is no next result row or an error happened while preparing it
		for rows.Next() {
			// Scan copies the columns from the matched row into the values pointed at by dest
			err := rows.Scan(&name)
			if err != nil {
				return "", err
			}

			names = append(names, name)
		}

		return names, nil
	})

	// Define the HTTP port
	app.Server.HTTP.Port = 8080

	// Inform the user that the server is running
	fmt.Printf("Server is running on http://localhost:%d\n", app.Server.HTTP.Port)

	app.Start()
}
