package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/hello", func(ctx *gofr.Context) (interface{}, error) {
		return "Hello Server", nil
	})

	app.GET("/count", func(ctx *gofr.Context) (interface{}, error) {
		var count int
		row, _ := ctx.DB().QueryContext(ctx, "select count(*) from user")
		row.Scan(&count)

		return count, nil
	})

	app.Server.HTTP.Port = 9000

	app.Start()
}
