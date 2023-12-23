package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/hello", func(c *gofr.Context) (interface{}, error) {
		return "Hello World", nil
	})

	app.Server.HTTP.Port = 9000

	app.Start()
}
