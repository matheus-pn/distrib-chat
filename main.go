package main

import (
	"distrib-chat/app"
)

func main() {
	db := app.Database()
	app.Migrate(db) // Auto migrations
	app.Routes(db)
}
