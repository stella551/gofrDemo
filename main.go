package main

import (
	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"main.go/migrations"
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

	err := migration.Migrate("demo-app", dbmigration.NewGorm(app.GORM()), map[string]dbmigration.Migrator{
		"20231119174053": migrations.K20231119174053{}, "20231119175509": migrations.K20231119175509{}}, "UP", app.Logger)
	if err != nil {
		app.Logger.Error(err)

		return
	}

	app.Server.HTTP.Port = 9000

	app.Start()
}
